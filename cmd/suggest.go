package cmd

import (
	"github.com/olekukonko/tablewriter"
	"github.com/phodal/coca/core/domain/todo"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "simple holmes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		if path != "" {
			app := todo.NewTodoApp()
			todos := app.AnalysisPath(path)

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Date", "Author", "Messages", "FileName", "Line"})
			for _, todo := range todos {
				table.Append([]string{todo.Date, todo.Author, strings.Join(todo.Message, "\n"), todo.FileName, todo.Line})
			}

			table.Render()
		}
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)

	suggestCmd.PersistentFlags().StringVarP(&rootCmdConfig.Path, "path", "p", ".", "path")
}
