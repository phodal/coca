package refactor

import (
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	. "./base"
	. "./base/models"
	. "./utils"
)

var currentFile string
var moveConfig string
var path string
var nodes map[string]JFullIdentifier

type MoveClassApp struct {
}

func NewMoveClassApp(config string, pPath string) *MoveClassApp {
	moveConfig = config
	path = pPath

	nodes = make(map[string]JFullIdentifier)
	return &MoveClassApp{}
}

func (j *MoveClassApp) Analysis() {
	files := GetJavaFiles(path)
	for index := range files {
		file := files[index]

		currentFile, _ = filepath.Abs(file)
		//displayName := filepath.Base(file)

		parser := ProcessFile(file)
		context := parser.CompilationUnit()

		node := NewJFullIdentifier()
		listener := new(JavaRefactorListener)
		listener.InitNode(node)

		antlr.NewParseTreeWalker().Walk(listener, context)

		pkgPrefix := node.Pkg + "." + node.Name
		if node.Pkg == "" {
			pkgPrefix = node.Name
		}

		nodes[pkgPrefix] = *node
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
		moveClass(scanner.Text())
	}
}

func moveClass(text string) {
	splitStr := strings.Split(text, " -> ")
	if len(splitStr) < 2 {
		return
	}

	originFile, _ := filepath.Abs(path + splitStr[0])
	newFile, _ := filepath.Abs(path + splitStr[1])

	originFile = strings.ReplaceAll(originFile, ".", "/") + ".java"
	newFile = strings.ReplaceAll(newFile, ".", "/") + ".java"

	fmt.Println(originFile, newFile)
	_, err := copy(originFile, newFile)
	if err != nil {
		panic(err)
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}
