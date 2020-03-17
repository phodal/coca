package unused

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	base2 "github.com/phodal/coca/pkg/application/refactor/base"
	models2 "github.com/phodal/coca/pkg/application/refactor/base/models"
	"github.com/phodal/coca/pkg/infrastructure/ast/ast_java"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var currentFile string
var configPath string

type RemoveUnusedImportApp struct {
}

func NewRemoveUnusedImportApp(pPath string) *RemoveUnusedImportApp {
	configPath = pPath

	return &RemoveUnusedImportApp{}
}

func (j *RemoveUnusedImportApp) Analysis() []models2.JFullIdentifier {
	files := cocafile.GetJavaFiles(configPath)

	var nodes []models2.JFullIdentifier = nil
	for index := range files {
		file := files[index]

		currentFile, _ = filepath.Abs(file)
		displayName := filepath.Base(file)
		fmt.Println("parse java call: " + displayName)

		parser := ast_java.ProcessJavaFile(file)
		context := parser.CompilationUnit()

		node := models2.NewJFullIdentifier()
		listener := new(base2.JavaRefactorListener)
		listener.InitNode(node)

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodes = append(nodes, listener.GetNodeInfo())
	}

	return nodes
}

func (j *RemoveUnusedImportApp) Refactoring(resultNodes []models2.JFullIdentifier) {
	for _, node := range resultNodes {
		if node.Name != "" {
			errorLines := BuildErrorLines(node)
			removeImportByLines(currentFile, errorLines)
		}
	}
}

func BuildErrorLines(node models2.JFullIdentifier) []int {
	var fields = node.GetFields()
	var imports = node.GetImports()

	var errorLines []int
	for index := range imports {
		imp := imports[index]
		ss := strings.Split(imp.Name, ".")
		lastField := ss[len(ss)-1]

		var isOk = false
		for _, field := range fields {
			if field.Name == lastField || lastField == "*" {
				isOk = true
			}
		}

		if !isOk {
			errorLines = append(errorLines, imp.StartLine)
		}
	}

	return errorLines
}

func removeImportByLines(file string, errorLines []int) {
	removedErrorCount := 1
	for _, line := range errorLines {
		newStart := line - removedErrorCount
		removeLine(file, newStart)
		removedErrorCount++
	}
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
