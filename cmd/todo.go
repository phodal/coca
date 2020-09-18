package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/application/todo"
	"github.com/spf13/cobra"
	"strconv"
)

type RootCmdConfig struct {
	Path       string
	WithGit    bool
	Extensions string
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
		filters := []string{".go", ".py", ".js", ".ts", ".java", ".kotlin", ".groovy"}
		todos := app.AnalysisPath(path, filters)

		simple, _ := json.MarshalIndent(todos, "", "\t")
		cmd_util.WriteToCocaFile("simple-todos.json", string(simple))

		fmt.Fprintf(output, "Todos Count %d\n", len(todos))

		if todoCmdConfig.WithGit {
			gitTodos := app.BuildWithGitHistory(todos)

			cModel, _ := json.MarshalIndent(todos, "", "\t")
			cmd_util.WriteToCocaFile("todos.json", string(cModel))

			table := cmd_util.NewOutput(output)
			table.SetHeader([]string{"Date", "Author", "Messages", "FileName", "Line"})
			for _, todo := range gitTodos {
				table.Append([]string{todo.Date, todo.Author, todo.Message, todo.FileName, todo.Line})
			}

			table.Render()
		} else {
			table := cmd_util.NewOutput(output)
			table.SetHeader([]string{"Filename", "Messages", "Assignee", "Line"})
			for _, todo := range todos {
				table.Append([]string{todo.Filename, todo.Message, todo.Assignee, strconv.Itoa(todo.Line)})
			}

			table.Render()
		}
	},
}

func init() {
	todoCmd.SetOut(output)
	todoCmd.PersistentFlags().StringVarP(&todoCmdConfig.Extensions, "ext", "e", ".java,.py,.go,.ts,.js", "ext=\".java,.go\"")
	todoCmd.PersistentFlags().StringVarP(&todoCmdConfig.Path, "path", "p", ".", "path")
	todoCmd.PersistentFlags().BoolVarP(&todoCmdConfig.WithGit, "git", "g", false, "is with git info")

	rootCmd.AddCommand(todoCmd)
}
