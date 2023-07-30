
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yourproject",
	Short: "A commandline tool to run a long-running command and get a notification when the command completes or errors out.",
	Long: `A commandline tool built with love in Go. 
It allows you to run a long-running command and get a notification on telegram/whatsapp/discord when the command completes or errors out.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Here you will initialize your configuration and environment variables
}
