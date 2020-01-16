package testhelper

import (
	"github.com/phodal/coca/pkg/application/analysis"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"path/filepath"
)

func BuildAnalysisDeps(codePath string) ([]core_domain.CodeDataStruct, map[string]core_domain.JIdentifier, []core_domain.JIdentifier) {
	codePath = filepath.FromSlash(codePath)

	identifierApp := analysis.NewJavaIdentifierApp()
	identifiers := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.NodeName)
	}

	callApp := analysis.NewJavaFullApp()
	callNodes := callApp.AnalysisPath(codePath, classes, identifiers)

	identifiersMap := core_domain.BuildIdentifierMap(identifiers)
	return callNodes, identifiersMap, identifiers
}

