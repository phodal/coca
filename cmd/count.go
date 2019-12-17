package cmd

import (
	"coca/src/models"
	"coca/src/support"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

type CountCmdConfig struct {
	Refs           bool
	DependencePath string
}

var (
	countCmdConfig CountCmdConfig
)

var cparsedDeps []models.JClassNode

var countCmd *cobra.Command = &cobra.Command{
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

		var callMap = make(map[string]int)
		for _, clz := range cparsedDeps {
			for _, call := range clz.MethodCalls {
				if callMap[call.Package+"."+call.Class+"."+call.MethodName] == 0 {
					callMap[call.Package+"."+call.Class+"."+call.MethodName] = 1
				} else {
					callMap[call.Package+"."+call.Class+"."+call.MethodName]++
				}
			}
		}

		callMapSort := support.RankByWordCount(callMap)

		for _, count := range callMapSort {
			fmt.Println(count.Value, count.Key)
		}
	},
}

func init() {
	rootCmd.AddCommand(countCmd)

	countCmd.PersistentFlags().StringVarP(&countCmdConfig.DependencePath, "dependence", "d", "", "get dependence file")
	//countCmd.PersistentFlags().BoolVarP(&countCmdConfig.Refs, "refs", "s", false, "count refs")
}
