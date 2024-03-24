package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls the specified model to the local machine",
	Long:  `This command will pull the specified model from the remote server to the local machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull called")
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
