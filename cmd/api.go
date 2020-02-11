package cmd

import (
	"encoding/json"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cmd/config"
	"github.com/phodal/coca/pkg/application/api"
	"github.com/phodal/coca/pkg/application/call"
	api_domain2 "github.com/phodal/coca/pkg/domain/api_domain"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
	"regexp"
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
	apiCmdConfig   ApiCmdConfig
	restApis       []api_domain2.RestAPI
	identifiers    []core_domain.CodeDataStruct
	identifiersMap map[string]core_domain.CodeDataStruct
	diMap          map[string]string
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "scan HTTP api from annotation",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		identifiers = cmd_util.LoadIdentify(apiCmdConfig.DependencePath)
		identifiersMap = core_domain.BuildIdentifierMap(identifiers)
		diMap = core_domain.BuildDIMap(identifiers, identifiersMap)

		depPath := apiCmdConfig.DependencePath
		apiPrefix := apiCmdConfig.AggregateApi

		parsedDeps = nil
		depFile := cmd_util.ReadFile(depPath)
		if depFile == nil {
			log.Fatal("lost deps")
		}

		_ = json.Unmarshal(depFile, &parsedDeps)

		if apiCmdConfig.ForceUpdate {
			forceUpdateApi()
		} else {
			apiContent := cmd_util.ReadCocaFile("apis.json")
			if apiContent == nil {
				forceUpdateApi()
			}
			_ = json.Unmarshal(apiContent, &restApis)
		}

		parsedDeps := cmd_util.GetDepsFromJson(depPath)

		filterAPIs := api_domain2.FilterApiByPrefix(apiPrefix, restApis)

		analyser := call.NewCallGraph()
		dotContent, counts := analyser.AnalysisByFiles(filterAPIs, parsedDeps, diMap)

		if apiCmdConfig.Sort {
			api_domain2.SortAPIs(counts)
		}

		if apiCmdConfig.ShowCount {
			table := cmd_util.NewOutput(output)

			table.SetHeader([]string{"Size", "Method", "URI", "Caller"})

			for _, v := range counts {
				table.Append([]string{strconv.Itoa(v.Size), v.HTTPMethod, v.URI, replacePackage(v.Caller)})
			}
			table.Render()
		}

		if apiCmdConfig.RemovePackageNames != "" {
			dotContent = replacePackage(dotContent)
		}

		cmd_util.WriteToCocaFile("api.dot", dotContent)
		cmd_util.ConvertToSvg("api")
	},
}

func forceUpdateApi() {
	app := new(api.JavaApiApp)
	restApis = app.AnalysisPath(filepath.FromSlash(apiCmdConfig.Path), parsedDeps, identifiersMap, diMap)
	cModel, _ := json.MarshalIndent(restApis, "", "\t")
	cmd_util.WriteToCocaFile("apis.json", string(cModel))
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
	apiCmd.PersistentFlags().StringVarP(&apiCmdConfig.RemovePackageNames, "remove", "r", "", "remove package ParamName")
	apiCmd.PersistentFlags().BoolVarP(&apiCmdConfig.ShowCount, "count", "c", false, "count api size")
	apiCmd.PersistentFlags().BoolVarP(&apiCmdConfig.ForceUpdate, "force", "f", false, "force api update")
	apiCmd.PersistentFlags().BoolVarP(&apiCmdConfig.Sort, "sort", "s", false, "sort api")
	apiCmd.PersistentFlags().StringVarP(&apiCmdConfig.AggregateApi, "aggregate", "a", "", "aggregate api")
}
