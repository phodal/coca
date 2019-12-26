package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/domain/evaluate"
	"github.com/phodal/coca/core/models"
	. "github.com/phodal/coca/core/support"
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

		var identifiers []models.JIdentifier
		identContent := ReadCocaFile("identify.json")

		_ = json.Unmarshal(identContent, &identifiers)
		_ = json.Unmarshal(file, &parsedDeps)

		result := analyser.Analysis(parsedDeps, identifiers)

		cModel, _ := json.MarshalIndent(result, "", "\t")
		WriteToCocaFile("evaluate.json", string(cModel))

		fmt.Println(" ----- same type in service ------ ")
		fmt.Println(result.ServiceSummary.ReturnTypeMap)

		fmt.Println("-------- Null -------- Method")
		for _, nullItem := range result.Nullable.Items {
			fmt.Println(nullItem)
		}
	},
}

func init() {
	rootCmd.AddCommand(evaluateCmd)

	evaluateCmd.PersistentFlags().StringVarP(&evaluateConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
