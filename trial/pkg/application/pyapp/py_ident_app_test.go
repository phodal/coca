package pyapp

import (
	"fmt"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"io/ioutil"
	"strings"
	"testing"
)


func Test_AllPythonGrammar(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonApiApp)

	var PyFileFilter = func(path string) bool {
		return strings.HasSuffix(path, ".py")
	}
	files := cocafile.GetFilesWithFilter("testdata/grammar/", PyFileFilter)

	for _, file := range files {
		fmt.Println(file)
		file, _ := ioutil.ReadFile(file)
		app.Analysis(string(file), "")
	}

	g.Expect(1).To(Equal(1))
}

func Test_PythonClass(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonApiApp)

	file, _ := ioutil.ReadFile("testdata/grammar/class.py")
	app.Analysis(string(file), "")

	g.Expect(1).To(Equal(1))
}

func Test_PythonVersion(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonApiApp)

	app.Analysis("print 'a' ", "")
	app.Analysis("print('a')", "")

	g.Expect(1).To(Equal(1))
}

func Test_PythonFuncDefStm(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonApiApp)

	file, _ := ioutil.ReadFile("testdata/grammar/class_or_func_def_stmt.py")
	app.Analysis(string(file), "")

	g.Expect(1).To(Equal(1))
}

