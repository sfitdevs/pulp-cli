package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save <pulp-code>",
	Short: "Save the pulp to local filesystem",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var getResponse GetResponse
		_, err := Client.SetResult(&getResponse).Get(fmt.Sprintf("%s/%s", Api, args[0]))
		if err != nil {
			fmt.Println("- error: " + err.Error())
		} else {
			if getResponse.Key != "" {
				f, err := os.Create(fmt.Sprintf("%s.%s", getResponse.Key, getResponse.Language))
				if err != nil {
					fmt.Println("- error: " + err.Error())
				} else {
					f.WriteString(getResponse.Content)
					fmt.Printf("- message: pulp saved successfully as %s.%s\n", getResponse.Key, getResponse.Language)
					defer f.Close()
				}
			} else {
				fmt.Println("- error: pulp not found")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
