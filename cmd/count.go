package cmd

import (
	"github.com/spf13/cobra"
)

type CountCmdConfig struct {
	Refs bool
}

var (
	countCmdConfig CountCmdConfig
)

var countCmd *cobra.Command = &cobra.Command{
	Use:   "count",
	Short: "count code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if countCmdConfig.Refs {

		}
	},
}

func init() {
	rootCmd.AddCommand(countCmd)

	countCmd.PersistentFlags().BoolVarP(&countCmdConfig.Refs, "refs", "s", false, "count refs")
}
