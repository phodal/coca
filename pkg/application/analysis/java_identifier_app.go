package analysis

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/jdomain"
	"github.com/phodal/coca/pkg/infrastructure/ast"
	"github.com/phodal/coca/pkg/infrastructure/ast/identifier"
)


type JavaIdentifierApp struct {
}

func NewJavaIdentifierApp() JavaIdentifierApp {
	return JavaIdentifierApp{}
}

func (j *JavaIdentifierApp) AnalysisPath(codeDir string) []jdomain.JIdentifier {
	files := cocafile.GetJavaFiles(codeDir)
	return j.AnalysisFiles(files)
}

func (j *JavaIdentifierApp) AnalysisFiles(files []string) []jdomain.JIdentifier {
	var nodeInfos []jdomain.JIdentifier = nil

	for _, file := range files {
		parser := ast.ProcessJavaFile(file)
		context := parser.CompilationUnit()

		listener := identifier.NewJavaIdentifierListener()

		antlr.NewParseTreeWalker().Walk(listener, context)

		identifiers := listener.GetNodes()
		nodeInfos = append(nodeInfos, identifiers...)
	}

	return nodeInfos
}
