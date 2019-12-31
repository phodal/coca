package analysis

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/adapter/coca_file"
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure/ast/identifier"
)


type JavaIdentifierApp struct {
}

func NewJavaIdentifierApp() JavaIdentifierApp {
	return *&JavaIdentifierApp{}
}

func (j *JavaIdentifierApp) AnalysisPath(codeDir string) []domain.JIdentifier {
	files := coca_file.GetJavaFiles(codeDir)
	return j.AnalysisFiles(files)
}

func (j *JavaIdentifierApp) AnalysisFiles(files []string) []domain.JIdentifier {
	var nodeInfos []domain.JIdentifier = nil

	for _, file := range files {
		parser := coca_file.ProcessFile(file)
		context := parser.CompilationUnit()

		listener := identifier.NewJavaIdentifierListener()

		antlr.NewParseTreeWalker().Walk(listener, context)

		identifiers := listener.GetNodes()
		nodeInfos = append(nodeInfos, identifiers...)
	}

	return nodeInfos
}
