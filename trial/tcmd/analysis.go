package tcmd

import (
	"encoding/json"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain"
	"github.com/phodal/coca/trial/pkg/application/ts"
	"github.com/spf13/cobra"
	"io/ioutil"
)

type TrialAnalysisCmdConfig struct {
	Path string
}

var (
	analysisCmdConfig TrialAnalysisCmdConfig
)

var analysisCmd = &cobra.Command{
	Use:   "analysis",
	Short: "analysis code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		importPath := analysisCmdConfig.Path

		var results []domain.CodeFile
		files := cocafile.GetFilesWithFilter(importPath, cocafile.TypeScriptFileFilter)
		app := new(ts.TypeScriptApiApp)
		for _, file := range files {
			bytes, _ := ioutil.ReadFile(file)
			result := app.Analysis(string(bytes), file)
			results = append(results, result)
		}

		cModel, _ := json.MarshalIndent(results, "", "\t")
		cmd_util.WriteToCocaFile("tsdeps.json", string(cModel))
	},
}

func init() {
	trialRootCmd.AddCommand(analysisCmd)

	analysisCmd.PersistentFlags().StringVarP(&analysisCmdConfig.Path, "path", "p", ".", "example -p core/main")
}