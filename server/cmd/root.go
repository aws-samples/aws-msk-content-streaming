package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs the gRPC service",
	Long:  `Not yet`,
	RunE:  runE,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// initialize cobra
	cobra.OnInitialize(initConfig)

	// adding flags
	addFlags(RootCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// set the default format, which is basically text
	log.SetFormatter(&log.TextFormatter{})

	viper.AutomaticEnv() // read in environment variables that match

	// config logger
	logConfig()
}

func logConfig() {
	// reset log format
	if viper.GetString("log-format") == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	// set the configured log level
	if level, err := log.ParseLevel(viper.GetString("log-level")); err == nil {
		log.SetLevel(level)
	}
}
