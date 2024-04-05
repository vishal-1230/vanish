package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the specified model",
	Long:  `This command will stop the specified model on the local machine.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   stop,
}

func stop(cmd *cobra.Command, args []string) {
	fmt.Println("stop called")
}

func init() {
	stopCmd.PersistentFlags().StringP("port", "p", "8080", "Model running on this port will be stopped")
	stopCmd.PersistentFlags().StringP("mid", "m", "", "Model ID of the model to be stopped")
	rootCmd.AddCommand(stopCmd)
}
