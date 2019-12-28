package cmd

import (
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/adapter"
	"github.com/phodal/coca/core/domain/arch"
	"github.com/spf13/cobra"
	"strings"
)

type ArchCmdConfig struct {
	DependencePath string
}

var (
	archCmdConfig ArchCmdConfig
)

var archCmd = &cobra.Command{
	Use:   "arch",
	Short: "generate arch",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		identifiers = adapter.LoadIdentify(apiCmdConfig.DependencePath)
		identifiersMap = adapter.BuildIdentifierMap(identifiers)

		parsedDeps := cmd_util.GetDepsFromJson(archCmdConfig.DependencePath)
		archApp := arch.NewArchApp()
		dotContent := archApp.Analysis(parsedDeps, identifiersMap)

		ignores := strings.Split("", ",")
		var nodeFilter = func(key string) bool {
			for _, f := range ignores {
				if key == f {
					return true
				}
			}
			return false
		}

		dotContent.ToDot("coca_reporter/arch.dot", ".", nodeFilter)
		cmd_util.ConvertToSvg("arch")
	},
}

func init() {
	rootCmd.AddCommand(archCmd)

	archCmd.PersistentFlags().StringVarP(&archCmdConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
