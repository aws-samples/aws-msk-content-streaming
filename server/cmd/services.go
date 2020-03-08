package cmd

import (
	"context"
	"crypto/tls"
	"math"
	"net"
	"time"

	pb "github.com/katallaxie/content_streaming_msk/proto"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

type srv struct {
}

func (s *srv) Start(ctx context.Context, ready func()) func() error {
	return func() error {
		lis, err := net.Listen("tcp", viper.GetString("addr"))
		if err != nil {
			return err
		}

		sarama.Logger = log.New()

		srv := &service{}

		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = true
		srv.tlsCfg = tlsConfig

		if err := s.Setup(tlsConfig); err != nil {
			return err
		}

		var kaep = keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
			PermitWithoutStream: true,            // Allow pings even when there are no active streams
		}

		var kasp = keepalive.ServerParameters{
			MaxConnectionIdle:     time.Duration(math.MaxInt64), // If a client is idle for 15 seconds, send a GOAWAY
			MaxConnectionAge:      time.Duration(math.MaxInt64), // If any connection is alive for more than 30 seconds, send a GOAWAY
			MaxConnectionAgeGrace: 5 * time.Second,              // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
			Time:                  5 * time.Second,              // Ping the client if it is idle for 5 seconds to ensure the connection is still active
			Timeout:               1 * time.Second,              // Wait 1 second for the ping ack before assuming the connection is dead
		}

		ss := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
		pb.RegisterMicroServer(ss, srv)
		health.RegisterHealthServer(ss, srv)

		ready()

		if err := ss.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}

func (s *srv) Setup(tlsCfg *tls.Config) error {
	config := sarama.NewConfig()
	config.ClientID = "server"
	config.Version = sarama.V2_2_0_0
	config.Net.TLS.Enable = true
	config.Net.TLS.Config = tlsCfg

	client, err := sarama.NewClient(viper.GetStringSlice("brokers"), config)
	if err != nil {
		return err
	}

	defer client.Close()

	admin, err := sarama.NewClusterAdminFromClient(client)
	if err != nil {
		return err
	}

	defer admin.Close()

	topics, err := admin.ListTopics()
	if err != nil {
		return err
	}

	if _, ok := topics[viper.GetString("topic")]; ok {
		return nil
	}

	details := &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 2,
		ConfigEntries:     map[string]*string{"cleanup.policy": asString("compact"), "segment.ms": asString("100"), "min.cleanable.dirty.ratio": asString("0.01")},
	}

	return admin.CreateTopic(viper.GetString("topic"), details, false)
}

type service struct {
	tlsCfg *tls.Config
	pb.UnimplementedMicroServer
}

func (s *service) Insert(ctx context.Context, req *pb.Insert_Request) (*pb.Insert_Response, error) {
	uuid := uuid.New()

	config := sarama.NewConfig()
	config.ClientID = "server"
	config.Producer.Return.Successes = true
	config.Version = sarama.V2_2_0_0
	config.Net.TLS.Enable = true
	config.Net.TLS.Config = s.tlsCfg

	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("brokers"), config)
	if err != nil {
		return nil, err
	}

	defer producer.Close()

	// extracing the item and assigning an uuid.
	// the uuid is a message id, so that the log compaction works.
	item := req.GetItem()

	switch i := item.Item.(type) {
	case *pb.Item_Article:
		i.Article.Uuid = uuid.String()
	}

	// marshal into new message
	b, err := proto.Marshal(item)
	if err != nil {
		return nil, err
	}

	msg := &sarama.ProducerMessage{Topic: viper.GetString("topic"), Key: sarama.StringEncoder(uuid.String()), Value: sarama.ByteEncoder(b)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return nil, err
	}

	// logging where we send this
	log.Infof("send to partition %s with offset %s", partition, offset)

	return &pb.Insert_Response{Uuid: uuid.String()}, nil
}

func (s *service) Update(ctx context.Context, req *pb.Update_Request) (*pb.Update_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func (s *service) ListArticles(req *pb.ListArticles_Request, srv pb.Micro_ListArticlesServer) error {
	config := sarama.NewConfig()
	config.ClientID = "server"
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_2_0_0
	config.Net.TLS.Enable = true
	config.Net.TLS.Config = s.tlsCfg

	// Create new consumer
	master, err := sarama.NewConsumer(viper.GetStringSlice("brokers"), config)
	if err != nil {
		return err
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	consumer, errors := consume(viper.GetString("topic"), master)

	for {
		select {
		case msg := <-consumer:
			item := &pb.Item{}
			// marshal into new message
			if err := proto.Unmarshal(msg.Value, item); err != nil {
				return err
			}

			if err := srv.Send(&pb.ListArticles_Response{Articles: []*pb.Article{item.GetArticle()}}); err != nil {
				return err
			}
		case err := <-errors:
			return err
		}
	}
}

func (s *service) Check(ctx context.Context, in *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	log.Infof("Received Check request: %v", in)

	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}, nil
}

func (s *service) Watch(in *health.HealthCheckRequest, _ health.Health_WatchServer) error {
	log.Infof("Received Watch request: %v", in)

	return status.Error(codes.Unimplemented, "unimplemented")
}

func consume(topic string, master sarama.Consumer) (chan *sarama.ConsumerMessage, chan *sarama.ConsumerError) {
	consumers := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)

	partitions, _ := master.Partitions(topic)
	// this only consumes partition no 1, you would probably want to consume all partitions
	consumer, err := master.ConsumePartition(topic, partitions[0], sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	log.Infof(" Start consuming topic ", topic)

	go func(topic string, consumer sarama.PartitionConsumer) {
		for {
			select {
			case consumerError := <-consumer.Errors():
				errors <- consumerError
				log.Info("consumerError: %s", consumerError.Err)
			case msg := <-consumer.Messages():
				consumers <- msg
				log.Infof("Got message on topic ", topic)
			}
		}
	}(topic, consumer)

	return consumers, errors
}

func asString(s string) *string {
	return &s
}
