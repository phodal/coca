package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/phodal/coca/core/domain/todo"
	"github.com/phodal/coca/core/support"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type RootCmdConfig struct {
	Path    string
	WithGit bool
}

var (
	todoCmdConfig RootCmdConfig
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "scan all todo, and list with time",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		app := todo.NewTodoApp()
		todos := app.AnalysisPath(path)

		simple, _ := json.MarshalIndent(todos, "", "\t")
		support.WriteToCocaFile("simple-todos.json", string(simple))

		fmt.Println("Todos Count", len(todos))

		if todoCmdConfig.WithGit {
			gitTodos := app.BuildWithGitHistory(todos)

			cModel, _ := json.MarshalIndent(todos, "", "\t")
			support.WriteToCocaFile("todos.json", string(cModel))

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Date", "Author", "Messages", "FileName", "Line"})
			for _, todo := range gitTodos {
				table.Append([]string{todo.Date, todo.Author, strings.Join(todo.Message, "\n"), todo.FileName, todo.Line})
			}

			table.Render()
		}
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)

	todoCmd.PersistentFlags().StringVarP(&todoCmdConfig.Path, "path", "p", ".", "path")
	todoCmd.PersistentFlags().BoolVarP(&todoCmdConfig.WithGit, "git", "g", false, "path")
}
