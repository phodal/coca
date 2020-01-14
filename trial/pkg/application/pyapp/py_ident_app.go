package pyapp

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/python"
	"github.com/phodal/coca/trial/pkg/ast/pyast"
)

func processStream(is antlr.CharStream) *parser.PythonParser {
	lexer := parser.NewPythonLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	return parser.NewPythonParser(stream)
}

func ProcessTsString(code string) *parser.PythonParser {
	is := antlr.NewInputStream(code)
	return processStream(is)
}

type PythonApiApp struct {
}

func (j *PythonApiApp) Analysis(code string, fileName string) {
	scriptParser := ProcessTsString(code)
	context := scriptParser.Root()

	listener := pyast.NewPythonIdentListener(fileName)
	antlr.NewParseTreeWalker().Walk(listener, context)
}
