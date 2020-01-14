package pyapp

import (
	. "github.com/onsi/gomega"
	"io/ioutil"
	"testing"
)


func Test_TypeScriptConsoleLog(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonApiApp)

	file, _ := ioutil.ReadFile("testdata/grammar/class.py")
	app.Analysis(string(file), "")

	g.Expect(1).To(Equal(1))
}

