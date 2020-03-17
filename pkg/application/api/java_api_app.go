package api

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	api_domain2 "github.com/phodal/coca/pkg/domain/api_domain"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/ast_java"
	"github.com/phodal/coca/pkg/infrastructure/ast/ast_java/ast_api_java"
	"path/filepath"
)

var allApis []api_domain2.RestAPI

type JavaApiApp struct {
}

func (j *JavaApiApp) AnalysisPath(codeDir string, parsedDeps []core_domain.CodeDataStruct, identifiersMap map[string]core_domain.CodeDataStruct, diMap map[string]string) []api_domain2.RestAPI {
	files := cocafile.GetJavaFiles(codeDir)
	allApis = nil
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("parse java call: " + displayName)

		parser := ast_java.ProcessJavaFile(file)
		context := parser.CompilationUnit()

		listener := ast_api_java.NewJavaAPIListener(identifiersMap, diMap)
		listener.AppendClasses(parsedDeps)

		antlr.NewParseTreeWalker().Walk(listener, context)

		currentRestApis := listener.GetClassApis()
		allApis = append(allApis, currentRestApis...)
	}

	return allApis
}
