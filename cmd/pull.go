package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"vanish/types"
	"vanish/utils"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls the specified model to the local machine",
	Long:  `This command will pull the specified model from the remote server to the local machine.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   pull,
}

const (
	green  = "\033[32m"
	reset  = "\033[0m"
	yellow = "\033[33m"
	red    = "\033[31m"
)

func pull(cmd *cobra.Command, args []string) {
	var model = args[0]
	fmt.Println("[", yellow, "•", reset, "]", "Pulling model:", model)

	if config, err := downloadModel(model); err != nil {
		fmt.Println("[", red, "✗", reset, "]", err)
	} else {
		// adding the model to the vanish-config.json
		jsonFileHandler := &utils.FileHandler{
			Type_of_file: "json",
		}
		var data vanish_types.VanishConfig
		err := jsonFileHandler.ReadFile("vanish-config.json", &data)

		if err != nil {
			fmt.Println("[", red, "✗", reset, "]", err)
			return
		} else {
			modelData := vanish_types.Model{
				Name:         model,
				Size:         "Unknown",
				Description:  "Unknown",
				HFURL:        "https://huggingface.co/" + model,
				Architecture: config.(map[string]interface{})["architectures"].([]interface{})[0].(string),
			}
			data.Models = append(data.Models, modelData)
			err := jsonFileHandler.WriteFile("vanish-config.json", data)
			if err != nil {
				fmt.Println("[", red, "✗", reset, "]", err)
				return
			}
			fmt.Println("[", green, "✔", reset, "]", "Model added to vanish-config.json")
		}
	}
}

func init() {
	rootCmd.AddCommand(pullCmd)
}

func downloadModel(modelName string) (interface{}, error) {

	// Define the path to the models directory
	modelsDir := "models"

	// Install Git lfs if not already installed
	gitLfsCmd := exec.Command("git", "lfs", "install")
	if err := gitLfsCmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to install Git LFS: %v", err)
	} else {
		fmt.Println("[", green, "✔", reset, "]", "Git LFS installed successfully.")
	}

	// Create the models directory if it doesn't exist
	if _, err := os.Stat(modelsDir); os.IsNotExist(err) {
		os.Mkdir(modelsDir, 0755)
	}

	// Define the path where the model will be cloned
	modelPath := filepath.Join(modelsDir, modelName)

	var config interface{}
	fileHandler := &utils.FileHandler{
		Type_of_file: "json",
	}

	// Check if the model directory already exists
	if _, err := os.Stat(modelPath); err == nil {
		fmt.Println("[", green, "✔", reset, "]", "Model already exists locally.")
		err := fileHandler.ReadFile(filepath.Join(modelPath, "config.json"), &config)
		if err != nil {
			return nil, errors.New("failed to read config file: " + err.Error())
		} else {
			return config, nil
		}
	}

	// Execute Git clone command
	gitCloneCmd := exec.Command("git", "clone", "https://huggingface.co/"+modelName, modelPath)
	gitCloneCmd.Stdout = os.Stdout
	if err := gitCloneCmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to clone model: %v", err)
	}

	// Running git lfs pull in the models/<model-name> directory
	gitLfsPullCmd := exec.Command("git", "lfs", "pull")
	gitLfsPullCmd.Dir = modelPath
	gitLfsPullCmd.Stdout = os.Stdout
	if err := gitLfsPullCmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to pull model: %v", err)
	}

	fmt.Println("[", green, "✔", reset, "]", "Model '"+modelName+"' downloaded successfully.")

	err := fileHandler.ReadFile(filepath.Join(modelPath, "config.json"), &config)
	if err != nil {
		return nil, errors.New("failed to read config file: " + err.Error())
	} else {
		return config, nil
	}
}
