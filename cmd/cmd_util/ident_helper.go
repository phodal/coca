package cmd_util

import (
	"encoding/json"
	"github.com/phodal/coca/pkg/application/analysis/javaapp"
	"github.com/phodal/coca/pkg/domain/core_domain"
)

func LoadIdentify(importPath string) []core_domain.CodeDataStruct {
	return readIdentify(importPath, "identify.json", analysisByPath)
}

func LoadTestIdentify(files []string) []core_domain.CodeDataStruct {
	return readIdentify(files, "tidentify.json", analysisByFiles)
}

func readIdentify(importPath interface{}, fileName string, analysisApp func(importPath interface{}) []core_domain.CodeDataStruct) []core_domain.CodeDataStruct {
	var identifiers []core_domain.CodeDataStruct

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

func analysisByPath(importPath interface{}) []core_domain.CodeDataStruct {
	identifierApp := new(javaapp.JavaIdentifierApp)
	ident := identifierApp.AnalysisPath(importPath.(string))
	return ident
}

func analysisByFiles(files interface{}) []core_domain.CodeDataStruct {
	identifierApp := javaapp.NewJavaIdentifierApp()
	ident := identifierApp.AnalysisFiles(files.([]string))
	return ident
}
