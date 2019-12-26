package cmd

import (
	"encoding/json"
	"fmt"
	. "github.com/phodal/coca/core/adapter/call"
	. "github.com/phodal/coca/core/adapter/identifier"
	"github.com/phodal/coca/core/models"
	. "github.com/phodal/coca/core/support"
	"github.com/spf13/cobra"
)

type AnalysisCmdConfig struct {
	Path        string
	ForceUpdate bool
}

var (
	analysisCmdConfig AnalysisCmdConfig
)

var analysisCmd = &cobra.Command{
	Use:   "analysis",
	Short: "analysis package",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		importPath := analysisCmdConfig.Path

		if importPath != "" {
			var iNodes []models.JIdentifier
			if IsExistCocaFile("identify.json") && !analysisCmdConfig.ForceUpdate {
				fmt.Println("use exist identify")
				identContent := ReadCocaFile("identify.json")
				_ = json.Unmarshal(identContent, &iNodes)
			} else {
				fmt.Println("force update identify")
				identifierApp := NewJavaIdentifierApp()
				iNodes := identifierApp.AnalysisPath(importPath)

				identModel, _ := json.MarshalIndent(iNodes, "", "\t")
				WriteToCocaFile("identify.json", string(identModel))
			}

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
	analysisCmd.PersistentFlags().BoolVarP(&analysisCmdConfig.ForceUpdate, "force", "f", false, "force update -f")
}
