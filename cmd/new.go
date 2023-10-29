package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

var description string

var newCmd = &cobra.Command{
	Use:     "new <filename>",
	Short:   "Create a new pulp to share",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"n", "create"},
	Run: func(cmd *cobra.Command, args []string) {
		var postResponse PostResponse
		var file string = string(args[0])
		var language []string = strings.Split(file, ".")
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("- error: " + err.Error())
		} else {
			body := PostData{Content: string(content), Language: string(language[len(language)-1]), Description: string(description)}
			resp, err := Client.SetHeader("Content-Type", "application/json").SetBody(body).SetResult(&postResponse).Post(API)
			statuscode := resp.StatusCode()
			if err != nil {
				fmt.Println("- error: " + err.Error())
			} else {
				if statuscode == 200 {
					var pulpsStored []Pulp
					var pulpsAppend []Pulp
					var newPulp Pulp
					content, _ := os.ReadFile(Path)
					json.Unmarshal(content, &pulpsStored)
					newPulp.Key = postResponse.Key
					newPulp.AccessKey = postResponse.AccessKey
					newPulp.TimeStamp = postResponse.TimeStamp
					pulpsAppend = append(pulpsStored, newPulp)
					newContent, _ := json.Marshal(pulpsAppend)
					os.WriteFile(Path, newContent, 0644)
					fmt.Println("- code: " + postResponse.Key + "\n- link: \033[4mhttps://pulp.deta.dev/" + postResponse.Key + "\033[0m\n- access key: " + color.InGray(postResponse.AccessKey) + " (used to delete pulp)")
				} else if statuscode == 500 {
					fmt.Println("- error: internal server error")
				} else {
					fmt.Println("- error: something went wrong")
				}
			}
		}
	},
}

func init() {
	newCmd.Flags().StringVarP(&description, "description", "d", "", "description of the file")
	rootCmd.AddCommand(newCmd)
}
