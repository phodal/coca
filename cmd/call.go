package cmd

import (
	"coca/config"
	. "coca/core/domain/call_graph"
	"coca/core/models"
	. "coca/core/support"
	"encoding/json"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"strings"
)

type CallCmdConfig struct {
	Path string
}

var (
	callCmdConfig CallCmdConfig
)

var callGraphCmd = &cobra.Command{
	Use:   "call",
	Short: "call graph api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var parsedDeps []models.JClassNode
		dependence := *&callCmdConfig.Path

		className := cmd.Flag("className").Value.String()
		remove := cmd.Flag("remove").Value.String()

		if dependence != "" {
			analyser := NewCallGraph()
			file := ReadFile(dependence)
			if file == nil {
				log.Fatal("lost file:" + dependence)
			}

			_ = json.Unmarshal(file, &parsedDeps)

			content := analyser.Analysis(className, *&parsedDeps)
			if remove != "" {
				content = strings.ReplaceAll(content, remove, "")
			}

			WriteToCocaFile("call.dot", content)

			cmd := exec.Command("dot", []string{"-Tsvg", config.CocaConfig.ReporterPath + "/call.dot", "-o", config.CocaConfig.ReporterPath + "/call.svg"}...)
			_, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalf("cmd.Run() failed with %s\n", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(callGraphCmd)

	callGraphCmd.PersistentFlags().StringP("className", "c", "", "path")
	callGraphCmd.PersistentFlags().StringVarP(&callCmdConfig.Path, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
	callGraphCmd.PersistentFlags().StringP("remove", "r", "", "remove package name")
}
