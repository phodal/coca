package di

import (
	"coca/core/models"
)

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
