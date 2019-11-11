package cmd

import (
	. "github.com/phodal/coca/adapter/sql"
	"github.com/spf13/cobra"
)

var sqlCmd *cobra.Command = &cobra.Command{
	Use:   "sql",
	Short: "scan sql",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		if path != "" {
			app := new(SqlIdentifierApp)
			app.AnalysisPath(path)
		}
	},
}

func init() {
	rootCmd.AddCommand(sqlCmd)

	sqlCmd.PersistentFlags().StringP("path", "p", "", "path")
}
