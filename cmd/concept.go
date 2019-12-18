package cmd

import (
	"coca/config"
	"coca/core/domain/concept"
	"coca/core/models"
	. "coca/core/support"
	"encoding/json"
	"github.com/spf13/cobra"
	"log"
)

var parsedDeps []models.JClassNode

var conceptCmd = &cobra.Command{
	Use:   "concept",
	Short: "concept api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dependence := cmd.Flag("dependence").Value.String()

		if dependence != "" {
			analyser := concept.NewConceptAnalyser()
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
	rootCmd.AddCommand(conceptCmd)

	conceptCmd.PersistentFlags().StringP("dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
