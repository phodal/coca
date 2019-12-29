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

	stodos := app.AnalysisPath(codePath)
	todos := app.BuildWithGitHistory(stodos)

	g.Expect(todos[0].Line).To(Equal("6"))
	//g.Expect(todos[0].Date).To(Equal("2019-12-28")) test: will failure in CI
	g.Expect(todos[0].FileName).To(ContainSubstring("_fixtures/todo/Toodo.java"))
	g.Expect(todos[0].Author).To(ContainSubstring("Phodal Huang"))
	g.Expect(todos[1].Line).To(Equal("13"))
	g.Expect(todos[1].Message[0]).To(Equal("add more content"))
	g.Expect(todos[1].Assignee).To(Equal("phodal"))
}
