package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new <filename>",
	Short: "create a new pulp",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var postResponse PostResponse
		content, _ := os.ReadFile(args[0])
		body := PostData{Content: string(content), Language: "js"}
		Client.SetHeader("Content-Type", "application/json").SetBody(body).SetResult(&postResponse).Post(Api)
		fmt.Printf("https://j2me.eu.org/%s\n", postResponse.Key)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
