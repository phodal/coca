package cmd

import (
	"coca/config"
	"coca/core/domain/call_graph/rcall"
	. "coca/core/support"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
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
	Short: "reverse call",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dependence := *&reverseConfig.DependencePath
		className := *&reverseConfig.ClassName
		remove := *&reverseConfig.RemovePackage

		if className == "" {
			log.Fatal("lost ClassName")
		}

		analyser := rcall.NewRCallGraph()
		file := ReadFile(dependence)
		if file == nil {
			log.Fatal("lost file:" + dependence)
		}

		_ = json.Unmarshal(file, &parsedDeps)

		fmt.Println("start rcall class :", className)
		content := analyser.Analysis(className, *&parsedDeps)

		if remove != "" {
			content = strings.ReplaceAll(content, remove, "")
		}

		WriteToCocaFile("rcall.dot", content)

		acmd := exec.Command("dot", []string{"-Tsvg", config.CocaConfig.ReporterPath + "/rcall.dot", "-o", config.CocaConfig.ReporterPath + "/rcall.svg"}...)
		output, err := acmd.CombinedOutput()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err, string(output))
		}
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)

	reverseCmd.PersistentFlags().StringVarP(&reverseConfig.RemovePackage, "remove", "r", "", "remove package name")
	reverseCmd.PersistentFlags().StringVarP(&reverseConfig.ClassName, "className", "c", "", "path")
	reverseCmd.PersistentFlags().StringVarP(&reverseConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
