package testhelper

import (
	"github.com/phodal/coca/pkg/application/analysis/javaapp"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"path/filepath"
)

func BuildAnalysisDeps(codePath string) ([]core_domain.CodeDataStruct, map[string]core_domain.CodeDataStruct, []core_domain.CodeDataStruct) {
	codePath = filepath.FromSlash(codePath)

	identifierApp := javaapp.NewJavaIdentifierApp()
	identifiers := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.NodeName)
	}

	callApp := javaapp.NewJavaFullApp()
	callNodes := callApp.AnalysisPath(codePath, classes, identifiers)

	identifiersMap := core_domain.BuildIdentifierMap(identifiers)
	return callNodes, identifiersMap, identifiers
}

