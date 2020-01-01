package main

import (
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/adapter/coca_file"
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure/ast/full"
	"github.com/phodal/coca/core/infrastructure/ast/identifier"
	"syscall/js"
)

func registerCallbacks() {
	js.Global().Set("compileCode", js.FuncOf(CompileCodeCallback))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}

func CompileCodeCallback(value js.Value, args []js.Value) interface{} {
	callback := args[len(args)-1:][0]
	message := args[0].String()

	results := CompileCode(message)

	identModel, _ := json.Marshal(results)
	callback.Invoke(js.Null(), string(identModel))
	return nil
}

func CompileCode(code string) []domain.JClassNode {
	classes, identMap := prepareForAnalysis(code)

	parser := coca_file.ProcessString(code)
	context := parser.CompilationUnit()

	listener := full.NewJavaFullListener(identMap, "hello")
	listener.AppendClasses(classes)

	antlr.NewParseTreeWalker().Walk(listener, context)

	nodes := listener.GetNodeInfo()
	fmt.Println(nodes)

	return nodes

}

func prepareForAnalysis(code string) ([]string, map[string]domain.JIdentifier) {
	parser := coca_file.ProcessString(code)
	context := parser.CompilationUnit()

	listener := identifier.NewJavaIdentifierListener()
	antlr.NewParseTreeWalker().Walk(listener, context)

	identifiers := listener.GetNodes()

	var classes []string = nil

	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	var identMap = make(map[string]domain.JIdentifier)
	for _, ident := range identifiers {
		identMap[ident.GetClassFullName()] = ident
	}
	return classes, identMap
}
