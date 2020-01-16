package testhelper

import (
	"github.com/phodal/coca/pkg/application/analysis"
	"github.com/phodal/coca/pkg/domain/jdomain"
	"path/filepath"
)

func BuildAnalysisDeps(codePath string) ([]jdomain.JClassNode, map[string]jdomain.JIdentifier, []jdomain.JIdentifier) {
	codePath = filepath.FromSlash(codePath)

	identifierApp := analysis.NewJavaIdentifierApp()
	identifiers := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := analysis.NewJavaFullApp()
	callNodes := callApp.AnalysisPath(codePath, classes, identifiers)

	identifiersMap := jdomain.BuildIdentifierMap(identifiers)
	return callNodes, identifiersMap, identifiers
}

