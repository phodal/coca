package unused

import (
	base2 "coca/core/domain/refactor/base"
	models2 "coca/core/domain/refactor/base/models"
	utils2 "coca/core/domain/refactor/utils"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var currentFile string
var moveConfig string
var configPath string

type RemoveUnusedImportApp struct {
}

var nodes []models2.JMoveStruct

func NewRemoveUnusedImportApp(config string, pPath string) *RemoveUnusedImportApp {
	moveConfig = config
	configPath = pPath

	nodes = nil
	return &RemoveUnusedImportApp{}
}

func (j *RemoveUnusedImportApp) Analysis() {
	files := utils2.GetJavaFiles(configPath)
	for index := range files {
		file := files[index]

		currentFile, _ = filepath.Abs(file)
		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := utils2.ProcessFile(file)
		context := parser.CompilationUnit()

		node := models2.NewJFullIdentifier()
		listener := new(base2.JavaRefactorListener)
		listener.InitNode(node)

		antlr.NewParseTreeWalker().Walk(listener, context)

		if node.Name != "" {
			handleNode(node)
		}
	}
}

func handleNode(node *models2.JFullIdentifier) {
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

func removeImportByLineNum(imp models2.JImport, line int) {
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
