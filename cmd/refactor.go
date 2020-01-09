package cmd

import (
	"encoding/json"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cmd/config"
	. "github.com/phodal/coca/pkg/application/refactor/moveclass"
	. "github.com/phodal/coca/pkg/application/refactor/rename"
	. "github.com/phodal/coca/pkg/application/refactor/unused"
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
			file := cmd_util.ReadFile(dependence)
			if file == nil {
				return
			}

			_ = json.Unmarshal(file, &parsedDeps)

			renameApp := RenameMethodApp(parsedDeps)

			configBytes := cmd_util.ReadFile(rename)
			if configBytes == nil {
				return
			}

			conf := string(configBytes)
			renameApp.Refactoring(conf)
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
