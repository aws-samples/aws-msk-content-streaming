package cmd

import (
	"context"
	"net"

	pb "github.com/katallaxie/content_streaming_msk/proto"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type root struct {
	logger *log.Entry
}

type server struct {
	pb.UnimplementedMonologServer
}

func (s *server) Insert(ctx context.Context, req *pb.Insert_Request) (*pb.Insert_Response, error) {
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

	// logging where we send this
	log.Infof("send to partition %s with offset %s", partition, offset)

	return &pb.Insert_Response{Uuid: uuid.String()}, nil
}

func (s *server) Update(ctx context.Context, req *pb.Update_Request) (*pb.Update_Response, error) {
	return nil, nil
}

func runE(cmd *cobra.Command, args []string) error {
	lis, err := net.Listen("tcp", viper.GetString("port"))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterMonologServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
