package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/wiliehidayat87/testMQ/src/config"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long:  `Cobra is a CLI library for Go that empowers applications.`,
	}
)

func init() {

	loc, _ := time.LoadLocation(config.APP_TZ)
	time.Local = loc

	/**
	 * WEBSERVER SERVICE
	 */
	rootCmd.AddCommand(serverCmd)

	/**
	 * CONSUME SERVICE
	 */
	rootCmd.AddCommand(consumerCmd)

	/**
	 * PUBLISH SERVICE
	 */
	rootCmd.AddCommand(publisherCmd)

}

func Execute() error {
	return rootCmd.Execute()
}
