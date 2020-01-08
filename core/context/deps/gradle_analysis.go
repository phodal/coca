package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure/ast"
	"github.com/phodal/coca/core/infrastructure/ast/groovy"
)

func AnalysisGradleFile(path string) []domain.JDependency {
	bytes := cmd_util.ReadFile(path)
	return AnalysisGradleString(string(bytes))
}

func AnalysisGradleString(str string) []domain.JDependency {
	parser := ast.ProcessGroovyString(str)
	context := parser.CompilationUnit()
	listener := groovy.NewGroovyIdentListener()
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetDepsInfo()
}
