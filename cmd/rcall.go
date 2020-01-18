package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cmd/config"
	"github.com/phodal/coca/pkg/application/rcall"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

type ReverseCmdConfig struct {
	DependencePath string
	ClassName      string
	RemovePackage  string
}

var (
	reverseConfig ReverseCmdConfig
)

var reverseCmd = &cobra.Command{
	Use:   "rcall",
	Short: "reverse call graph visualization",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dependence := reverseConfig.DependencePath
		className := reverseConfig.ClassName
		remove := reverseConfig.RemovePackage

		if className == "" {
			log.Fatal("lost NodeName")
		}

		analyser := rcall.NewRCallGraph()
		file := cmd_util.ReadFile(dependence)

		_ = json.Unmarshal(file, &parsedDeps)

		fmt.Fprintf(output,"start rcall class: %s\n", className)
		content := analyser.Analysis(className, parsedDeps, WriteCallMap)

		if remove != "" {
			content = strings.ReplaceAll(content, remove, "")
		}

		cmd_util.WriteToCocaFile("rcall.dot", content)
		cmd_util.ConvertToSvg("rcall")
	},
}

func WriteCallMap(rcallMap map[string][]string) {
	mapJson, _ := json.MarshalIndent(rcallMap, "", "\t")
	cmd_util.WriteToCocaFile("rcallmap.json", string(mapJson))
}

func init() {
	rootCmd.AddCommand(reverseCmd)

	reverseCmd.PersistentFlags().StringVarP(&reverseConfig.RemovePackage, "remove", "r", "", "remove package ParamName")
	reverseCmd.PersistentFlags().StringVarP(&reverseConfig.ClassName, "className", "c", "", "path")
	reverseCmd.PersistentFlags().StringVarP(&reverseConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
