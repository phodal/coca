package main

import (
	//"./cmd"

	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"

	. "./language/java"
)

func main() {
	//cmd.Execute()

	analysisPath("examples/step2-Java")
}

func analysisPath(codeDir string) {
	files := javaFiles(codeDir)
	for index := range files {
		file := files[index]
		Parser(file)
	}
}

func javaFiles(codeDir string) []string {
	files := make([]string, 0)
	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if (strings.HasSuffix(path, ".java") && !strings.Contains(path, "Test.java")) {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func Parser(path string) ICompilationUnitContext {
	is, _ := antlr.NewFileStream(path)
	fmt.Println(is)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewJavaParser(stream)
	return parser.CompilationUnit()
}
