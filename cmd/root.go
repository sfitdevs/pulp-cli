package cmd

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var API string = "https://p.aadi.lol/api/"
var Host string = "https://p.aadi.lol"
var Client *resty.Request = resty.New().R()
var home, _ = os.UserHomeDir()
var Path = fmt.Sprintf("%s\\.pulp\\pulps.json", home)

var rootCmd = &cobra.Command{
	Use:   "pulp",
	Short: "Pulp - share code seamlessly and effectively!",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func initConfig() {
	if _, err := os.Stat(Path); err != nil {
		os.Create(Path)
	}
}
