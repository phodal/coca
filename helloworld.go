package main

import (
	//"./cmd"

	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"

	. "./language/java"
)

func main() {
	//cmd.Execute()

	// Setup the input
	//is := antlr.NewInputStream("1 + 2 * 3")

	is, _ := antlr.NewFileStream("examples/step2-Java/domain/AggregateRoot.java")
	fmt.Println(is);

	lexer := NewJavaLexer(is);
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewJavaParser(stream);
	localctx := parser.CompilationUnit();

	fmt.Println(localctx.GetText())
}
