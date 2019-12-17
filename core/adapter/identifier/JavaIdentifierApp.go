package identifier

import (
	parser2 "coca/core/language/java"
	"coca/core/models"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"
)

var nodeInfos []models.JsonIdentifier = nil

type JavaIdentifierApp struct {
}

func (j *JavaIdentifierApp) AnalysisPath(codeDir string) []models.JsonIdentifier {
	nodeInfos = nil
	files := (*JavaIdentifierApp)(nil).javaFiles(codeDir)
	for index := range files {
		file := files[index]
		node := models.NewJsonIdentifier()

		parser := (*JavaIdentifierApp)(nil).processFile(file)
		context := parser.CompilationUnit()

		clzInfo := models.NewJIdentifier()
		listener := new(JavaIdentifierListener)
		listener.InitNode(clzInfo)

		antlr.NewParseTreeWalker().Walk(listener, context)

		if clzInfo.Name != "" {
			node = &models.JsonIdentifier{clzInfo.Pkg, clzInfo.Name, clzInfo.Type, clzInfo.GetMethods()}
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

func (j *JavaIdentifierApp) processFile(path string) *parser2.JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := parser2.NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := parser2.NewJavaParser(stream)
	return parser
}
