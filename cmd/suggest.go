package cmd

import (
	"encoding/json"
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/domain/suggest"
	"github.com/phodal/coca/core/support"
	"github.com/spf13/cobra"
	"log"
)

var (
	suggestConfig ApiCmdConfig
)

var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "simple holmes",
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
	rootCmd.AddCommand(suggestCmd)

	suggestCmd.PersistentFlags().StringVarP(&suggestConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence D")
}
