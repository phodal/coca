package call

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	. "github.com/phodal/coca/adapter/models"
	. "github.com/phodal/coca/language/java"
)

var nodeInfos []JClassNode
type BadSmellModel struct {
	File string
	Line string
	Bs   string
}

type BadSmellApp struct {

}

func (j *BadSmellApp) AnalysisPath(codeDir string) []BadSmellModel {
	nodeInfos = nil
	files := (*BadSmellApp)(nil).javaFiles(codeDir)
	for index := range files {
		nodeInfo := NewClassNode()
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := (*BadSmellApp)(nil).processFile(file)
		context := parser.CompilationUnit()

		listener := NewBadSmellListener()

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodeInfo = listener.getNodeInfo()
		nodeInfo.Path = file
		nodeInfos = append(nodeInfos, *nodeInfo)
	}

	bsList := analysisBadSmell(nodeInfos)

	return bsList
}

func analysisBadSmell(nodes []JClassNode) []BadSmellModel {
	var badSmellList []BadSmellModel
	for _, node := range nodes {
		for _, method := range node.Methods {
			if method.StopLine - method.StartLine > 50 {
				longMethod := &BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "longMethod"}
				badSmellList = append(badSmellList, *longMethod)
			}
		}
	}

	return badSmellList
}

func (j *BadSmellApp) javaFiles(codeDir string) []string {
	files := make([]string, 0)
	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".java") && !strings.Contains(path, "Test.java") {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func (j *BadSmellApp) processFile(path string) *JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewJavaParser(stream)
	return parser
}
