package arch

import (
	"github.com/phodal/coca/core/domain/arch/tequila"
	"github.com/phodal/coca/core/models"
)

type ArchApp struct {
}

func NewArchApp() ArchApp {
	return *&ArchApp{}
}

func (a ArchApp) Analysis(deps []models.JClassNode, identifiersMap map[string]models.JIdentifier) tequila.FullGraph {
	fullGraph := tequila.FullGraph{
		NodeList:     make(map[string]string),
		RelationList: make(map[string]*tequila.Relation),
	}

	for _, clz := range deps {
		src := clz.Package + "." + clz.Class
		fullGraph.NodeList[src] = src

		for _, method := range clz.Methods {
			if method.Name == "Main" {
				continue
			}

			// TODO: add implements, extends support
			for _, call := range method.MethodCalls {
				dst := call.Package + "." + call.Class
				if _, ok := identifiersMap[dst]; ok {
					relation := &tequila.Relation{
						From:  src,
						To:    dst,
						Style: "\"solid\"",
					}

					fullGraph.RelationList[relation.From+"->"+relation.To] = relation
				}
			}
		}
	}

	return *&fullGraph
}
