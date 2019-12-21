package cmd

import (
	"coca/config"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version number of Coca",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Coca Version: " + config.VERSION + " -- HEAD")
	},
}