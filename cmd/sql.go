package cmd

import (
	. "github.com/phodal/coca/core/adapter/sql"
	"github.com/spf13/cobra"
)

var sqlCmd = &cobra.Command{
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

