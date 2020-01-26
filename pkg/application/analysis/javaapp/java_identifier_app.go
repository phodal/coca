package javaapp

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/ast_java"
	"github.com/phodal/coca/pkg/infrastructure/ast/ast_java/java_identify"
)


type JavaIdentifierApp struct {
}

func NewJavaIdentifierApp() JavaIdentifierApp {
	return JavaIdentifierApp{}
}

func (j *JavaIdentifierApp) AnalysisPath(codeDir string) []core_domain.CodeDataStruct {
	files := cocafile.GetJavaFiles(codeDir)
	return j.AnalysisFiles(files)
}

func (j *JavaIdentifierApp) AnalysisFiles(files []string) []core_domain.CodeDataStruct {
	var nodeInfos []core_domain.CodeDataStruct = nil

	for _, file := range files {
		parser := ast_java.ProcessJavaFile(file)
		context := parser.CompilationUnit()

		listener := java_identify.NewJavaIdentifierListener()

		antlr.NewParseTreeWalker().Walk(listener, context)

		identifiers := listener.GetNodes()
		nodeInfos = append(nodeInfos, identifiers...)
	}

	return nodeInfos
}
