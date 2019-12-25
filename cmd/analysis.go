package cmd

import (
	"encoding/json"
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/domain/suggest"
	"github.com/phodal/coca/core/support"
	"github.com/spf13/cobra"
	"log"
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

		parsedDeps = nil
		depFile := support.ReadFile(apiCmdConfig.DependencePath)
		if depFile == nil {
			log.Fatal("lost deps")
		}

		_ = json.Unmarshal(depFile, &parsedDeps)

		if importPath != "" {
			app := suggest.NewSuggestApp()
			 app.AnalysisPath(parsedDeps)
		}
	},
}

func init() {
	rootCmd.AddCommand(analysisCmd)

	apiCmd.PersistentFlags().StringVarP(&apiCmdConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
