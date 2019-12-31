package cmd

import (
	"encoding/json"
	"github.com/phodal/coca/core/ast/bs"
	"github.com/phodal/coca/core/domain/bs_domain"
	"github.com/phodal/coca/core/infrastructure/coca_file"
	"github.com/phodal/coca/core/infrastructure/str_helper"
	"github.com/spf13/cobra"
	"sort"
	"strings"
)

type BsCmdConfig struct {
	Path string
}

var (
	bsCmdConfig BsCmdConfig
)

var badsmellCmd = &cobra.Command{
	Use:   "bs",
	Short: "generate bad smell list and suggestions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		importPath := *&bsCmdConfig.Path
		ignoreStr := cmd.Flag("ignore").Value.String()
		sortType := cmd.Flag("sort").Value.String()

		ignoreRules := strings.Split(ignoreStr, ",")

		bsApp := *bs.NewBadSmellApp()
		bsList := bsApp.AnalysisPath(importPath, ignoreRules)

		bsModel, _ := json.MarshalIndent(bsList, "", "\t")

		if sortType == "type" {
			sortSmells := sortSmellByType(bsList)
			bsModel, _ = json.MarshalIndent(sortSmells, "", "\t")
		}

		coca_file.WriteToCocaFile("bs.json", string(bsModel))
	},
}

func sortSmellByType(models []bs_domain.BadSmellModel) map[string][]bs_domain.BadSmellModel {
	sortSmells := make(map[string][]bs_domain.BadSmellModel)
	for _, model := range models {
		sortSmells[model.Bs] = append(sortSmells[model.Bs], model)
	}

	for key, smells := range sortSmells {
		if isSmellHaveSize(key) {
			sort.Slice(smells, func(i, j int) bool {
				return smells[i].Size > (smells[j].Size)
			})

			sortSmells[key] = smells
		}
	}

	return sortSmells
}

func isSmellHaveSize(key string) bool {
	var smellList = []string{
		"largeClass",
		"repeatedSwitches",
		"longParameterList",
		"longMethod",
		"dataClass",
	}
	return str_helper.StringArrayContains(smellList, key)
}

func init() {
	rootCmd.AddCommand(badsmellCmd)

	badsmellCmd.PersistentFlags().StringVarP(&bsCmdConfig.Path, "path", "p", ".", "example -p core/main")
	badsmellCmd.PersistentFlags().StringP("ignore", "x", "", "-x=dataClass,lazyElement,longMethod,refusedBequest")
	badsmellCmd.PersistentFlags().StringP("sort", "s", "", "sort bad smell -s=type")
}
