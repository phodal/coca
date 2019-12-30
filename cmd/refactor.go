package cmd

import (
	"encoding/json"
	"github.com/phodal/coca/config"
	. "github.com/phodal/coca/core/domain/refactor/move_class"
	. "github.com/phodal/coca/core/domain/refactor/rename"
	. "github.com/phodal/coca/core/domain/refactor/unused"
	"github.com/phodal/coca/core/infrastructure"
	"github.com/spf13/cobra"
)

var refactorCmd = &cobra.Command{
	Use:   "refactor",
	Short: "auto refactor code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		moveConfig := cmd.Flag("move").Value.String()
		path := cmd.Flag("path").Value.String()
		rename := cmd.Flag("rename").Value.String()
		dependence := cmd.Flag("dependence").Value.String()

		if moveConfig != "" && path != "" {
			app := NewMoveClassApp(moveConfig, path)
			app.Analysis()

			app2 := NewRemoveUnusedImportApp(path)
			results := app2.Analysis()
			app2.Refactoring(results)
		}

		if dependence != "" && rename != "" {
			file := infrastructure.ReadFile(dependence)
			if file == nil {
				return
			}

			_ = json.Unmarshal(file, &parsedDeps)


			renameApp := RenameMethodApp(parsedDeps, rename)
			renameApp.Start()
		}
	},
}

func init() {
	rootCmd.AddCommand(refactorCmd)

	refactorCmd.PersistentFlags().StringP("path", "p", "", "path")
	refactorCmd.PersistentFlags().StringP("move", "m", "", "with config example -m config.file")
	refactorCmd.PersistentFlags().StringP("rename", "R", "", "rename method -R config.file")
	refactorCmd.PersistentFlags().StringP("dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence D")
}
