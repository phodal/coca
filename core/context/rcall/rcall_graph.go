package rcall

import (
	"github.com/phodal/coca/core/context/call"
	"github.com/phodal/coca/core/domain"
)

type RCallGraph struct {
}

func NewRCallGraph() RCallGraph {
	return *&RCallGraph{}
}

func (c RCallGraph) Analysis(funcName string, clzs []domain.JClassNode, writeCallback func(rcallMap map[string][]string)) string {
	var projectMethodMap = BuildProjectMethodMap(clzs)
	rcallMap := BuildRCallMethodMap(clzs, projectMethodMap)

	writeCallback(rcallMap)

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
			for _, jMethodCall := range method.MethodCalls {
				if jMethodCall.Class != "" {
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

func (c RCallGraph) buildRCallChain(funcName string, methodMap map[string][]string) string {
	if loopCount >= 6 {
		return "\n"
	}
	loopCount++

	if len(methodMap[funcName]) > 0 {
		var arrayResult = ""
		for _, child := range methodMap[funcName] {
			if len(methodMap[child]) > 0 {
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
