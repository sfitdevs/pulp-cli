package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new <filename>",
	Short: "Create a new pulp to share",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var postResponse PostResponse
		var file string = string(args[0])
		var language []string = strings.Split(file, ".")
		content, _ := os.ReadFile(file)
		body := PostData{Content: string(content), Language: string(language[len(language)-1])}
		Client.SetHeader("Content-Type", "application/json").SetBody(body).SetResult(&postResponse).Post(Api)
		fmt.Println("- code: " + postResponse.Key + "\n- link: " + color.InBlue(fmt.Sprintf("https://p.aadi.lol/%s", postResponse.Key)) + "\n- access key: " + color.InGray(postResponse.AccessKey) + " (used to delete pulp)")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
