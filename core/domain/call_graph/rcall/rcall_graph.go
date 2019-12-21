package rcall

import (
	"github.com/phodal/coca/core/domain/call_graph"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"encoding/json"
)

type RCallGraph struct {
}

func NewRCallGraph() RCallGraph {
	return *&RCallGraph{}
}

func (c RCallGraph) Analysis(funcName string, clzs []models.JClassNode) string {
	var projectMethodMap map[string]int = BuildProjectMethodMap(clzs)
	rcallMap := BuildRCallMethodMap(clzs, projectMethodMap)


	mapJson, _ := json.MarshalIndent(rcallMap, "", "\t")
	support.WriteToCocaFile("rcallmap.json", string(mapJson))

	chain := c.buildRCallChain(funcName, rcallMap)

	graphvizReverse := "rankdir = LR;\nedge [dir=\"back\"];\n"
	chain = graphvizReverse + chain
	dotContent := call_graph.ToGraphviz(chain)
	return dotContent
}

func BuildProjectMethodMap(clzs []models.JClassNode) map[string]int {
	var maps = make(map[string]int)
	for _, clz := range clzs {
		for _, method := range clz.Methods {
			maps[clz.Package+"."+clz.Class+"."+method.Name] = 1
		}
	}

	return maps
}

func BuildRCallMethodMap(clzs []models.JClassNode, projectMaps map[string]int) map[string][]string {
	var methodMap = make(map[string][]string)
	for _, clz := range clzs {
		for _, method := range clz.Methods {
			var caller = clz.Package + "." + clz.Class + "." + method.Name
			for _, call := range method.MethodCalls {
				if call.Class != "" {
					callee := buildMethodFullName(call)
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

func buildMethodFullName(call models.JMethodCall) string {
	return call.Package + "." + call.Class + "." + call.MethodName
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
