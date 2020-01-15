package pyapp

import (
	"fmt"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/trial"
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

	defs, _ := ioutil.ReadFile("testdata/grammar/classdef.py")
	results := app.Analysis(string(defs), "testdata/grammar/classdef.py")
	g.Expect(len(results.DataStructures)).To(Equal(3))
}

func Test_PythonClassWithDecorator(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonApiApp)

	file, _ := ioutil.ReadFile("testdata/grammar/class_or_func_def_stmt.py")
	codeFile := app.Analysis(string(file), "testdata/grammar/class_or_func_def_stmt.py")

	g.Expect(len(codeFile.DataStructures)).To(Equal(1))
	g.Expect(len(codeFile.DataStructures[0].Annotations.([]trial.PythonAnnotation))).To(Equal(1))

	g.Expect(codeFile.Members[0].MethodNodes[0].Name).To(Equal("bar"))
	g.Expect(len(codeFile.Members[0].MethodNodes[0].Annotations.([]trial.PythonAnnotation))).To(Equal(2))
}
