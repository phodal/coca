package move_class

import (
	"bufio"
	"coca/src/refactor/base"
	"coca/src/refactor/base/models"
	"coca/src/refactor/utils"
	utils2 "coca/src/support"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var currentFile string
var moveConfig string
var configPath string

var nodes []models.JMoveStruct

type MoveClassApp struct {
}

func NewMoveClassApp(config string, pPath string) *MoveClassApp {
	moveConfig = config
	configPath = pPath

	nodes = nil
	return &MoveClassApp{}
}

func (j *MoveClassApp) Analysis() {
	// TODO: 使用 Deps.json 来移动包
	files := utils.GetJavaFiles(configPath)
	fmt.Println(files)
	for index := range files {
		file := files[index]

		currentFile, _ = filepath.Abs(file)
		//displayName := filepath.Base(file)

		parser := utils.ProcessFile(file)
		context := parser.CompilationUnit()

		node := models.NewJFullIdentifier()
		listener := new(base.JavaRefactorListener)
		listener.InitNode(node)

		antlr.NewParseTreeWalker().Walk(listener, context)

		moveStruct := &models.JMoveStruct{node, currentFile, node.GetImports()}
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

		originFile, _ := filepath.Abs(configPath + originImport)
		newFile, _ := filepath.Abs(configPath + newImport)

		copyClass(originFile, newFile)

		updatePackageInfo(nodes, originImport, newImport)

		updateImportSide(originImport, newImport)
	}
}

func updatePackageInfo(structs []models.JMoveStruct, originImport string, newImport string)  {
	var originNode models.JMoveStruct
	for index := range nodes {
		node := nodes[index]
		if originImport == node.Pkg + "." + node.Name {
			originNode = node
		}
	}

	if originNode.Name == "" {
		return
	}
	path := buildJavaPath(configPath + newImport)
	split := strings.Split(newImport, ".")
	pkg := strings.Join(split[:len(split) - 1], ".")
	fmt.Println(pkg)
	updateFile(path, originNode.GetPkgInfo().StartLine, "package " + pkg + ";")
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
			lines[i - 1] = newImp
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func copyClass(originFile string, newFile string) {
	originFile = buildJavaPath(originFile)
	// TODO: 适配 Windows
	if runtime.GOOS == "windows" {
		newFile = strings.ReplaceAll(newFile, ".", "\\") + ".java"
	} else {
		newFile = strings.ReplaceAll(newFile, ".", "/") + ".java"
	}

	fmt.Println(originFile, newFile)
	_, err := utils2.CopyFile(originFile, newFile)
	if err != nil {
		panic(err)
	}
}

func buildJavaPath(originFile string) string {
	// TODO: 适配 Windows

	str := ""
	if runtime.GOOS == "windows" {
		str = strings.ReplaceAll(originFile, ".", "\\") + ".java"
	} else {
		str = strings.ReplaceAll(originFile, ".", "/") + ".java"
	}
	return str
}
