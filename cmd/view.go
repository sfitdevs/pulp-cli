package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the content of the pulp",
	Run: func(cmd *cobra.Command, args []string) {
		var getResponse GetResponse
		Client.SetResult(&getResponse).Get(fmt.Sprintf("%s/%s", Api, args[0]))
		fmt.Println(getResponse.Content)
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
