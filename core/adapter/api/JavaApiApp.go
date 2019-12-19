package api

import (
	"coca/core/models"
	"coca/core/support"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
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

	files := support.GetJavaFiles(codeDir)
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := support.ProcessFile(file)
		context := parser.CompilationUnit()

		listener := NewJavaApiListener()
		listener.appendClasses(parsedDeps)

		antlr.NewParseTreeWalker().Walk(listener, context)

		allApis = listener.getClassApis()
	}

	return *&allApis
}

