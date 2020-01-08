package deps

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure/ast/groovy"
	"github.com/phodal/coca/languages/groovy"
)

func AnalysisGradle(str string) []domain.JDependency {
	parser := ProcessGroovyString(str)
	context := parser.CompilationUnit()

	listener := groovy.NewGroovyIdentListener()

	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetDepsInfo()
}

func ProcessGroovyString(code string) *parser.GroovyParser {
	is := antlr.NewInputStream(code)
	lexer := parser.NewGroovyLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewGroovyParser(stream)
	return parser
}
