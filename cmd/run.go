package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the specified model",
	Long:  `This command will run the specified model on the local machine on the next available port or on the specified port using -p.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	// printing name & port of the model
	fmt.Println("run called")
	fmt.Println("Model Name: ", args[0])
	port, _ := cmd.Flags().GetString("port")
	fmt.Println("Port: ", port)
}

func init() {
	runCmd.PersistentFlags().StringP("port", "p", "5051", "Port on which the model should run")
	rootCmd.AddCommand(runCmd)
}
