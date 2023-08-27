package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get more information about the pulp",
	Run: func(cmd *cobra.Command, args []string) {
		var getResponse GetResponse
		Client.SetResult(&getResponse).Get(fmt.Sprintf("%s/%s", Api, args[0]))
		fmt.Printf("Key: %s\nLanguage: %s\nCreated: %s\nSize: %d bytes\nViews: %d\n", getResponse.Key, getResponse.Language, time.UnixMilli(getResponse.TimeStamp).Format("02/01/2006 15:04 MST"), getResponse.Size, getResponse.Views)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
