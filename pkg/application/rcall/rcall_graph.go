package rcall

import (
	"github.com/phodal/coca/pkg/application/call"
	"github.com/phodal/coca/pkg/domain/core_domain"
)

type RCallGraph struct {
}

func NewRCallGraph() RCallGraph {
	return RCallGraph{}
}

func (c RCallGraph) Analysis(funcName string, clzs []core_domain.CodeDataStruct, writeCallback func(rcallMap map[string][]string)) string {
	var projectMethodMap = BuildProjectMethodMap(clzs)
	rcallMap := BuildRCallMethodMap(clzs, projectMethodMap)

	writeCallback(rcallMap)

	chain := c.buildRCallChain(funcName, rcallMap)

	graphvizReverse := "rankdir = LR;\nedge [dir=\"back\"];\n"
	chain = graphvizReverse + chain
	dotContent := call.ToGraphviz(chain)
	return dotContent
}

func BuildProjectMethodMap(clzs []core_domain.CodeDataStruct) map[string]int {
	var maps = make(map[string]int)
	for _, clz := range clzs {
		for _, method := range clz.Functions {
			maps[method.BuildFullMethodName(clz)] = 1
		}
	}

	return maps
}

func BuildRCallMethodMap(parserDeps []core_domain.CodeDataStruct, projectMaps map[string]int) map[string][]string {
	var methodMap = make(map[string][]string)
	for _, clz := range parserDeps {
		for _, method := range clz.Functions {
			var caller = method.BuildFullMethodName(clz)
			for _, jMethodCall := range method.FunctionCalls {
				if jMethodCall.NodeName != "" {
					callee := jMethodCall.BuildFullMethodName()
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
var lastChild = ""

func (c RCallGraph) buildRCallChain(funcName string, methodMap map[string][]string) string {
	if loopCount >= 6 {
		return "\n"
	}
	loopCount++

	if len(methodMap[funcName]) > 0 {
		var arrayResult = ""
		for _, child := range methodMap[funcName] {
			if child == lastChild {
				return ""
			}
			if len(methodMap[child]) > 0 {
				lastChild = child
				arrayResult = arrayResult + c.buildRCallChain(child, methodMap)
			}
			if funcName == child {
				continue
			}
			newCall := "\"" + funcName + "\" -> \"" + child + "\";\n"
			arrayResult = arrayResult + newCall
		}

		return arrayResult

	}
	return "\n"
}
