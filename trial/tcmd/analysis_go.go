package tcmd

import (
	"encoding/json"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/trial"
	"github.com/phodal/coca/trial/cocago"
	"github.com/spf13/cobra"
)

type TrialAnalysisGoCmdConfig struct {
	Path string
}

var (
	analysisGoCmdConfig TrialAnalysisGoCmdConfig
)

var analysisGoCmd = &cobra.Command{
	Use:   "go",
	Short: "analysis code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		importPath := analysisGoCmdConfig.Path

		var results []trial.CodeFile
		files := cocafile.GetFilesWithFilter(importPath, cocafile.GoFileFilter)
		for _, file := range files {
			parser := cocago.NewCocagoParser()
			parser.SetOutput(true)
			result := parser.ProcessFile(file)

			results = append(results, result)
		}

		cModel, _ := json.MarshalIndent(results, "", "\t")
		cmd_util.WriteToCocaFile("godeps.json", string(cModel))
	},
}

func init() {
	trialRootCmd.AddCommand(analysisGoCmd)

	analysisGoCmd.PersistentFlags().StringVarP(&analysisGoCmdConfig.Path, "path", "p", ".", "example -p core/main")
}