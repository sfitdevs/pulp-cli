package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save <pulp-id>",
	Short: "Saves the pulp to local filesystem",
	Run: func(cmd *cobra.Command, args []string) {
		var getResponse GetResponse
		Client.SetResult(&getResponse).Get(fmt.Sprintf("%s/%s", Api, args[0]))
		f, _ := os.Create(fmt.Sprintf("%s.%s", getResponse.Key, getResponse.Language))
		defer f.Close()
		f.WriteString(getResponse.Content)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
