package cmd

import (
	"coca/core/domain/bs"
	"coca/core/support"
	"encoding/json"
	"github.com/spf13/cobra"
	"strings"
)

var badsmellCmd *cobra.Command = &cobra.Command{
	Use:   "bs",
	Short: "bad smell analysis",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		importPath := cmd.Flag("path").Value.String()
		ignoreStr := cmd.Flag("ignore").Value.String()
		sortType := cmd.Flag("sort").Value.String()

		ignoreRules := strings.Split(ignoreStr, ",")

		if importPath != "" {
			bsApp := new(bs.BadSmellApp)
			bsList := bsApp.AnalysisPath(importPath, ignoreRules)

			bsModel, _ := json.MarshalIndent(bsList, "", "\t")

			if sortType == "type" {
				sortSmells := sortSmellByType(bsList)
				bsModel, _ = json.MarshalIndent(sortSmells, "", "\t")
			}

			support.WriteToFile("bs.json", string(bsModel))
		}
	},
}

func sortSmellByType(models []bs.BadSmellModel) map[string][]bs.BadSmellModel {
	sortSmells := make(map[string][]bs.BadSmellModel)
	for _, model := range models {
		sortSmells[model.Bs] = append(sortSmells[model.Bs], model)
	}

	return sortSmells
}

func init() {
	rootCmd.AddCommand(badsmellCmd)

	badsmellCmd.PersistentFlags().StringP("path", "p", "", "example -p core/main")
	badsmellCmd.PersistentFlags().StringP("ignore", "x", "", "-x=dataClass,lazyElement,longMethod,refusedBequest")
	badsmellCmd.PersistentFlags().StringP("sort", "s", "", "sort bad smell -s=type")
}
