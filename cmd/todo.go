package cmd

import (
	"github.com/olekukonko/tablewriter"
	"github.com/phodal/coca/core/domain/todo"
	"github.com/spf13/cobra"
	"os"
	"strings"
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
		app := todo.NewTodoApp()
		todos := app.AnalysisPath(path)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Date", "Author", "Messages", "FileName", "Line"})
		for _, todo := range todos {
			table.Append([]string{todo.Date, todo.Author, strings.Join(todo.Message, "\n"), todo.FileName, todo.Line})
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)

	todoCmd.PersistentFlags().StringVarP(&rootCmdConfig.Path, "path", "p", ".", "path")
}
