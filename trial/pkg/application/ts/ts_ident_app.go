package js_ident

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/ts"
	"github.com/phodal/coca/pkg/domain"
	"github.com/phodal/coca/trial/pkg/ast"
)

func processStream(is antlr.CharStream) *parser.TypeScriptParser {
	lexer := parser.NewTypeScriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewTypeScriptParser(stream)
	return parser
}

func ProcessTsString(code string) *parser.TypeScriptParser {
	is := antlr.NewInputStream(code)
	return processStream(is)
}

type TypeScriptApiApp struct {
}

func (j *TypeScriptApiApp) Analysis(code string) domain.JClassNode {
	scriptParser := ProcessTsString(code)
	context := scriptParser.Program()

	listener := ast.NewTypeScriptIdentListener()
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetNodeInfo()
}
