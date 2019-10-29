package base

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"

	. "../../language/java"
	. "./models"
)

var currentFile string

type JavaRefactorApp struct {
}

func (j *JavaRefactorApp) AnalysisPath(codeDir string) {
	files := (*JavaRefactorApp)(nil).javaFiles(codeDir)
	for index := range files {
		file := files[index]

		currentFile, _ = filepath.Abs(file)
		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := (*JavaRefactorApp)(nil).processFile(file)
		context := parser.CompilationUnit()

		node := NewJFullIdentifier()
		listener := new(JavaRefactorListener)
		listener.InitNode(node)

		antlr.NewParseTreeWalker().Walk(listener, context)

		if node.Name != "" {
			handleNode()
		}
	}
}

func handleNode() {
	var fields map[string]JField = node.GetFields()
	var imports map[string]JImport = node.GetImports()

	fmt.Println(node.Pkg+"."+node.Name, imports, node.GetMethods(), fields)

	if len(fields) == 0 {
		removeAllImports(imports)
	}

	for index := range fields {
		field := fields[index]
		fmt.Println(field)
	}
}

func removeAllImports(imports map[string]JImport) {
	for index := range imports {
		imp := imports[index]
		removeImportByLineNum(imp)
	}
}

func removeImportByLineNum(imp JImport) {
	fmt.Println(currentFile, imp)
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
