package cmd

import (
	"coca/config"
	"coca/core/domain/evaluate"
	. "coca/core/support"
	"encoding/json"
	"github.com/spf13/cobra"
	"log"
)

var evaluateCmd *cobra.Command = &cobra.Command{
	Use:   "concept",
	Short: "concept api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dependence := cmd.Flag("dependence").Value.String()

		if dependence != "" {
			analyser := evaluate.NewEvaluateAnalyser()
			file := ReadFile(dependence)
			if file == nil {
				log.Fatal("lost file:" + dependence)
			}

			_ = json.Unmarshal(file, &parsedDeps)

			analyser.Analysis(&parsedDeps)
		}
	},
}

func init() {
	rootCmd.AddCommand(evaluateCmd)

	evaluateCmd.PersistentFlags().StringP("dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
