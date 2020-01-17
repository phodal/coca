package arch

import (
	"github.com/phodal/coca/pkg/application/arch/tequila"
	"github.com/phodal/coca/pkg/domain/core_domain"
)

type ArchApp struct {
}

func NewArchApp() ArchApp {
	return ArchApp{}
}

func (a ArchApp) Analysis(deps []core_domain.CodeDataStruct, identifiersMap map[string]core_domain.CodeDataStruct) *tequila.FullGraph {
	fullGraph := &tequila.FullGraph{
		NodeList:     make(map[string]string),
		RelationList: make(map[string]*tequila.Relation),
	}

	for _, clz := range deps {
		if clz.NodeName == "Main" {
			continue
		}

		src := clz.Package + "." + clz.NodeName
		fullGraph.NodeList[src] = src

		for _, impl := range clz.Implements {
			relation := &tequila.Relation{
				From:  src,
				To:    impl,
				Style: "\"solid\"",
			}

			fullGraph.RelationList[relation.From+"->"+relation.To] = relation
		}

		addCallInField(clz, src, *fullGraph)
		addExtend(clz, src, *fullGraph)
		addCallInMethod(clz, identifiersMap, src, *fullGraph)
	}

	return fullGraph
}

func addCallInField(clz core_domain.CodeDataStruct, src string, fullGraph tequila.FullGraph) {
	for _, field := range clz.FunctionCalls {
		dst := field.Package + "." + field.NodeName
		relation := &tequila.Relation{
			From:  src,
			To:    dst,
			Style: "\"solid\"",
		}

		fullGraph.RelationList[relation.From+"->"+relation.To] = relation
	}
}

func addCallInMethod(clz core_domain.CodeDataStruct, identifiersMap map[string]core_domain.CodeDataStruct, src string, fullGraph tequila.FullGraph) {
	for _, method := range clz.Functions {
		if method.Name == "main" {
			continue
		}

		// TODO: add implements, extends support
		for _, call := range method.FunctionCalls {
			dst := call.Package + "." + call.NodeName
			if src == dst {
				continue
			}

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

func addExtend(clz core_domain.CodeDataStruct, src string, fullGraph tequila.FullGraph) {
	if clz.Extend != "" {
		relation := &tequila.Relation{
			From:  src,
			To:    clz.Extend,
			Style: "\"solid\"",
		}

		fullGraph.RelationList[relation.From+"->"+relation.To] = relation
	}
}
