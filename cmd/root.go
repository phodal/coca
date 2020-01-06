package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	output = os.Stdout
	rootCmd = &cobra.Command{
		Use:   "coca",
		Short: "A generator for Cobra based Applications",
		Long:  `coca`,
	}
)

func Execute() error {
	rootCmd.SetOut(output)
	return rootCmd.Execute()
}
