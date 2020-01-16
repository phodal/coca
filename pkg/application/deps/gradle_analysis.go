package deps

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/domain/support_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast"
	"github.com/phodal/coca/pkg/infrastructure/ast/groovy"
)

func AnalysisGradleFile(path string) []api_domain.JDependency {
	bytes := cmd_util.ReadFile(path)
	return AnalysisGradleString(string(bytes))
}

func AnalysisGradleString(str string) []api_domain.JDependency {
	parser := ast.ProcessGroovyString(str)
	context := parser.CompilationUnit()
	listener := groovy.NewGroovyIdentListener()
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetDepsInfo()
}
