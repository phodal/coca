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
		context := processFile(file)

		v := &BaseJavaParserVisitor{}
		visit := v.Visit(context)

		fmt.Println(visit)
	}
}

func javaFiles(codeDir string) []string {
	files := make([]string, 0)
	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".java") && !strings.Contains(path, "Test.java") {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func processFile(path string) ICompilationUnitContext {
	is, _ := antlr.NewFileStream(path)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewJavaParser(stream)
	parser.BuildParseTrees = true
	return parser.CompilationUnit()
}
