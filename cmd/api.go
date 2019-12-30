package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/adapter"
	. "github.com/phodal/coca/core/adapter/api"
	"github.com/phodal/coca/core/context/call_graph"
	. "github.com/phodal/coca/core/infrastructure"
	"github.com/spf13/cobra"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type ApiCmdConfig struct {
	Path               string
	DependencePath     string
	ShowCount          bool
	RemovePackageNames string
	AggregateApi       string
	ForceUpdate        bool
	Sort               bool
}

var (
	apiCmdConfig ApiCmdConfig
	restApis     []RestApi

	identifiers    = adapter.LoadIdentify(apiCmdConfig.DependencePath)
	identifiersMap = adapter.BuildIdentifierMap(identifiers)
	diMap          = adapter.BuildDIMap(identifiers, identifiersMap)
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "scan HTTP api from annotation",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		depPath := apiCmdConfig.DependencePath
		apiPrefix := apiCmdConfig.AggregateApi

		parsedDeps = nil
		depFile := ReadFile(depPath)
		if depFile == nil {
			log.Fatal("lost deps")
		}

		_ = json.Unmarshal(depFile, &parsedDeps)

		if *&apiCmdConfig.ForceUpdate {
			forceUpdateApi()
		} else {
			apiContent := ReadCocaFile("apis.json")
			if apiContent == nil {
				forceUpdateApi()
			}
			_ = json.Unmarshal(apiContent, &restApis)
		}

		parsedDeps := cmd_util.GetDepsFromJson(depPath)

		restFieldsApi := filterApi(apiPrefix, restApis)

		fmt.Println(restApis)
		analyser := call_graph.NewCallGraph()
		dotContent, counts := analyser.AnalysisByFiles(restFieldsApi, parsedDeps, diMap)

		if *&apiCmdConfig.Sort {
			sort.Slice(counts, func(i, j int) bool {
				return counts[i].Size < counts[j].Size
			})
		}

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

		WriteToCocaFile("api.dot", dotContent)
		cmd_util.ConvertToSvg("api")
	},
}

func forceUpdateApi() {
	app := new(JavaApiApp)
	restApis = app.AnalysisPath(apiCmdConfig.Path, parsedDeps, identifiersMap, diMap)
	cModel, _ := json.MarshalIndent(restApis, "", "\t")
	WriteToCocaFile("apis.json", string(cModel))
}

func filterApi(apiPrefix string, apis []RestApi, ) []RestApi {
	var restFieldsApi []RestApi
	if apiPrefix != "" {
		for _, api := range apis {
			if strings.HasPrefix(api.Uri, apiPrefix) {
				restFieldsApi = append(restFieldsApi, api)
			}
		}
	} else {
		restFieldsApi = apis
	}

	return restFieldsApi
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
	apiCmd.PersistentFlags().BoolVarP(&apiCmdConfig.ForceUpdate, "force", "f", false, "force api update")
	apiCmd.PersistentFlags().BoolVarP(&apiCmdConfig.Sort, "sort", "s", false, "sort api")
	apiCmd.PersistentFlags().StringVarP(&apiCmdConfig.AggregateApi, "aggregate", "a", "", "aggregate api")
}
