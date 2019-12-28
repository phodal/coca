package cmd

import (
	"fmt"
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/domain/arch"
	"github.com/spf13/cobra"
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
		parsedDeps := GetDepsFromJson(archCmdConfig.DependencePath)
		archApp := arch.NewArchApp()
		dotContent := archApp.Analysis(parsedDeps)

		fmt.Println(dotContent)
		//ConvertToSvg(dotContent)
	},
}

func init() {
	rootCmd.AddCommand(archCmd)

	archCmd.PersistentFlags().StringVarP(&archCmdConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
