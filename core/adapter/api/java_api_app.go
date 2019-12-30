package api

import (
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/infrastructure"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"path/filepath"
)

var allApis []RestApi

type JavaApiApp struct {
}

func (j *JavaApiApp) AnalysisPath(codeDir string, parsedDeps []models.JClassNode, identifiersMap map[string]models.JIdentifier, diMap map[string]string) []RestApi {
	files := infrastructure.GetJavaFiles(codeDir)
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := infrastructure.ProcessFile(file)
		context := parser.CompilationUnit()

		listener := NewJavaApiListener(identifiersMap, diMap)
		listener.appendClasses(parsedDeps)

		antlr.NewParseTreeWalker().Walk(listener, context)

		allApis = listener.getClassApis()
	}

	return *&allApis
}
