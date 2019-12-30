package rcall

import (
	"encoding/json"
	"github.com/phodal/coca/core/context/call"
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure/coca_file"
)

type RCallGraph struct {
}

func NewRCallGraph() RCallGraph {
	return *&RCallGraph{}
}

func (c RCallGraph) Analysis(funcName string, clzs []domain.JClassNode) string {
	var projectMethodMap = BuildProjectMethodMap(clzs)
	rcallMap := BuildRCallMethodMap(clzs, projectMethodMap)

	mapJson, _ := json.MarshalIndent(rcallMap, "", "\t")
	coca_file.WriteToCocaFile("rcallmap.json", string(mapJson))

	chain := c.buildRCallChain(funcName, rcallMap)

	graphvizReverse := "rankdir = LR;\nedge [dir=\"back\"];\n"
	chain = graphvizReverse + chain
	dotContent := call.ToGraphviz(chain)
	return dotContent
}

func BuildProjectMethodMap(clzs []domain.JClassNode) map[string]int {
	var maps = make(map[string]int)
	for _, clz := range clzs {
		for _, method := range clz.Methods {
			maps[method.BuildFullMethodName(clz)] = 1
		}
	}

	return maps
}

func BuildRCallMethodMap(parserDeps []domain.JClassNode, projectMaps map[string]int) map[string][]string {
	var methodMap = make(map[string][]string)
	for _, clz := range parserDeps {
		for _, method := range clz.Methods {
			var caller = method.BuildFullMethodName(clz)
			for _, call := range method.MethodCalls {
				if call.Class != "" {
					callee := call.BuildFullMethodName()
					if projectMaps[callee] < 1 {
						continue
					}
					methodMap[callee] = append(methodMap[callee], caller)
				}
			}
		}
	}

	return methodMap
}

var loopCount = 0

func (c RCallGraph) buildRCallChain(funcName string, methodMap map[string][]string) string {
	if loopCount > 6 {
		return "\n"
	}
	loopCount++

	if len(methodMap[funcName]) > 0 {
		var arrayResult = ""
		for _, child := range methodMap[funcName] {
			if len(methodMap[child]) > 0 {
				arrayResult = arrayResult + c.buildRCallChain(child, methodMap)
			}
			arrayResult = arrayResult + "\"" + funcName + "\" -> \"" + child + "\";\n"
		}

		return arrayResult

	}
	return "\n"
}
