package moveclass

import (
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	base2 "github.com/phodal/coca/pkg/application/refactor/base"
	models2 "github.com/phodal/coca/pkg/application/refactor/base/models"
	"github.com/phodal/coca/pkg/infrastructure/ast/ast_java"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var currentFile string
var moveConfig string
var configPath string

var nodes []models2.JMoveStruct

type MoveClassApp struct {
}

func NewMoveClassApp(config string, pPath string) *MoveClassApp {
	moveConfig = config
	configPath = pPath

	nodes = nil
	return &MoveClassApp{}
}

func (j *MoveClassApp) Analysis() []models2.JMoveStruct {
	// TODO: 使用 Deps.json 来移动包
	files := cocafile.GetJavaFiles(configPath)
	for index := range files {
		file := files[index]

		currentFile, _ = filepath.Abs(file)

		parser := ast_java.ProcessJavaFile(file)
		context := parser.CompilationUnit()

		node := models2.NewJFullIdentifier()
		listener := new(base2.JavaRefactorListener)
		listener.InitNode(node)

		antlr.NewParseTreeWalker().Walk(listener, context)

		node = listener.GetNodeInfo()
		moveStruct := &models2.JMoveStruct{JFullIdentifier: node, Path: currentFile, Deps: node.GetImports()}
		nodes = append(nodes, *moveStruct)
	}

	return nodes
}

func (j *MoveClassApp) Refactoring() {
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

		originFile  := buildJavaPath(configPath, originImport)
		newFile := buildJavaPath(configPath, newImport)

		// for travis test
		fmt.Println(originFile, newFile)
		copyClass(originFile, newFile)

		updatePackageInfo(originImport, newImport)
		updateImportSide(originImport, newImport)
	}
}

func updatePackageInfo(originImport string, newImport string) {
	var originNode models2.JMoveStruct
	for index := range nodes {
		node := nodes[index]
		if originImport == node.Pkg+"."+node.Name {
			originNode = node
		}
	}

	if originNode.Name == "" {
		return
	}

	path := buildJavaPath(configPath, newImport)
	split := strings.Split(newImport, ".")
	pkg := strings.Join(split[:len(split)-1], ".")
	updateFile(path, originNode.GetPkgInfo().StartLine, "package "+pkg+";")
}

func updateImportSide(originImport string, newImport string) {
	for index := range nodes {
		node := nodes[index]
		for j := range node.Deps {
			dep := node.Deps[j]
			if dep.Name == originImport {
				updateFile(node.Path, dep.StartLine, "import "+newImport+";")
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
			lines[i-1] = newImp
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func copyClass(originFile string, newFile string) {
	_, err := CopyFile(originFile, newFile)
	if err != nil {
		log.Fatalln(err)
	}
}

func buildJavaPath(configPath string, importStr string) string {
	if !strings.HasSuffix(configPath, "/") {
		configPath = configPath + "/"
	}
	path := configPath + strings.ReplaceAll(importStr, ".", "/") + ".java"
	return filepath.FromSlash(path)
}

func CopyFile(src, dst string) (int64, error) {
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
