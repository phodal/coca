package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cmd/config"
	"github.com/phodal/coca/core/context/rcall"
	"github.com/phodal/coca/core/infrastructure/coca_file"
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
		dependence := *&reverseConfig.DependencePath
		className := *&reverseConfig.ClassName
		remove := *&reverseConfig.RemovePackage

		if className == "" {
			log.Fatal("lost ClassName")
		}

		analyser := rcall.NewRCallGraph()
		file := coca_file.ReadFile(dependence)
		if file == nil {
			log.Fatal("lost file:" + dependence)
		}

		_ = json.Unmarshal(file, &parsedDeps)

		fmt.Println("start rcall class :", className)
		content := analyser.Analysis(className, *&parsedDeps)

		if remove != "" {
			content = strings.ReplaceAll(content, remove, "")
		}

		coca_file.WriteToCocaFile("rcall.dot", content)
		cmd_util.ConvertToSvg("call")
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)

	reverseCmd.PersistentFlags().StringVarP(&reverseConfig.RemovePackage, "remove", "r", "", "remove package name")
	reverseCmd.PersistentFlags().StringVarP(&reverseConfig.ClassName, "className", "c", "", "path")
	reverseCmd.PersistentFlags().StringVarP(&reverseConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
