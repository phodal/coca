package deps

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure/ast/groovy"
	. "github.com/phodal/coca/languages/groovy"
)

type DepApp struct {
}

func NewDepApp() DepApp {
	return *&DepApp{}
}

func (d *DepApp) BuildImportMap(deps []domain.JClassNode) map[string]domain.JImport {
	var impMap = make(map[string]domain.JImport)
	for _, clz := range deps {
		for _, imp := range clz.Imports {
			impMap[imp.Source] = imp
		}
	}

	return impMap
}

func Analysis(str string) {
	parser := ProcessGroovyString(str)
	context := parser.CompilationUnit()

	listener := groovy.NewGroovyIdentListener()

	antlr.NewParseTreeWalker().Walk(listener, context)
}

func ProcessGroovyString(code string) *GroovyParser {
	is := antlr.NewInputStream(code)
	lexer := NewGroovyLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewGroovyParser(stream)
	return parser
}
