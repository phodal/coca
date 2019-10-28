package identifier

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"

	. "../../language/java"
	. "./models"
)

type JavaIdentifierApp struct {

}

func (j *JavaIdentifierApp) AnalysisPath(codeDir string) {
	files := (*JavaIdentifierApp)(nil).javaFiles(codeDir)
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := (*JavaIdentifierApp)(nil).processFile(file)
		context := parser.CompilationUnit()

		interfaceIdent := &JIdentifier{"", "", ""}
		listener := new(JavaIdentifierListener)
		listener.InitNode(interfaceIdent)

		antlr.NewParseTreeWalker().Walk(listener, context)

		if interfaceIdent.Name != "" {
			fmt.Println(interfaceIdent.Type, interfaceIdent.Pkg, interfaceIdent.Name)
		}
	}
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
