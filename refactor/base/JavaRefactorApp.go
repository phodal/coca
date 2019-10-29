package base

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	. "./models"
	. "../utils"
)

var currentFile string

type JavaRefactorApp struct {
}

func (j *JavaRefactorApp) AnalysisPath(codeDir string) {
	files := GetJavaFiles(codeDir)
	for index := range files {
		file := files[index]

		currentFile, _ = filepath.Abs(file)
		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := ProcessFile(file)
		context := parser.CompilationUnit()

		node := NewJFullIdentifier()
		listener := new(JavaRefactorListener)
		listener.InitNode(node)

		antlr.NewParseTreeWalker().Walk(listener, context)

		if node.Name != "" {
			handleNode()
		}
	}
}

func handleNode() {
	var fields map[string]JField = node.GetFields()
	var imports map[string]JImport = node.GetImports()

	fmt.Println(node.Pkg+"."+node.Name, imports, node.GetMethods(), fields)
	if len(fields) == 0 {
		//removeAllImports(imports)
		return
	}

	//for index := range fields {
		//field := fields[index]
		//errorLine := 0
		//
		//for index := range imports {
		//	imp := imports[index]
		//	ss := strings.Split(imp.Name, ".")
		//	lastField := ss[len(ss)-1]
		//
		//	if (lastField == field.Name) {
		//		continue
		//	} else {
		//
		//	}
		//}
	//}
}

func removeAllImports(imports map[string]JImport) {
	for index := range imports {
		imp := imports[index]
		removeImportByLineNum(imp)
	}
}

func removeImportByLineNum(imp JImport) {
	removeLine(currentFile, imp.StartLine - 1)
}

func removeLine(path string, lineNumber int) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	info, _ := os.Stat(path)
	mode := info.Mode()

	array := strings.Split(string(file), "\n")
	array = append(array[:lineNumber], array[lineNumber+1:]...)
	_ = ioutil.WriteFile(path, []byte(strings.Join(array, "\n")), mode)
}
