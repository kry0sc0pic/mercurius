package cmd

import (
	"os"
	"strings"

	"github.com/kry0sc0pic/mercurius/notification"
	"github.com/kry0sc0pic/mercurius/utils"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a long-running command",
	Long:  `Run a long-running command and get a notification on telegram/whatsapp/discord when the command completes or errors out.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := strings.Join(args, " ")
		err := utils.ExecuteCommand(command)
		if err != nil {
			notifyError(err)
			os.Exit(1)
		} else {
			notifyCompletion(command)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func notifyError(err error) {
	// Here you can add logic to send error notification
	notification.NewDiscordNotifier().Send("Command execution failed with error: " + err.Error())
}

func notifyCompletion(command string) {
	// Here you can add logic to send completion notification
	// notification.discord.Send("Command '" + command + "' completed successfully")
	notification.NewDiscordNotifier().Send("Command '" + command + "' completed successfully")
}
