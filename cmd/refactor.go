package cmd

import (
	. "coca/src/domain/refactor/move_class"
	. "coca/src/domain/refactor/rename"
	. "coca/src/domain/refactor/unused"
	"github.com/spf13/cobra"
)

var refactorCmd *cobra.Command = &cobra.Command{
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

			app2 := NewRemoveUnusedImportApp(moveConfig, path)
			app2.Analysis()
		}

		if dependence != "" && rename != "" {
			renameApp := RenameMethodApp(dependence, rename)
			renameApp.Start()
		}
	},
}

func init() {
	rootCmd.AddCommand(refactorCmd)

	refactorCmd.PersistentFlags().StringP("path", "p", "", "path")
	refactorCmd.PersistentFlags().StringP("move", "m", "", "with config example -m config.file")
	refactorCmd.PersistentFlags().StringP("rename", "R", "", "rename method -R config.file")
	refactorCmd.PersistentFlags().StringP("dependence", "D", "", "get dependence D")
}
