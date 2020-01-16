package ts

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/ts"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/trial/pkg/ast/ts"
)

func processStream(is antlr.CharStream) *parser.TypeScriptParser {
	lexer := parser.NewTypeScriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	return parser.NewTypeScriptParser(stream)
}

func ProcessTsString(code string) *parser.TypeScriptParser {
	is := antlr.NewInputStream(code)
	return processStream(is)
}

type TypeScriptApiApp struct {
}

func (j *TypeScriptApiApp) Analysis(code string, fileName string) core_domain.CodeFile {
	scriptParser := ProcessTsString(code)
	context := scriptParser.Program()

	listener := ts.NewTypeScriptIdentListener(fileName)
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetNodeInfo()
}
