package call

import (
	parser2 "coca/core/languages/java"
	"coca/core/models"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"
)

var nodeInfos []models.JClassNode

type JavaCallApp struct {
}

func (j *JavaCallApp) AnalysisPath(codeDir string, classes []string, identNodes []models.JsonIdentifier) []models.JClassNode {
	nodeInfos = nil
	files := (*JavaCallApp)(nil).javaFiles(codeDir)
	for index := range files {
		nodeInfo := models.NewClassNode()
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := (*JavaCallApp)(nil).processFile(file)
		context := parser.CompilationUnit()

		listener := NewJavaCallListener(identNodes)
		listener.appendClasses(classes)

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodeInfo = listener.getNodeInfo()
		nodeInfo.Path = file
		nodeInfos = append(nodeInfos, *nodeInfo)
	}

	return nodeInfos
}

func (j *JavaCallApp) javaFiles(codeDir string) []string {
	files := make([]string, 0)
	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".java") && !strings.Contains(path, "Test.java") && !strings.Contains(path, "Tests.java") {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func (j *JavaCallApp) processFile(path string) *parser2.JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := parser2.NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser2.NewJavaParser(stream)
	return parser
}
