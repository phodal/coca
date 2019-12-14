package identifier

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"

	. "coca/adapter/models"
	. "coca/language/java"
)

var nodeInfos []JsonIdentifier = nil

type JavaIdentifierApp struct {
}

func (j *JavaIdentifierApp) AnalysisPath(codeDir string) []JsonIdentifier {
	nodeInfos = nil
	files := (*JavaIdentifierApp)(nil).javaFiles(codeDir)
	for index := range files {
		file := files[index]
		node := NewJsonIdentifier()

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := (*JavaIdentifierApp)(nil).processFile(file)
		context := parser.CompilationUnit()

		clzInfo := NewJIdentifier()
		listener := new(JavaIdentifierListener)
		listener.InitNode(clzInfo)

		antlr.NewParseTreeWalker().Walk(listener, context)

		if clzInfo.Name != "" {
			node = &JsonIdentifier{clzInfo.Pkg, clzInfo.Name, clzInfo.Type, clzInfo.GetMethods()}
			nodeInfos = append(nodeInfos, *node)

		}
	}

	return nodeInfos
}

func (j *JavaIdentifierApp) javaFiles(codeDir string) []string {
	files := make([]string, 0)
	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".java") && !strings.Contains(path, "Test.java") {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func (j *JavaIdentifierApp) processFile(path string) *JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewJavaParser(stream)
	return parser
}
