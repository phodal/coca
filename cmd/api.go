package cmd

import (
	. "github.com/phodal/coca/adapter/api"
	"github.com/spf13/cobra"
)

var apiCmd *cobra.Command = &cobra.Command{
	Use:   "api",
	Short: "scan api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		dependence := cmd.Flag("dependence").Value.String()

		if path != "" {
			app := new(JavaApiApp)
			app.AnalysisPath(path, dependence)
		}
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	rootCmd.PersistentFlags().StringP("path", "p", "", "path")
	rootCmd.PersistentFlags().StringP("dependence", "d", "", "get dependence file")
}
