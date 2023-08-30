package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:     "view <pulp-code>",
	Short:   "View the content of the pulp",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		var getResponse GetResponse
		resp, err := Client.SetResult(&getResponse).Get(API + args[0])
		statuscode := resp.StatusCode()
		if err != nil {
			fmt.Println("- error: " + err.Error())
		} else {
			if statuscode == 200 {
				fmt.Println(getResponse.Content)
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
	rootCmd.AddCommand(viewCmd)
}
