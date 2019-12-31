package cmd

import (
	"encoding/json"
	bs2 "github.com/phodal/coca/core/context/bs"
	"github.com/phodal/coca/core/domain/bs_domain"
	"github.com/phodal/coca/core/infrastructure/coca_file"
	"github.com/phodal/coca/core/infrastructure/string_helper"
	"github.com/spf13/cobra"
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

		bsApp := *bs2.NewBadSmellApp()
		bsList := bsApp.AnalysisPath(importPath, ignoreRules)

		bsModel, _ := json.MarshalIndent(bsList, "", "\t")

		if sortType == "type" {
			sortSmells := bs_domain.SortSmellByType(bsList, isSmellHaveSize)
			bsModel, _ = json.MarshalIndent(sortSmells, "", "\t")
		}

		coca_file.WriteToCocaFile("bs.json", string(bsModel))
	},
}


func isSmellHaveSize(key string) bool {
	var smellList = []string{
		"largeClass",
		"repeatedSwitches",
		"longParameterList",
		"longMethod",
		"dataClass",
	}
	return string_helper.StringArrayContains(smellList, key)
}

func init() {
	rootCmd.AddCommand(badsmellCmd)

	badsmellCmd.PersistentFlags().StringVarP(&bsCmdConfig.Path, "path", "p", ".", "example -p core/main")
	badsmellCmd.PersistentFlags().StringP("ignore", "x", "", "-x=dataClass,lazyElement,longMethod,refusedBequest")
	badsmellCmd.PersistentFlags().StringP("sort", "s", "", "sort bad smell -s=type")
}
