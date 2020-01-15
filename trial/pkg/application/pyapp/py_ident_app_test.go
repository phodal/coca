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

	var PyFileFilter = func(path string) bool {
		return strings.HasSuffix(path, ".py")
	}
	files := cocafile.GetFilesWithFilter("testdata/grammar", PyFileFilter)

	for _, file := range files {
		fmt.Println(file)
		file, _ := ioutil.ReadFile(file)
		app := new(PythonApiApp)
		app.Analysis(string(file), "")
	}

	g.Expect(1).To(Equal(1))
}

func Test_PythonClass(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonApiApp)

	file, _ := ioutil.ReadFile("testdata/grammar/class.py")
	codeFile := app.Analysis(string(file), "testdata/grammar/class.py")

	g.Expect(len(codeFile.DataStructures)).To(Equal(1))
}
