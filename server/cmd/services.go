package cmd

import (
	"context"
	"net"

	pb "github.com/katallaxie/content_streaming_msk/proto"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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

		ss := grpc.NewServer()
		pb.RegisterMicroServer(ss, &service{})

		ready()

		if err := ss.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}

type service struct {
	pb.UnimplementedMicroServer
}

func (s *service) Insert(ctx context.Context, req *pb.Insert_Request) (*pb.Insert_Response, error) {
	uuid := uuid.New()

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("brokers"), config)
	if err != nil {
		return nil, err
	}

	defer producer.Close()

	// extracing the item and assigning an uuid.
	// the uuid is a message id, so that the log compaction works.
	item := req.GetItem()

	switch i := item.Item.(type) {
	case *pb.Item_Image:
		i.Image.Uuid = uuid.String()
	case *pb.Item_Article:
		i.Article.Uuid = uuid.String()
	}

	// marshal into new message
	b, err := proto.Marshal(item)
	if err != nil {
		return nil, err
	}

	msg := &sarama.ProducerMessage{Topic: viper.GetString("topic"), Value: sarama.ByteEncoder(b)}
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
	config.ClientID = "go-kafka-consumer"
	config.Consumer.Return.Errors = true

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
