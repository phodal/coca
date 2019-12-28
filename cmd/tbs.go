package cmd

import (
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/adapter"
	"github.com/phodal/coca/core/domain/tbs"
	"github.com/spf13/cobra"
)

type TbsCmdConfig struct {
	DependencePath string
}

var (
	tbsCmdConfig TbsCmdConfig
)

var tbsCmd = &cobra.Command{
	Use:   "tbs",
	Short: "test bad smell",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		identifiers = adapter.LoadIdentify(apiCmdConfig.DependencePath)
		identifiersMap = adapter.BuildIdentifierMap(identifiers)

		parsedDeps := cmd_util.GetDepsFromJson(tbsCmdConfig.DependencePath)
		app := tbs.NewTbsApp()
		app.AnalysisPath(parsedDeps, identifiersMap)
	},
}

func init() {
	rootCmd.AddCommand(tbsCmd)

	tbsCmd.PersistentFlags().StringVarP(&tbsCmdConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
