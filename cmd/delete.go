package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete <access-key>",
	Args:    cobra.MinimumNArgs(1),
	Short:   "Delete a pulp using access-key",
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		body := DelData{AccessKey: args[0]}
		resp, err := Client.SetHeader("Content-Type", "application/json").SetBody(body).Delete(API)
		statuscode := resp.StatusCode()
		if err != nil {
			fmt.Println("- error: " + err.Error())
		} else {
			if statuscode == 200 {
				fmt.Println("- message: pulp deleted successfully")
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
	rootCmd.AddCommand(deleteCmd)
}
