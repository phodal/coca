package application

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/js"
	"github.com/phodal/coca/pkg/domain"
	"github.com/phodal/coca/trial/pkg/ast"
)

func processStream(is antlr.CharStream) *parser.JavaScriptParser {
	lexer := parser.NewJavaScriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewJavaScriptParser(stream)
	return parser
}

func ProcessJsString(code string) *parser.JavaScriptParser {
	is := antlr.NewInputStream(code)
	return processStream(is)
}

type JavaScriptApiApp struct {
}

func (j *JavaScriptApiApp) Analysis(code string) domain.JClassNode {
	jsParser := ProcessJsString(code)
	context := jsParser.Program()

	listener := ast.NewJavaScriptIdentListener()
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetNodeInfo()
}
