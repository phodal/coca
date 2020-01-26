package ast_groovy

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/languages/groovy"
)

func ProcessGroovyString(code string) *parser.GroovyParser {
	is := antlr.NewInputStream(code)
	lexer := parser.NewGroovyLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewGroovyParser(stream)
	return parser
}
