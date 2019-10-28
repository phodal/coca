package main

import (
	//"./cmd"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"

	. "./language/java"
	. "./visitor"
)

func main() {
	//cmd.Execute()
	path := "."

	if len(os.Args) > 1 {
		path = os.Args[1:][0]
	}
	analysisPath(path)
}

func analysisPath(codeDir string) {
	files := javaFiles(codeDir)
	for index := range files {
		file := files[index]
		parser := processFile(file)
		context := parser.CompilationUnit()

		context.GetStart()

		v := NewJavaCallVisitor()
		v.Visit(context)

		context.Accept(v);
		v.BaseParseTreeVisitor.Visit(context)
		//fmt.Println(context.GetText())
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

func processFile(path string) *JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewJavaParser(stream)
	return parser
}
