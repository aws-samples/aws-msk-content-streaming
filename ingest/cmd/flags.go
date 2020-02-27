package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("debug", false, "enable debug")
	cmd.Flags().String("port", ":50051", "port")
	cmd.Flags().Bool("verbose", false, "enable verbose")
	cmd.Flags().String("topic", "monolog", "monolog")
	cmd.Flags().StringSlice("brokers", []string{"localhost:9092"}, "brokers")

	// set the link between flags
	viper.BindPFlag("debug", cmd.Flags().Lookup("debug"))
	viper.BindPFlag("verbose", cmd.Flags().Lookup("verbose"))
	viper.BindPFlag("brokers", cmd.Flags().Lookup("brokers"))
	viper.BindPFlag("port", cmd.Flags().Lookup("port"))
	viper.BindPFlag("topic", cmd.Flags().Lookup("topic"))
}
