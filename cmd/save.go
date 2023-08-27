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
		Client.SetResult(&getResponse).Get(fmt.Sprintf("%s/%s", Api, args[0]))
		f, _ := os.Create(fmt.Sprintf("%s.%s", getResponse.Key, getResponse.Language))
		defer f.Close()
		f.WriteString(getResponse.Content)
		fmt.Printf("- pulp saved successfully as %s.%s\n", getResponse.Key, getResponse.Language)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
