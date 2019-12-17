package api

import (
	parser2 "coca/core/languages/java"
	"coca/core/models"
	"coca/core/support"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"
)

var parsedDeps []models.JClassNode
var allApis []RestApi

type JavaApiApp struct {

}

func (j *JavaApiApp) AnalysisPath(codeDir string, depPath string) []RestApi {
	parsedDeps = nil
	file := support.ReadFile(depPath)
	if file == nil {
		return nil
	}

	_ = json.Unmarshal(file, &parsedDeps)

	files := (*JavaApiApp)(nil).JavaFiles(codeDir)
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := (*JavaApiApp)(nil).ProcessFile(file)
		context := parser.CompilationUnit()

		listener := NewJavaApiListener()
		listener.appendClasses(parsedDeps)

		antlr.NewParseTreeWalker().Walk(listener, context)

		allApis = listener.getClassApis()
	}

	return allApis
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

func (j *JavaApiApp) ProcessFile(path string) *parser2.JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := parser2.NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := parser2.NewJavaParser(stream)
	return parser
}
