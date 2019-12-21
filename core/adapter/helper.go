package adapter

import (
	"github.com/phodal/coca/core/adapter/identifier"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"encoding/json"
)

func BuildIdentifierMap(identifiers []models.JIdentifier) map[string]models.JIdentifier {
	var identifiersMap = make(map[string]models.JIdentifier)

	for _, ident := range identifiers {
		identifiersMap[ident.Package+"."+ident.ClassName] = ident
	}
	return identifiersMap
}

func LoadIdentify(importPath string) []models.JIdentifier {
	var identifiers []models.JIdentifier

	apiContent := support.ReadCocaFile("identify.json")
	if apiContent == nil {
		identifierApp := new(identifier.JavaIdentifierApp)
		ident := identifierApp.AnalysisPath(importPath)

		identModel, _ := json.MarshalIndent(ident, "", "\t")
		support.WriteToCocaFile("identify.json", string(identModel))

		return *&ident
	}
	_ = json.Unmarshal(apiContent, &identifiers)

	return *&identifiers
}


func BuildDIMap(identifiers []models.JIdentifier, identifierMap map[string]models.JIdentifier) map[string]string {
	var diMap = make(map[string]string)
	for _, clz := range identifiers {
		if len(clz.Annotations) > 0 {
			for _, annotation := range clz.Annotations {
				name := annotation.QualifiedName
				if (name == "Component" || name == "Repository") && len(clz.Implements) > 0 {
					superClz := identifierMap[clz.Implements[0]]
					diMap[superClz.Package + "." + superClz.ClassName] = clz.Package + "." + clz.ClassName
				}
			}
		}
	}

	return diMap
}

