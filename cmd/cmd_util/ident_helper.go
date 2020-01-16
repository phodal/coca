package cmd_util

import (
	"encoding/json"
	"github.com/phodal/coca/pkg/application/analysis"
	"github.com/phodal/coca/pkg/domain/jdomain"
)

func LoadIdentify(importPath string) []jdomain.JIdentifier {
	return readIdentify(importPath, "identify.json", analysisByPath)
}

func LoadTestIdentify(files []string) []jdomain.JIdentifier {
	return readIdentify(files, "tidentify.json", analysisByFiles)
}

func readIdentify(importPath interface{}, fileName string, analysisApp func(importPath interface{}) []jdomain.JIdentifier) []jdomain.JIdentifier {
	var identifiers []jdomain.JIdentifier

	apiContent := ReadCocaFile(fileName)
	if apiContent == nil || string(apiContent) == "null" {
		ident := analysisApp(importPath)

		identModel, _ := json.MarshalIndent(ident, "", "\t")
		WriteToCocaFile(fileName, string(identModel))

		return ident
	}
	_ = json.Unmarshal(apiContent, &identifiers)
	return identifiers
}

func analysisByPath(importPath interface{}) []jdomain.JIdentifier {
	identifierApp := new(analysis.JavaIdentifierApp)
	ident := identifierApp.AnalysisPath(importPath.(string))
	return ident
}

func analysisByFiles(files interface{}) []jdomain.JIdentifier {
	identifierApp := analysis.NewJavaIdentifierApp()
	ident := identifierApp.AnalysisFiles(files.([]string))
	return ident
}
