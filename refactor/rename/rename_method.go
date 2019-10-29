package unused

import (
	. "../base"
	. "../base/models"
	. "../utils"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
)

var currentFile string
var moveConfig string
var configPath string

type RemoveMethodApp struct {
}

var nodes []JMoveStruct

func NewRemoveMethodApp(config string, pPath string) *RemoveMethodApp {
	moveConfig = config
	configPath = pPath

	nodes = nil
	return &RemoveMethodApp{}
}

func (j *RemoveMethodApp) Analysis() {
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

}
