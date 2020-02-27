package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("debug", false, "enable debug")
	cmd.Flags().String("addr", "localhost:50051", "address")

	// set the link between flags
	viper.BindPFlag("debug", cmd.Flags().Lookup("debug"))
	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
}
