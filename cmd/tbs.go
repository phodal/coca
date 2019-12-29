package cmd

import (
	"encoding/json"
	"github.com/phodal/coca/core/adapter"
	"github.com/phodal/coca/core/adapter/call"
	"github.com/phodal/coca/core/domain/tbs"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"github.com/spf13/cobra"
)

type TbsCmdConfig struct {
	Path string
}

var (
	tbsCmdConfig TbsCmdConfig
)

var tbsCmd = &cobra.Command{
	Use:   "tbs",
	Short: "test bad smell",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		files := support.GetJavaTestFiles(tbsCmdConfig.Path)
		var identifiers []models.JIdentifier

		identifiers = adapter.LoadTestIdentify(files)
		identifiersMap := adapter.BuildIdentifierMap(identifiers)

		var classes []string = nil
		for _, node := range identifiers {
			classes = append(classes, node.Package+"."+node.ClassName)
		}

		analysisApp := call.NewJavaCallApp()
		classNodes := analysisApp.AnalysisFiles(identifiers, files, classes)

		nodeContent, _ := json.MarshalIndent(classNodes, "", "\t")
		support.WriteToCocaFile("tdeps.json", string(nodeContent))

		app := tbs.NewTbsApp()
		result := app.AnalysisPath(classNodes, identifiersMap)

		resultContent, _ := json.MarshalIndent(result, "", "\t")
		support.WriteToCocaFile("tbs.json", string(resultContent))
	},
}

func init() {
	rootCmd.AddCommand(tbsCmd)

	tbsCmd.PersistentFlags().StringVarP(&tbsCmdConfig.Path, "path", "p", ".", "example -p core/main")
}
