package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"

	. "coca/core/adapter/call"
	. "coca/core/adapter/identifier"
	. "coca/core/support"
)

type AnalysisCmdConfig struct {
	Path string
}

var (
	analysisCmdConfig AnalysisCmdConfig
)

var analysisCmd = &cobra.Command{
	Use:   "analysis",
	Short: "analysis package",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		importPath := *&analysisCmdConfig.Path

		if importPath != "" {
			identifierApp := new(JavaIdentifierApp)
			iNodes := identifierApp.AnalysisPath(importPath)

			identModel, _ := json.MarshalIndent(iNodes, "", "\t")
			WriteToCocaFile("identify.json", string(identModel))

			var classes []string = nil

			for _, node := range iNodes {
				classes = append(classes, node.Package+"."+node.ClassName)
			}

			callApp := new(JavaCallApp)

			callNodes := callApp.AnalysisPath(importPath, classes, iNodes)
			cModel, _ := json.MarshalIndent(callNodes, "", "\t")
			WriteToCocaFile("deps.json", string(cModel))
		}
	},
}

func init() {
	rootCmd.AddCommand(analysisCmd)

	analysisCmd.PersistentFlags().StringVarP(&analysisCmdConfig.Path, "path", "p", ".", "example -p core/main")
}
