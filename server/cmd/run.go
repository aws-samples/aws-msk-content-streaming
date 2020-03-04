package cmd

import (
	"context"

	"github.com/andersnormal/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type root struct {
	logger *log.Entry
}

func runE(cmd *cobra.Command, args []string) error {
	// create a new root
	root := new(root)

	// init logger
	root.logger = log.WithFields(log.Fields{
		"verbose": viper.GetBool("verbose"),
		"brokers": viper.GetStringSlice("brokers"),
		"topic":   viper.GetString("topic"),
	})

	// create root context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create server
	s, _ := server.WithContext(ctx)

	// log ...
	root.logger.Info("Starting server ...")

	// debug listener
	debug := server.NewDebugListener(
		server.WithPprof(),
		server.WithStatusAddr(":8443"),
	)
	s.Listen(debug, true)

	// listen for grpc
	s.Listen(&srv{}, true)

	// listen for the server and wait for it to fail,
	// or for sys interrupts
	if err := s.Wait(); err != nil {
		root.logger.Error(err)
	}

	// noop
	return nil
}
