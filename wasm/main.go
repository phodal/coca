package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/adapter/coca_file"
	"github.com/phodal/coca/core/infrastructure/ast/identifier"
)

func main()  {
	parser := coca_file.ProcessString("package com.phodal.coca;")
	context := parser.CompilationUnit()

	listener := identifier.NewJavaIdentifierListener()

	antlr.NewParseTreeWalker().Walk(listener, context)

	identifiers := listener.GetNodes()
	fmt.Println(identifiers)
}