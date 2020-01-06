package deps

import (
	"fmt"
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

func (d *DepApp) CountDeps(deps []domain.JClassNode) {
	fmt.Println(deps)
}

func Analysis(str string) {
	parser := ProcessGroovyString(str)
	context := parser.CompilationUnit()

	listener := new(groovy.GroovyIdentifierListener)

	antlr.NewParseTreeWalker().Walk(listener, context)
}

func ProcessGroovyString(code string) *GroovyParser {
	is := antlr.NewInputStream(code)
	lexer := NewGroovyLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewGroovyParser(stream)
	return parser
}
