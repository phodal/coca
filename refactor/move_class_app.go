package refactor

import (
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"os"
	"path/filepath"

	. "./base"
	. "./base/models"
	. "./utils"
)

var currentFile string
var moveConfig string
var path string
var nodes []JFullIdentifier

type MoveClassApp struct {
}

func NewMoveClassApp(config string, pPath string) *MoveClassApp {
	moveConfig = config
	path = pPath

	return &MoveClassApp{}
}

func (j *MoveClassApp) Analysis() {
	files := GetJavaFiles(path)
	fmt.Println(path)
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

		nodes = append(nodes, *node)
	}

	parseRename()
}

func parseRename() {
	file, err := os.Open(moveConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		moveClass(scanner.Text())
	}
}

func moveClass(text string) {

}
