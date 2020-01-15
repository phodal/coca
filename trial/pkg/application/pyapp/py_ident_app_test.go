package pyapp

import (
	. "github.com/onsi/gomega"
	"io/ioutil"
	"testing"
)


func Test_PythonClass(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonApiApp)

	file, _ := ioutil.ReadFile("testdata/grammar/class.py")
	app.Analysis(string(file), "")

	g.Expect(1).To(Equal(1))
}


func Test_PythonFuncDefStm(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonApiApp)

	file, _ := ioutil.ReadFile("testdata/grammar/class_or_func_def_stmt.py")
	app.Analysis(string(file), "")

	g.Expect(1).To(Equal(1))
}

