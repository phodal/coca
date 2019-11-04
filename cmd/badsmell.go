package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"

	. "github.com/phodal/coca/bs"
	. "github.com/phodal/coca/utils"
)

var badsmellCmd *cobra.Command = &cobra.Command{
	Use:   "badsmell",
	Short: "Bad Code Smell",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		importPath := cmd.Flag("path").Value.String()

		if importPath != "" {
			bsApp := new(BadSmellApp)
			bsList := bsApp.AnalysisPath(importPath)

			bsModel, _ := json.MarshalIndent(bsList, "", "\t")

			WriteToFile("bs.json", string(bsModel))
		}
	},
}

func init() {
	rootCmd.AddCommand(badsmellCmd)

	badsmellCmd.PersistentFlags().StringP("path", "p", "Code Path", "example -p src/main")
}
