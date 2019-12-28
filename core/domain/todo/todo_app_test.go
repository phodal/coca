package todo

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestNewTodoApp(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/todo"


	app := NewTodoApp()
	todos := app.AnalysisPath(codePath)

	g.Expect(todos[0].Line).To(Equal("6"))
	g.Expect(todos[0].Date).To(Equal("2019-12-28"))
	g.Expect(todos[0].FileName).To(ContainSubstring("_fixtures/todo/Toodo.java"))
	g.Expect(todos[0].Author).To(ContainSubstring("Phodal Huang"))
	g.Expect(todos[1].Line).To(Equal("13"))
	g.Expect(todos[1].Message[0]).To(Equal("add more content"))
	g.Expect(todos[1].Assignee).To(Equal("phodal"))
}
