package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the list of your created pulps",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Remove(Path)
		if err != nil {
			fmt.Println("- list is already empty")
		} else {
			fmt.Println("- successfully cleared list")
		}
	},
}

func init() {
	ListCmd.AddCommand(clearCmd)
}
