package cmd

import (
	"fmt"
	"strconv"
	"text/tabwriter"
	"time"
	vanish_types "vanish/types"
	"vanish/utils"

	// vanish_types "vanish/types"

	"github.com/spf13/cobra"
	// "vanish/types"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all the pulled models",
	Long: `This command will display a list of all the models that are available 
	locally to be used by vanish.`,
	Run: list,
}

func list(cmd *cobra.Command, args []string) {
	fmt.Println("list called")
	jsonFileHandler := &utils.FileHandler{
		Type_of_file: "json",
	}
	var data vanish_types.VanishConfig
	err := jsonFileHandler.ReadFile("vanish-config.json", &data)

	if err == nil {
		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 1, ' ', tabwriter.Debug)

		if cmd.Flag("all").Value.String() == "true" {
			fmt.Fprintln(w, "Model Name\tDisk Size\tDescription")
			for _, model := range data.Models {
				fmt.Fprintln(w, model.Name+"\t"+model.Size+"\t"+model.Description)
			}
			w.Flush()
			return
		} else if cmd.Flag("running").Value.String() == "true" {
			fmt.Fprintln(w, "Model ID\tModel Name\tPort\tTime Elapsed")
			for _, service := range data.ServicesRunning {
				for _, model := range data.Models {
					if service.Name == model.Name {
						var time_elapsed time.Duration
						utils.GetTimeElapsed(service.StartedAt, &time_elapsed)
						fmt.Fprintln(w, service.ModelId+"\t"+service.Name+"\t"+strconv.Itoa(service.Port)+"\t"+time_elapsed.String())
					}
				}
			}
			w.Flush()
			return
		} else if cmd.Flag("installed").Value.String() == "true" {
			fmt.Fprintln(w, "Model Name\tDisk Size\tPorts\tDescription")
			for _, model := range data.Models {
				ports := ""
				for _, service := range data.ServicesRunning {
					if service.Name == model.Name {
						ports += strconv.Itoa(service.Port) + ","
					}
				}
				if len(ports) > 0 {
					ports = ports[:len(ports)-1]
				}
				fmt.Fprintln(w, model.Name+"\t"+model.Size+"\t"+ports+"\t"+model.Description)
			}
			w.Flush()
			return
		}
	} else {
		fmt.Println(err)
	}
}

func init() {
	listCmd.PersistentFlags().BoolP("all", "a", false, "List all the models from the remote server")
	listCmd.PersistentFlags().BoolP("installed", "i", true, "List all the installed models locally")
	listCmd.PersistentFlags().BoolP("running", "r", false, "List all the running models locally")
	rootCmd.AddCommand(listCmd)
}
