package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "coca",
		Short: "A generator for Cobra based Applications",
		Long:  `coca`,
	}
)

// Execute executes the root command.
func Execute() error {
	fmt.Println(rootCmd.Flag("config").Value.String())
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(collCmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
}

