package cmd

import (
	"coca/config"
	. "coca/core/adapter/api"
	"coca/core/domain/call_graph"
	"coca/core/models"
	. "coca/core/support"
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
	Path               string
	DependencePath     string
	ShowCount          bool
	RemovePackageNames string
}

var (
	apiCmdConfig ApiCmdConfig
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "scan api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := *&apiCmdConfig.Path
		dependence := *&apiCmdConfig.DependencePath

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

			analyser := call_graph.NewCallGraph()
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

			cmd := exec.Command("dot", []string{"-Tsvg", config.CocaConfig.ReporterPath + "/api.dot", "-o", config.CocaConfig.ReporterPath + "/api.svg"}...)
			_, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalf("cmd.Run() failed with %s\n", err)
			}
		}
	},
}

func replacePackage(content string) string {
	var packageRegex string
	packageNameArray := strings.Split(apiCmdConfig.RemovePackageNames, ",")
	for index, name := range packageNameArray {
		packageRegex = packageRegex + strings.ReplaceAll(name, ".", "\\.")
		if index != len(packageNameArray)-1 {
			packageRegex = packageRegex + "|"
		}
	}

	re, _ := regexp.Compile(packageRegex)

	return re.ReplaceAllString(content, "")
}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.PersistentFlags().StringVarP(&apiCmdConfig.Path, "path", "p", ".", "path")
	apiCmd.PersistentFlags().StringVarP(&apiCmdConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
	apiCmd.PersistentFlags().StringVarP(&apiCmdConfig.RemovePackageNames, "remove", "r", "", "remove package name")
	apiCmd.PersistentFlags().BoolVarP(&apiCmdConfig.ShowCount, "count", "c", false, "count api size")
}
