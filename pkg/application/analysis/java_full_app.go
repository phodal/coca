package analysis

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/jdomain"
	"github.com/phodal/coca/pkg/infrastructure/ast"
	"github.com/phodal/coca/pkg/infrastructure/ast/full"
	"path/filepath"
)

type JavaFullApp struct {
}

func NewJavaFullApp() JavaFullApp {
	return JavaFullApp{}
}

func (j *JavaFullApp) AnalysisPath(codeDir string, classes []string, identNodes []jdomain.JIdentifier) []jdomain.JClassNode {
	files := cocafile.GetJavaFiles(codeDir)
	return j.AnalysisFiles(identNodes, files, classes)
}

func (j *JavaFullApp) AnalysisFiles(identNodes []jdomain.JIdentifier, files []string, classes []string) []jdomain.JClassNode {
	var nodeInfos []jdomain.JClassNode

	var identMap = make(map[string]jdomain.JIdentifier)
	for _, ident := range identNodes {
		identMap[ident.GetClassFullName()] = ident
	}

	for _, file := range files {
		displayName := filepath.Base(file)
		fmt.Println("Refactoring parse java call: " + displayName)

		parser := ast.ProcessJavaFile(file)
		context := parser.CompilationUnit()

		listener := full.NewJavaFullListener(identMap, file)
		listener.AppendClasses(classes)

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodes := listener.GetNodeInfo()
		nodeInfos = append(nodeInfos, nodes...)
	}

	return nodeInfos
}
