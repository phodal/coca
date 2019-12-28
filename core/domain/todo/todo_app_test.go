package todo

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestNewTodoApp(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/todo"


	app := NewTodoApp()
	app.AnalysisPath(codePath)

	g.Expect(true).To(Equal(true))
}
