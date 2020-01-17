package pyapp

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/python"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/pyast"
)

func streamToParser(is antlr.CharStream) *parser.PythonParser {
	lexer := parser.NewPythonLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	return parser.NewPythonParser(tokens)
}

func ProcessTsString(code string) *parser.PythonParser {
	is := antlr.NewInputStream(code)
	return streamToParser(is)
}

type PythonIdentApp struct {

}

func (j *PythonIdentApp) Analysis(code string, fileName string) core_domain.CodeFile {
	scriptParser := ProcessTsString(code)
	context := scriptParser.Root()

	listener := pyast.NewPythonIdentListener(fileName)
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetCodeFileInfo()
}
