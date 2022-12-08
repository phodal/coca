package deps

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/modernizing/coca/cmd/cmd_util"
	"github.com/modernizing/coca/pkg/domain/core_domain"
	"github.com/modernizing/coca/pkg/infrastructure/ast/ast_groovy"
)

func AnalysisGradleFile(path string) []core_domain.CodeDependency {
	bytes := cmd_util.ReadFile(path)
	return AnalysisGradleString(string(bytes))
}

func AnalysisGradleString(str string) []core_domain.CodeDependency {
	parser := ast_groovy.ProcessGroovyString(str)
	context := parser.CompilationUnit()
	listener := ast_groovy.NewGroovyIdentListener()
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetDepsInfo()
}
