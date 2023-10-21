package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var pulps []Pulp
		content, _ := os.ReadFile(Path)
		json.Unmarshal(content, &pulps)
		if len(pulps) == 0 {
			fmt.Println("- no pulps created by you")
		} else {
			tbl := table.New("no.", "key", "date created", "access key")
			for i := 0; i < len(pulps); i++ {
				tbl.AddRow(i+1, pulps[i].Key, time.UnixMilli(pulps[i].TimeStamp).Format("02/01/06 15:04"), pulps[i].AccessKey)
			}
			tbl.Print()
		}
	},
}

func init() {
	rootCmd.AddCommand(ListCmd)
}
