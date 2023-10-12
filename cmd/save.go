package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var filename string

var saveCmd = &cobra.Command{
	Use:     "save <pulp-code>",
	Short:   "Save the pulp to local filesystem",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"s", "dl", "clone"},
	Run: func(cmd *cobra.Command, args []string) {
		var getResponse GetResponse
		resp, err := Client.SetResult(&getResponse).Get(API + args[0])
		statuscode := resp.StatusCode()
		if err != nil {
			fmt.Println("- error: " + err.Error())
		} else {
			if statuscode == 200 {
				if filename == "" {
					filename = getResponse.Key + "." + getResponse.Language
				}
				file, err := os.Create(filename)
				if err != nil {
					fmt.Println("- error: " + err.Error())
				} else {
					file.WriteString(getResponse.Content)
					fmt.Printf("- message: pulp saved successfully as %s\n", filename)
					defer file.Close()
				}
			} else if statuscode == 404 {
				fmt.Println("- error: pulp not found")
			} else if statuscode == 500 {
				fmt.Println("- error: internal server error")
			} else {
				fmt.Println("- error: something went wrong")
			}
		}
	},
}

func init() {
	saveCmd.Flags().StringVarP(&filename, "filename", "f", "", "filename to save file as")
	rootCmd.AddCommand(saveCmd)
}
