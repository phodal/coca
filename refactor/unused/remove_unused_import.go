package unused

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	. "github.com/phodal/coca/refactor/base"
	. "github.com/phodal/coca/refactor/base/models"
	. "github.com/phodal/coca/refactor/utils"
)

var currentFile string
var moveConfig string
var configPath string

type RemoveUnusedImportApp struct {
}

var nodes []JMoveStruct

func NewRemoveUnusedImportApp(config string, pPath string) *RemoveUnusedImportApp {
	moveConfig = config
	configPath = pPath

	nodes = nil
	return &RemoveUnusedImportApp{}
}

func (j *RemoveUnusedImportApp) Analysis() {
	files := GetJavaFiles(configPath)
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
			handleNode(node)
		}
	}
}

func handleNode(node *JFullIdentifier) {
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

	removeImportByLines(currentFile, errorLines)
}

func removeImportByLines(file string, errorLines []int) {
	removedErrorCount := 1
	for _, line := range errorLines {
		newStart := line - removedErrorCount
		removeLine(file, newStart)
		removedErrorCount++
	}
}

func removeImportByLineNum(imp JImport, line int) {
	removeLine(currentFile, line)
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
