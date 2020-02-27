package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("debug", false, "enable debug")
	cmd.Flags().Bool("verbose", false, "enable verbose")
	cmd.Flags().String("addr", ":9090", "address")
	cmd.Flags().String("topic", "monolog", "topic")
	cmd.Flags().StringSlice("brokers", []string{"localhost:9092"}, "brokers")
	cmd.Flags().String("version", "2.3.2", "kafka version")

	// set the link between flags
	viper.BindPFlag("debug", cmd.Flags().Lookup("debug"))
	viper.BindPFlag("verbose", cmd.Flags().Lookup("verbose"))
	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("topic", cmd.Flags().Lookup("topic"))
	viper.BindPFlag("brokers", cmd.Flags().Lookup("brokers"))
	viper.BindPFlag("version", cmd.Flags().Lookup("version"))
}
