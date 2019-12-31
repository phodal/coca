package cmd_util

import (
	"encoding/json"
	"github.com/phodal/coca/core/context/analysis"
	"github.com/phodal/coca/core/domain"
)

func LoadIdentify(importPath string) []domain.JIdentifier {
	var identifiers []domain.JIdentifier

	apiContent := ReadCocaFile("identify.json")
	if apiContent == nil || string(apiContent) == "null" {
		identifierApp := new(analysis.JavaIdentifierApp)
		ident := identifierApp.AnalysisPath(importPath)

		identModel, _ := json.MarshalIndent(ident, "", "\t")
		WriteToCocaFile("identify.json", string(identModel))

		return *&ident
	}
	_ = json.Unmarshal(apiContent, &identifiers)

	return *&identifiers
}

func LoadTestIdentify(files []string) []domain.JIdentifier {
	var identifiers []domain.JIdentifier

	apiContent := ReadCocaFile("tidentify.json")

	if apiContent == nil || string(apiContent) == "null" {
		identifierApp := analysis.NewJavaIdentifierApp()
		ident := identifierApp.AnalysisFiles(files)

		identModel, _ := json.MarshalIndent(ident, "", "\t")
		WriteToCocaFile("tidentify.json", string(identModel))

		return *&ident
	}
	_ = json.Unmarshal(apiContent, &identifiers)

	return *&identifiers
}

