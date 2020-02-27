package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("debug", false, "enable debug")
	cmd.Flags().Bool("verbose", false, "enable verbose")
	cmd.Flags().String("api-key", "", "api key")

	// set the link between flags
	viper.BindPFlag("debug", cmd.Flags().Lookup("debug"))
	viper.BindPFlag("verbose", cmd.Flags().Lookup("verbose"))
	viper.BindPFlag("api-key", cmd.Flags().Lookup("api-key"))
}
