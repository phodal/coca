package identifier

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
)

var nodeInfos []models.JIdentifier = nil

type JavaIdentifierApp struct {
}

func NewJavaIdentifierApp() JavaIdentifierApp {
	return *&JavaIdentifierApp{}
}

func (j *JavaIdentifierApp) AnalysisPath(codeDir string) []models.JIdentifier {
	nodeInfos = nil
	files := support.GetJavaFiles(codeDir)

	for _, file := range files {
		parser := support.ProcessFile(file)
		context := parser.CompilationUnit()

		listener := NewJavaIdentifierListener()

		antlr.NewParseTreeWalker().Walk(listener, context)

		identifiers := listener.getNodes()
		nodeInfos = append(nodeInfos, identifiers...)
	}

	return nodeInfos
}
