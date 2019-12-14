package api

import (
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"

	. "coca/adapter/models"
	. "coca/language/java"
	. "coca/utils"
)

var parsedDeps []JClassNode
var allApis []RestApi

type JavaApiApp struct {

}

func (j *JavaApiApp) AnalysisPath(codeDir string, depPath string) []RestApi {
	parsedDeps = nil
	file := ReadFile(depPath)
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

		apis := listener.getClassApis()
		allApis = append(allApis, apis...)
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

func (j *JavaApiApp) ProcessFile(path string) *JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewJavaParser(stream)
	return parser
}
