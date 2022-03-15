package todo

import (
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func TestNewTodoApp(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/todo"
	codePath = filepath.FromSlash(codePath)
	app := NewTodoApp()

	stodos := app.AnalysisPath(codePath, []string{".go", ".py", ".js", ".ts", ".java", ".kotlin", ".groovy"})
	todos := app.BuildWithGitHistory(stodos)

	// todo: add support for python
	g.Expect(len(todos)).To(Equal(4))
	g.Expect(todos[0].Line).To(Equal("3"))
	g.Expect(todos[1].FileName).To(ContainSubstring(filepath.FromSlash("_fixtures/todo/Todo.java")))
	g.Expect(todos[1].Author).To(ContainSubstring("Phodal Huang"))
	g.Expect(todos[1].Line).To(Equal("6"))
	g.Expect(todos[2].Message).To(Equal("add more content"))
	g.Expect(todos[2].Assignee).To(Equal("phodal"))
}

func Test_ShouldReturnNullWhenNotTodo(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/tbs"
	codePath = filepath.FromSlash(codePath)
	app := NewTodoApp()

	stodos := app.AnalysisPath(codePath, []string{".go", ".py", ".js", ".ts", ".java", ".kotlin", ".groovy"})
	todos := app.BuildWithGitHistory(stodos)

	g.Expect(len(todos)).To(Equal(0))
}
