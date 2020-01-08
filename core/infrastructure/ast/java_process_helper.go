package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/languages/java"
)

func ProcessJavaFile(path string) *parser.JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := parser.NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewJavaParser(stream)
	return parser
}

func ProcessJavaString(code string) *parser.JavaParser {
	is := antlr.NewInputStream(code)
	lexer := parser.NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewJavaParser(stream)
	return parser
}
