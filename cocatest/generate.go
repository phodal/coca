package cocatest

import (
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/core/adapter/cocafile"
	"github.com/phodal/coca/core/context/analysis"
	"github.com/phodal/coca/core/domain"
	"path/filepath"
)

func BuildTestAnalysisResultsByPath(codePath string) (map[string]domain.JIdentifier, []domain.JClassNode) {
	files := cocafile.GetJavaTestFiles(codePath)
	var identifiers []domain.JIdentifier

	identifiers = cmd_util.LoadTestIdentify(files)
	identifiersMap := domain.BuildIdentifierMap(identifiers)

	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	analysisApp := analysis.NewJavaFullApp()
	classNodes := analysisApp.AnalysisFiles(identifiers, files, classes)
	return identifiersMap, classNodes
}

func BuildAnalysisDeps(codePath string) ([]domain.JClassNode, map[string]domain.JIdentifier) {
	codePath = filepath.FromSlash(codePath)

	identifierApp := new(analysis.JavaIdentifierApp)
	identifiers := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := analysis.NewJavaFullApp()
	callNodes := callApp.AnalysisPath(codePath, classes, identifiers)

	identifiersMap := domain.BuildIdentifierMap(identifiers)
	return callNodes, identifiersMap
}

