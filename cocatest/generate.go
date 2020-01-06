package cocatest

import (
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/core/adapter/coca_file"
	"github.com/phodal/coca/core/context/analysis"
	"github.com/phodal/coca/core/domain"
)

func BuildAnalysisResultsByPath(codePath string) (map[string]domain.JIdentifier, []domain.JClassNode) {
	files := coca_file.GetJavaTestFiles(codePath)
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
