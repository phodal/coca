package app

import (
	"encoding/json"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/application/analysis/javaapp"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/spf13/cobra"
)

type AnalysisCmdConfig struct {
	Path        string
	ForceUpdate bool
	Lang        string
}

var (
	analysisCmdConfig AnalysisCmdConfig
)

var analysisCmd = &cobra.Command{
	Use:   "analysis",
	Short: "analysis code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var outputName string
		var ds []core_domain.CodeDataStruct
		ds = AnalysisJava()
		outputName = "deps.json"

		cModel, _ := json.MarshalIndent(ds, "", "\t")
		cmd_util.WriteToCocaFile(outputName, string(cModel))
	},
}

func AnalysisJava() []core_domain.CodeDataStruct {
	importPath := analysisCmdConfig.Path
	identifierApp := javaapp.NewJavaIdentifierApp()
	iNodes := identifierApp.AnalysisPath(importPath)

	identModel, _ := json.MarshalIndent(iNodes, "", "\t")
	cmd_util.WriteToCocaFile("identify.json", string(identModel))

	callApp := javaapp.NewJavaFullApp()

	callNodes := callApp.AnalysisPath(importPath, iNodes)
	return callNodes
}

func init() {
	rootCmd.AddCommand(analysisCmd)

	analysisCmd.PersistentFlags().StringVarP(&analysisCmdConfig.Path, "path", "p", ".", "example -p core/main")
	analysisCmd.PersistentFlags().StringVarP(&analysisCmdConfig.Lang, "lang", "l", "java", "example coca analysis -l java, typescript, python")
	analysisCmd.PersistentFlags().BoolVarP(&analysisCmdConfig.ForceUpdate, "force", "f", false, "force update -f")
}
