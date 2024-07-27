package console

import (
	"golang-rest-api-articles/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tubagus",
	Short: "Tubagus Apps",
	Long:  `Tubagus is a CLI application that can be used to create, read, update, delete data in MySQL databases.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	config.SetupLogger()
}
