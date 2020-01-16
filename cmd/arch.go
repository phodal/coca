package cmd

import (
	"bufio"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cmd/config"
	"github.com/phodal/coca/pkg/application/arch"
	"github.com/phodal/coca/pkg/application/arch/tequila"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type ArchCmdConfig struct {
	DependencePath string
	IsMergePackage bool
}

var (
	archCmdConfig ArchCmdConfig
)

var archCmd = &cobra.Command{
	Use:   "arch",
	Short: "project package visualization",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		identifiers = cmd_util.LoadIdentify(apiCmdConfig.DependencePath)
		identifiersMap = core_domain.BuildIdentifierMap(identifiers)

		parsedDeps := cmd_util.GetDepsFromJson(archCmdConfig.DependencePath)
		archApp := arch.NewArchApp()
		result := archApp.Analysis(parsedDeps, identifiersMap)

		ignores := strings.Split("", ",")
		var nodeFilter = func(key string) bool {
			for _, f := range ignores {
				if key == f {
					return true
				}
			}
			return false
		}


		if archCmdConfig.IsMergePackage {
			result = result.MergeHeaderFile(tequila.MergePackageFunc)
		}

		graph := result.ToDot(".", nodeFilter)
		f, _ := os.Create("coca_reporter/arch.dot")
		w := bufio.NewWriter(f)
		_, _ = w.WriteString("di" + graph.String())
		_ = w.Flush()

		cmd_util.ConvertToSvg("arch")
	},
}

func init() {
	rootCmd.AddCommand(archCmd)

	archCmd.PersistentFlags().StringVarP(&archCmdConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
	archCmd.PersistentFlags().BoolVarP(&archCmdConfig.IsMergePackage, "mergePackage", "P", false, "merge package/folder for include dependencies")
}
