package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the specified model",
	Long:  `This command will restart the specified model on the local machine on the next available port or on the specified port using -p.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   restart,
}

func restart(cmd *cobra.Command, args []string) {
	fmt.Println("restart called")
}

func init() {
	restartCmd.PersistentFlags().StringP("port", "p", "8080", "Port on which the model should run")
	rootCmd.AddCommand(restartCmd)
}
