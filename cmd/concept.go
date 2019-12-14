package cmd

import (
	"coca/src/adapter/models"
	. "coca/src/domain"
	. "coca/src/utils"
	"encoding/json"
	"github.com/spf13/cobra"
	"log"
)

var parsedDeps []models.JClassNode

var conceptCmd *cobra.Command = &cobra.Command{
	Use:   "concept",
	Short: "concept api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		dependence := cmd.Flag("dependence").Value.String()

		if path != "" {
			analyser := NewConceptAnalyser()
			file := ReadFile(dependence)
			if file == nil {
				log.Fatal("lost file:" + dependence)
			}

 			_ = json.Unmarshal(file, &parsedDeps)

			analyser.Analysis(path, &parsedDeps)
		}
	},
}

func init() {
	rootCmd.AddCommand(conceptCmd)

	conceptCmd.PersistentFlags().StringP("path", "p", "", "path")
	conceptCmd.PersistentFlags().StringP("dependence", "d", "", "get dependence file")
}
