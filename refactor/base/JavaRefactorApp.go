package base

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"

	. "../../language/java"
)

type JavaRefactorApp struct {

}

func (j *JavaRefactorApp) AnalysisPath(codeDir string) {
	files := (*JavaRefactorApp)(nil).javaFiles(codeDir)
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := (*JavaRefactorApp)(nil).processFile(file)
		context := parser.CompilationUnit()

		listener := new(JavaRefactorCallListener)

		antlr.NewParseTreeWalker().Walk(listener, context)
	}
}

func (j *JavaRefactorApp) javaFiles(codeDir string) []string {
	files := make([]string, 0)
	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".java") && !strings.Contains(path, "Test.java") {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func (j *JavaRefactorApp) processFile(path string) *JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewJavaParser(stream)
	return parser
}
