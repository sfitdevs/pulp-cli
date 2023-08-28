package cmd

import (
	"fmt"
	"time"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info <pulp-code>",
	Short: "Get more information about the pulp",
	Run: func(cmd *cobra.Command, args []string) {
		var getResponse GetResponse
		_, err := Client.SetResult(&getResponse).Get(fmt.Sprintf("%s/%s", Api, args[0]))
		if err != nil {
			fmt.Println("- error: " + err.Error())
		} else {
			if getResponse.Key != "" {
				fmt.Printf("- key: %s\n- link: %shttps://p.aadi.lol/%s%s\n- language: %s\n- created: %s\n- size: %d bytes\n- views: %d\n", getResponse.Key, color.Underline, getResponse.Key, color.Reset, getResponse.Language, time.UnixMilli(getResponse.TimeStamp).Format("02/01/2006 15:04 MST"), getResponse.Size, getResponse.Views)
			} else {
				fmt.Println("- error: pulp not found")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
