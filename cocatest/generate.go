package cocatest

import (
	"github.com/phodal/coca/core/context/analysis"
	"github.com/phodal/coca/core/domain"
	"path/filepath"
)

func BuildAnalysisDeps(codePath string) ([]domain.JClassNode, map[string]domain.JIdentifier, []domain.JIdentifier) {
	codePath = filepath.FromSlash(codePath)

	identifierApp := analysis.NewJavaIdentifierApp()
	identifiers := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := analysis.NewJavaFullApp()
	callNodes := callApp.AnalysisPath(codePath, classes, identifiers)

	identifiersMap := domain.BuildIdentifierMap(identifiers)
	return callNodes, identifiersMap, identifiers
}

