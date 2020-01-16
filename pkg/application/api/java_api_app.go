package api

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/domain/support_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast"
	"github.com/phodal/coca/pkg/infrastructure/ast/api"
	"path/filepath"
)

var allApis []api_domain.RestAPI

type JavaApiApp struct {
}

func (j *JavaApiApp) AnalysisPath(codeDir string, parsedDeps []core_domain.CodeDataStruct, identifiersMap map[string]core_domain.JIdentifier, diMap map[string]string) []api_domain.RestAPI {
	files := cocafile.GetJavaFiles(codeDir)
	allApis = nil
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Refactoring parse java call: " + displayName)

		parser := ast.ProcessJavaFile(file)
		context := parser.CompilationUnit()

		listener := api.NewJavaAPIListener(identifiersMap, diMap)
		listener.AppendClasses(parsedDeps)

		antlr.NewParseTreeWalker().Walk(listener, context)

		currentRestApis := listener.GetClassApis()
		allApis = append(allApis, currentRestApis...)
	}

	return allApis
}
