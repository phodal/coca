package refactor

import (
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io"
	"io/ioutil"
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

type JMoveStruct struct {
	*JFullIdentifier

	path string
	deps []JImport
}

var nodes []JMoveStruct

type MoveClassApp struct {
}

func NewMoveClassApp(config string, pPath string) *MoveClassApp {
	moveConfig = config
	path = pPath

	nodes = nil
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

		moveStruct := &JMoveStruct{node, currentFile, node.GetImports()}
		nodes = append(nodes, *moveStruct)
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
		splitStr := strings.Split(scanner.Text(), " -> ")
		if len(splitStr) < 2 {
			return
		}

		originImport := splitStr[0]
		newImport := splitStr[1]

		originFile, _ := filepath.Abs(path + originImport)
		newFile, _ := filepath.Abs(path + newImport)

		moveClass(originFile, newFile)

		for index := range nodes {
			node := nodes[index]
			for j := range node.deps {
				dep := node.deps[j]
				if dep.Name == originImport {
					updateFile(node.path, dep.StartLine, "import " + newImport + ";")
				}
			}
		}
	}
}

func updateFile(path string, lineNum int, newImp string) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i := range lines {
		if i == lineNum {
			lines[i - 1] = newImp
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func moveClass(originFile string, newFile string) {
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
