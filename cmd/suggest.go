package cmd

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cmd/config"
	"github.com/phodal/coca/core/context/suggest"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	suggestConfig ApiCmdConfig
)

var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "find usable Design Patterns from code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		parsedDeps = nil
		depFile := cmd_util.ReadFile(apiCmdConfig.DependencePath)
		if depFile == nil {
			log.Fatal("lost deps")
		}

		_ = json.Unmarshal(depFile, &parsedDeps)

		app := suggest.NewSuggestApp()
		results := app.AnalysisPath(parsedDeps)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Class", "Pattern", "Reason"})

		for _, result := range results {
			table.Append([]string{result.Class, result.Pattern, result.Reason})
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)

	suggestCmd.PersistentFlags().StringVarP(&suggestConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence D")
}
