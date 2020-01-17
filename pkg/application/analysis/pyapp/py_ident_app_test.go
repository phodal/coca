package pyapp

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cocatest"
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
		file, _ := ioutil.ReadFile(file)
		app := new(PythonIdentApp)
		app.Analysis(string(file), "")
	}

	g.Expect(1).To(Equal(1))
}

func Test_PythonClassForLexer(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonIdentApp)

	file, _ := ioutil.ReadFile("testdata/grammar/class.py")
	app.Analysis(string(file), "")

	g.Expect(1).To(Equal(1))
}

func Test_PythonArgumentsForLexer(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonIdentApp)

	file, _ := ioutil.ReadFile("testdata/grammar/argument.py")
	app.Analysis(string(file), "")

	g.Expect(1).To(Equal(1))
}

func Test_PythonArgumentsForTryStmt(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonIdentApp)

	file, _ := ioutil.ReadFile("testdata/grammar/try_stmt.py")
	app.Analysis(string(file), "")

	g.Expect(1).To(Equal(1))
}

func Test_PythonClassDef(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonIdentApp)

	file, _ := ioutil.ReadFile("testdata/grammar/class.py")
	codeFile := app.Analysis(string(file), "testdata/grammar/class.py")

	g.Expect(len(codeFile.DataStructures)).To(Equal(1))

	defs, _ := ioutil.ReadFile("testdata/grammar/classdef.py")
	results := app.Analysis(string(defs), "testdata/grammar/classdef.py")
	g.Expect(len(results.DataStructures)).To(Equal(3))
}

func Test_PythonClassWithDecorator(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonIdentApp)

	file, _ := ioutil.ReadFile("testdata/grammar/class_or_func_def_stmt.py")
	codeFile := app.Analysis(string(file), "testdata/grammar/class_or_func_def_stmt.py")

	g.Expect(len(codeFile.DataStructures)).To(Equal(1))
	g.Expect(len(codeFile.DataStructures[0].Annotations)).To(Equal(1))

	g.Expect(codeFile.Members[0].FunctionNodes[0].Name).To(Equal("bar"))
	g.Expect(len(codeFile.Members[0].FunctionNodes[0].Annotations)).To(Equal(2))
}

func Test_PythonImport(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonIdentApp)

	file, _ := ioutil.ReadFile("testdata/grammar/import_stmt.py")
	codeFile := app.Analysis(string(file), "import_stmt")

	g.Expect(len(codeFile.Imports)).To(Equal(10))
	g.Expect(len(codeFile.Imports[2].UsageName)).To(Equal(2))
	g.Expect(cocatest.JSONFileBytesEqual(codeFile.Imports, "testdata/compare/import_stmt"+".json")).To(Equal(true))
}

func Test_PythonClassWithFunctionDef(t *testing.T) {
	g := NewGomegaWithT(t)

	app := new(PythonIdentApp)

	file, _ := ioutil.ReadFile("testdata/compare/blog_entity.py")
	codeFile := app.Analysis(string(file), "testdata/compare/blog_entity.py")

	g.Expect(len(codeFile.DataStructures)).To(Equal(1))
	g.Expect(len(codeFile.DataStructures[0].Functions)).To(Equal(4))
}
