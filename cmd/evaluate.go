package cmd

import (
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/domain/evaluate"
	. "github.com/phodal/coca/core/support"
	"encoding/json"
	"github.com/spf13/cobra"
	"log"
)

type EvaluateConfig struct {
	DependencePath string
}

var (
	evaluateConfig EvaluateConfig
)

var evaluateCmd = &cobra.Command{
	Use:   "evaluate",
	Short: "evaluate refactor effort",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dependence := *&evaluateConfig.DependencePath

		analyser := evaluate.NewEvaluateAnalyser()
		file := ReadFile(dependence)
		if file == nil {
			log.Fatal("lost file:" + dependence)
		}

		// util

		_ = json.Unmarshal(file, &parsedDeps)

		analyser.Analysis(*&parsedDeps)
	},
}

func init() {
	rootCmd.AddCommand(evaluateCmd)

	evaluateCmd.PersistentFlags().StringVarP(&evaluateConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
