package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"

	. "coca/adapter/call"
	. "coca/adapter/identifier"
	. "coca/utils"
)

var analysisCmd *cobra.Command = &cobra.Command{
	Use:   "analysis",
	Short: "analysis package",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		importPath := cmd.Flag("path").Value.String()

		if importPath != "" {
			identifierApp := new(JavaIdentifierApp)
			iNodes := identifierApp.AnalysisPath(importPath)

			var classes []string = nil

			for _, node := range iNodes {
				classes = append(classes, node.Package + "." + node.Name)
			}

			callApp := new(JavaCallApp)
			callNodes := callApp.AnalysisPath(importPath, classes)

			cModel, _ := json.MarshalIndent(callNodes, "", "\t")
			WriteToFile("deps.json", string(cModel))
		}
	},
}

func init() {
	rootCmd.AddCommand(analysisCmd)

	analysisCmd.PersistentFlags().StringP("path", "p", "Code Path", "example -p src/main")
}
