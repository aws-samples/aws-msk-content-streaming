package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addFlags(cmd *cobra.Command) {
	cmd.Flags().String("log-format", "text", "log format")
	cmd.Flags().String("log-level", "info", "log-level")
	cmd.Flags().String("addr", ":9090", "address")
	cmd.Flags().String("topic", "monolog", "kafka, topic")
	cmd.Flags().StringSlice("brokers", []string{"localhost:9092"}, "kafka brokers")
	cmd.Flags().String("version", "2.3.2", "kafka version")

	// set the link between flags
	viper.BindPFlag("log-format", cmd.Flags().Lookup("log-format"))
	viper.BindPFlag("log-level", cmd.Flags().Lookup("log-level"))
	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("topic", cmd.Flags().Lookup("topic"))
	viper.BindPFlag("brokers", cmd.Flags().Lookup("brokers"))
	viper.BindPFlag("version", cmd.Flags().Lookup("version"))
}
