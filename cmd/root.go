package cmd

import (
	"os"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "vanish",
	Short: "Vanish is a command line tool to quikly have ML models running on your local machine.",
	Long:  `Vanish runs a specified machine learning model on the local machine and 
	exposes a set of REST API to interact with the model.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
