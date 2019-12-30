package identifier

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure"
)


type JavaIdentifierApp struct {
}

func NewJavaIdentifierApp() JavaIdentifierApp {
	return *&JavaIdentifierApp{}
}

func (j *JavaIdentifierApp) AnalysisPath(codeDir string) []domain.JIdentifier {
	files := infrastructure.GetJavaFiles(codeDir)
	return j.AnalysisFiles(files)
}

func (j *JavaIdentifierApp) AnalysisFiles(files []string) []domain.JIdentifier {
	var nodeInfos []domain.JIdentifier = nil

	for _, file := range files {
		parser := infrastructure.ProcessFile(file)
		context := parser.CompilationUnit()

		listener := NewJavaIdentifierListener()

		antlr.NewParseTreeWalker().Walk(listener, context)

		identifiers := listener.getNodes()
		nodeInfos = append(nodeInfos, identifiers...)
	}

	return nodeInfos
}
