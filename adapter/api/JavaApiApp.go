package api

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"

	. "github.com/phodal/coca/adapter/models"
	. "github.com/phodal/coca/language/java"
)

var nodeInfos []JClassNode

type JavaApiApp struct {

}

func (j *JavaApiApp) AnalysisPath(codeDir string, classes []JClassNode) []JClassNode {
	nodeInfos = nil
	files := (*JavaApiApp)(nil).JavaFiles(codeDir)
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := (*JavaApiApp)(nil).ProcessFile(file)
		context := parser.CompilationUnit()

		listener := NewJavaApiListener()
		listener.appendClasses(classes)

		antlr.NewParseTreeWalker().Walk(listener, context)

		apis := listener.getApis()
		fmt.Println(apis)
		//nodeInfo = listener.getNodeInfo()
		//nodeInfo.Path = file
		//nodeInfos = append(nodeInfos, *nodeInfo)
	}

	return nodeInfos
}

func (j *JavaApiApp) JavaFiles(codeDir string) []string {
	files := make([]string, 0)
	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".java") && !strings.Contains(path, "Test.java") {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func (j *JavaApiApp) ProcessFile(path string) *JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewJavaParser(stream)
	return parser
}
