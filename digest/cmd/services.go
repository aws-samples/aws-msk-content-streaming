package cmd

import (
	"context"
	"fmt"
	"net"

	pb "github.com/katallaxie/content_streaming_msk/proto"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
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
		pb.RegisterMonologServer(ss, &service{})

		ready()

		if err := ss.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}

type service struct {
	pb.UnimplementedMonologServer
}

func (s *service) Insert(ctx context.Context, req *pb.Insert_Request) (*pb.Insert_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}

func (s *service) Update(ctx context.Context, req *pb.Update_Request) (*pb.Update_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func (s *service) ListArticles(req *pb.ListArticles_Request, srv pb.Monolog_ListArticlesServer) error {
	fmt.Println("here")
	config := sarama.NewConfig()
	config.ClientID = "go-kafka-consumer"
	config.Consumer.Return.Errors = true

	// Create new consumer
	master, err := sarama.NewConsumer(viper.GetStringSlice("brokers"), config)
	if err != nil {
		fmt.Println(err)
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
	if nil != err {
		fmt.Printf("Topic %v Partitions: %v", topic, partitions)
		panic(err)
	}

	fmt.Println(" Start consuming topic ", topic)

	go func(topic string, consumer sarama.PartitionConsumer) {
		for {
			select {
			case consumerError := <-consumer.Errors():
				errors <- consumerError
				fmt.Println("consumerError: ", consumerError.Err)

			case msg := <-consumer.Messages():
				consumers <- msg
				fmt.Println("Got message on topic ", topic, msg.Value)
			}
		}
	}(topic, consumer)

	return consumers, errors
}
