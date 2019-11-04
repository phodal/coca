package cmd

import (
	. "github.com/phodal/coca/bs"
	"github.com/spf13/cobra"
)

var badsmellCmd *cobra.Command = &cobra.Command{
	Use:   "badsmell",
	Short: "badsmell recognized",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		depFile := cmd.Flag("dependence").Value.String()

		if depFile != "" {
			bsApp := NewBadSmellApp(depFile)
			bsApp.Start()
		}
	},
}

func init() {
	rootCmd.AddCommand(badsmellCmd)

	badsmellCmd.PersistentFlags().StringP("dependence", "d", "", "dependence path")
}
