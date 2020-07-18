package cmd

import (
	"encoding/json"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cmd/config"
	. "github.com/phodal/coca/pkg/application/call"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

type CallCmdConfig struct {
	Path       string
	ClassName  string
	RemoveName string
	Lookup     bool
}

var (
	callCmdConfig CallCmdConfig
)

var callGraphCmd = &cobra.Command{
	Use:   "call",
	Short: "show call graph with specific method",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var parsedDeps []core_domain.CodeDataStruct
		dependence := callCmdConfig.Path

		if dependence != "" {
			analyser := NewCallGraph()
			file := cmd_util.ReadFile(dependence)
			if file == nil {
				log.Fatal("lost file:" + dependence)
			}

			_ = json.Unmarshal(file, &parsedDeps)

			content := analyser.Analysis(callCmdConfig.ClassName, parsedDeps, callCmdConfig.Lookup)
			if callCmdConfig.RemoveName != "" {
				content = strings.ReplaceAll(content, callCmdConfig.RemoveName, "")
			}

			cmd_util.WriteToCocaFile("call.dot", content)
			cmd_util.ConvertToSvg("call")
		}
	},
}

func init() {
	rootCmd.AddCommand(callGraphCmd)

	callGraphCmd.PersistentFlags().StringVarP(&callCmdConfig.ClassName, "className", "c", "", "class")
	callGraphCmd.PersistentFlags().StringVarP(&callCmdConfig.Path, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
	callGraphCmd.PersistentFlags().StringVarP(&callCmdConfig.RemoveName, "remove", "r", "", "remove package ParamName")
	callGraphCmd.PersistentFlags().BoolVarP(&callCmdConfig.Lookup, "lookup", "l", false, "call with rcall")
}
