package cmd

import (
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/domain/concept"
	"github.com/phodal/coca/core/models"
	. "github.com/phodal/coca/core/support"
	"encoding/json"
	"fmt"
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

			wordCounts := analyser.Analysis(&parsedDeps)
			for _, word := range wordCounts {
				if word.Value > 0 {
					fmt.Println(word.Key, word.Value)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(conceptCmd)

	conceptCmd.PersistentFlags().StringP("dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
