package cmd

import (
	"encoding/json"
	"github.com/modernizing/coca/cmd/cmd_util"
	"github.com/modernizing/coca/cmd/config"
	"github.com/modernizing/coca/pkg/application/count"
	"github.com/modernizing/coca/pkg/domain/core_domain"
	"github.com/modernizing/coca/pkg/infrastructure/string_helper"
	"github.com/spf13/cobra"
	"strconv"
)

type CountCmdConfig struct {
	DependencePath string
	Top            int
}

var (
	countCmdConfig CountCmdConfig
)

var cparsedDeps []core_domain.CodeDataStruct

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "count most refs function",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dependence := countCmdConfig.DependencePath
		file := cmd_util.ReadFile(dependence)

		_ = json.Unmarshal(file, &cparsedDeps)

		callMap := count.BuildCallMap(cparsedDeps)

		callMapSort := string_helper.SortWord(callMap)

		if countCmdConfig.Top > 0 {
			callMapSort = callMapSort[:countCmdConfig.Top]
		}

		table := cmd_util.NewOutput(output)
		table.SetHeader([]string{"Refs Count", "method"})

		for _, count := range callMapSort {
			table.Append([]string{strconv.Itoa(count.Value), count.Key})
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(countCmd)

	countCmd.PersistentFlags().StringVarP(&countCmdConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
	countCmd.PersistentFlags().IntVarP(&countCmdConfig.Top, "top", "t", 0, "top x")
}
