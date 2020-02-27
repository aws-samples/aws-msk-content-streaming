package cmd

import (
	"context"
	"time"

	pb "github.com/katallaxie/content_streaming_msk/proto"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type root struct {
	logger *log.Entry
}

func runE(cmd *cobra.Command, args []string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(viper.GetString("addr"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewMonologClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.Insert_Request{Item: &pb.Item{Item: &pb.Item_Article{Article: &pb.Article{Body: "test"}}}}

	r, err := c.Insert(ctx, req)
	if err != nil {
		return err
	}

	log.Info(r)

	return nil
}
