package cmd

import (
	"github.com/phodal/coca/core/domain/todo"
	"github.com/spf13/cobra"
)

type RootCmdConfig struct {
	Path string
}

var (
	rootCmdConfig RootCmdConfig
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "scan todo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		if path != "" {
			app := todo.NewTodoApp()
			app.AnalysisPath(path)
		}
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)

	todoCmd.PersistentFlags().StringVarP(&rootCmdConfig.Path, "path", "p", ".", "path")
}
