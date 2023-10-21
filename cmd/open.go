package cmd

import (
	"fmt"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:     "open <pulp-code>",
	Short:   "Open the pulp in browser",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"o"},
	Run: func(cmd *cobra.Command, args []string) {
		err := browser.OpenURL(fmt.Sprintf("%s/%s", Host, args[0]))
		if err != nil {
			fmt.Printf("- cant open the pulp\nplease visit directly: %s/%s", Host, args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
