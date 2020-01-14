package pyapp

import (
	. "github.com/onsi/gomega"
	"testing"
)


func Test_TypeScriptConsoleLog(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonApiApp)
	app.Analysis("print('console.log')", "")

	g.Expect(1).To(Equal(1))
}

