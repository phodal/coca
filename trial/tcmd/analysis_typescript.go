package tcmd

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/trial"
	"github.com/phodal/coca/trial/pkg/application/ts"
	"github.com/spf13/cobra"
	"io/ioutil"
)

type TrialAnalysisCmdConfig struct {
	Path string
}

var (
	analysisTypeScriptCmdConfig TrialAnalysisCmdConfig
)

var analysisTypeScriptCmd = &cobra.Command{
	Use:   "analysis",
	Short: "analysis code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		importPath := analysisTypeScriptCmdConfig.Path

		var results []trial.CodeFile
		files := cocafile.GetFilesWithFilter(importPath, cocafile.TypeScriptFileFilter)
		app := new(ts.TypeScriptApiApp)
		for _, file := range files {
			bytes, _ := ioutil.ReadFile(file)
			fmt.Println("Process parse TypeScript file : " + file)

			result := app.Analysis(string(bytes), file)
			results = append(results, result)
		}

		cModel, _ := json.MarshalIndent(results, "", "\t")
		cmd_util.WriteToCocaFile("tsdeps.json", string(cModel))
	},
}

func init() {
	trialRootCmd.AddCommand(analysisTypeScriptCmd)

	analysisTypeScriptCmd.PersistentFlags().StringVarP(&analysisTypeScriptCmdConfig.Path, "path", "p", ".", "example -p core/main")
}