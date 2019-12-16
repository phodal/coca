package cmd

import (
	. "coca/src/adapter/api"
	"coca/src/adapter/models"
	"coca/src/domain"
	. "coca/src/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"strings"
)

var apiCmd *cobra.Command = &cobra.Command{
	Use:   "api",
	Short: "scan api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		dependence := cmd.Flag("dependence").Value.String()
		remove := cmd.Flag("remove").Value.String()
		isShowApiCount := cmd.Flag("count").Value.String()

		if path != "" {
			app := new(JavaApiApp)
			restApis := app.AnalysisPath(path, dependence)

			cModel, _ := json.MarshalIndent(restApis, "", "\t")
			WriteToFile("apis.json", string(cModel))

			var parsedDeps []models.JClassNode
			file := ReadFile(dependence)
			if file == nil {
				log.Fatal("lost file:" + dependence)
			}
			_ = json.Unmarshal(file, &parsedDeps)

			analyser := domain.NewCallGraph()
			dotContent, countMap := analyser.AnalysisByFiles(restApis, parsedDeps)

			if isShowApiCount != "" {
				for _, count := range countMap {
					fmt.Println(count.Value, count.Key)
				}
			}

			if remove != "" {
				dotContent = strings.ReplaceAll(dotContent, remove, "")
			}

			WriteToFile("api.dot", dotContent)

			cmd := exec.Command("dot", []string{"-Tsvg", "api.dot", "-o", "api.svg"}...)
			_, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalf("cmd.Run() failed with %s\n", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.PersistentFlags().StringP("path", "p", "", "path")
	apiCmd.PersistentFlags().StringP("dependence", "d", "", "get dependence file")
	apiCmd.PersistentFlags().StringP("remove", "r", "", "remove package name")
	apiCmd.PersistentFlags().StringP("count", "c", "", "count api size")
}
