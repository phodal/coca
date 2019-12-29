package cmd

import (
	"github.com/phodal/coca/config"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Coca Version: " + config.VERSION + " -- HEAD")
	},
}