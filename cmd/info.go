package cmd

import (
	"fmt"
	"time"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info <pulp-code>",
	Short:   "Get more information about the pulp",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		var getResponse GetResponse
		resp, err := Client.SetResult(&getResponse).Get(API + args[0])
		statuscode := resp.StatusCode()
		if err != nil {
			fmt.Println("- error: " + err.Error())
		} else {
			if statuscode == 200 {
				fmt.Printf("- key: %s\n- link: %shttps://pulp.deta.dev/%s%s\n- language: %s\n- created: %s\n- size: %d bytes\n- views: %d\n", getResponse.Key, color.Underline, getResponse.Key, color.Reset, getResponse.Language, time.UnixMilli(getResponse.TimeStamp).Format("02/01/2006 15:04 MST"), getResponse.Size, getResponse.Views)
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
	rootCmd.AddCommand(infoCmd)
}
