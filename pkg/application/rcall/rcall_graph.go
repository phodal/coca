package rcall

import (
	"github.com/modernizing/coca/pkg/domain/core_domain"
)

type RCallGraph struct {
}

func NewRCallGraph() RCallGraph {
	return RCallGraph{}
}

func (c RCallGraph) Analysis(funcName string, clzs []core_domain.CodeDataStruct, writeCallback func(rcallMap map[string][]string)) string {
	var projectMethodMap = BuildProjectMethodMap(clzs)
	methodCallMap := BuildMethodCallMap(clzs, projectMethodMap)

	writeCallback(methodCallMap)

	chain := c.BuildRCallChain(funcName, methodCallMap)
	dotContent := ToGraphviz(chain)
	return dotContent
}

// TODO: be a utils
func ToGraphviz(chain string) string {
	var result = "digraph G {\n"
	//result += "rankdir = LR;\n"
	result = result + chain
	result = result + "}\n"
	return result
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

func BuildMethodCallMap(dataStructs []core_domain.CodeDataStruct, projectMaps map[string]int) map[string][]string {
	var methodCallMap = make(map[string][]string)
	for _, clz := range dataStructs {
		for _, method := range clz.Functions {
			var caller = method.BuildFullMethodName(clz)
			for _, jMethodCall := range method.FunctionCalls {
				if jMethodCall.NodeName != "" {
					callee := jMethodCall.BuildFullMethodName()
					if projectMaps[callee] < 1 {
						continue
					}
					methodCallMap[callee] = append(methodCallMap[callee], caller)
				}
			}
		}
	}

	return methodCallMap
}

var loopCount = 0
var lastChild = ""

func (c RCallGraph) BuildRCallChain(funcName string, methodMap map[string][]string) string {
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
				arrayResult = arrayResult + c.BuildRCallChain(child, methodMap)
			}
			if funcName == child {
				continue
			}
			newCall := "\"" + child + "\" -> \"" + funcName + "\";\n"
			arrayResult = arrayResult + newCall
		}

		return arrayResult

	}
	return "\n"
}
