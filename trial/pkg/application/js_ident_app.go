package application

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/js"
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

func (j *JavaScriptApiApp) Analysis(code string) {
	jsParser := ProcessJsString(code)
	program := jsParser.Program()

	fmt.Println(program)
}
