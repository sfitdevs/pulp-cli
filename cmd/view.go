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
		_, err := Client.SetResult(&getResponse).Get(fmt.Sprintf("%s/%s", Api, args[0]))
		if err != nil {
			fmt.Println("- error: " + err.Error())
		} else {
			if getResponse.Key != "" {
				fmt.Println(getResponse.Content)
			} else {
				fmt.Println("- error: pulp not found")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
