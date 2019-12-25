package cmd

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

type CountCmdConfig struct {
	DependencePath string
	Top            int
}

var (
	countCmdConfig CountCmdConfig
)

var cparsedDeps []models.JClassNode

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "count code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dependence := countCmdConfig.DependencePath
		if dependence == "" {
			return
		}

		file := support.ReadFile(dependence)
		if file == nil {
			log.Fatal("lost file:" + dependence)
		}

		_ = json.Unmarshal(file, &cparsedDeps)

		var projectMethods = make(map[string]bool)
		for _, clz := range cparsedDeps {
			for _, method := range clz.Methods {
				projectMethods[clz.Package+"."+clz.Class+"."+method.Name] = true
			}
		}

		// TODO: support identify data class
		var callMap = make(map[string]int)
		for _, clz := range cparsedDeps {
			for _, call := range clz.MethodCalls {
				callMethod := call.Package + "." + call.Class + "." + call.MethodName
				if projectMethods[callMethod] {
					if callMap[callMethod] == 0 {
						callMap[callMethod] = 1
					} else {
						callMap[callMethod]++
					}
				}
			}
		}

		callMapSort := support.RankByWordCount(callMap)

		if *&countCmdConfig.Top > 0 {
			callMapSort = callMapSort[:*&countCmdConfig.Top]
		}

		table := tablewriter.NewWriter(os.Stdout)
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
