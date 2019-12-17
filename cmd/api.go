package cmd

import (
	. "coca/src/adapter/api"
	"coca/src/adapter/models"
	"coca/src/domain"
	. "coca/src/utils"
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type ApiCmdConfig struct {
	ShowCount          bool
	RemovePackageNames string
}

var (
	apiCmdConfig ApiCmdConfig
)

var apiCmd *cobra.Command = &cobra.Command{
	Use:   "api",
	Short: "scan api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		dependence := cmd.Flag("dependence").Value.String()

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
			dotContent, counts := analyser.AnalysisByFiles(restApis, parsedDeps)

			if apiCmdConfig.ShowCount {
				table := tablewriter.NewWriter(os.Stdout)

				table.SetHeader([]string{"Size", "Method", "Uri", "Caller"})

				for _, v := range counts {
					table.Append([]string{strconv.Itoa(v.Size), v.HttpMethod, v.Uri, replacePackage(v.Caller)})
				}
				table.Render()
			}

			if apiCmdConfig.RemovePackageNames != "" {
				dotContent = replacePackage(dotContent)
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

func replacePackage(content string) string {
	var packagegsReggex string
	packageNameArray := strings.Split(apiCmdConfig.RemovePackageNames, ",")
	for index, name := range packageNameArray {
		packagegsReggex = packagegsReggex + strings.ReplaceAll(name, ".", "\\.")
		if index != len(packageNameArray) - 1 {
			packagegsReggex = packagegsReggex + "|"
		}
	}

	re, _ := regexp.Compile(packagegsReggex)


	return re.ReplaceAllString(content, "")
}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.PersistentFlags().StringP("path", "p", "", "path")
	apiCmd.PersistentFlags().StringP("dependence", "d", "", "get dependence file")
	apiCmd.PersistentFlags().StringVarP(&apiCmdConfig.RemovePackageNames, "remove", "r", "", "remove package name")
	apiCmd.PersistentFlags().BoolVarP(&apiCmdConfig.ShowCount, "count", "c", false, "count api size")
}
