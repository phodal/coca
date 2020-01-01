package wadapter

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/adapter/coca_file"
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure/ast/full"
	"github.com/phodal/coca/core/infrastructure/ast/identifier"
)

func CompileCode(code string) []domain.JClassNode {
	classes, identMap := prepareForAnalysis(code)

	parser := coca_file.ProcessString(code)
	context := parser.CompilationUnit()

	listener := full.NewJavaFullListener(identMap, "hello")
	listener.AppendClasses(classes)

	antlr.NewParseTreeWalker().Walk(listener, context)

	nodes := listener.GetNodeInfo()
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
