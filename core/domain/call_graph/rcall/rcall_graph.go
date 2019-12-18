package rcall

import (
	"coca/core/domain/call_graph"
	"coca/core/models"
)

type RCallGraph struct {
}

func NewRCallGraph() RCallGraph {
	return *&RCallGraph{}
}

func (c RCallGraph) Analysis(funcName string, clzs []models.JClassNode) string {
	methodMap := BuildRCallMethodMap(clzs)

	chain := c.buildRCallChain(funcName, methodMap)

	dotContent := call_graph.ToGraphviz(chain)
	return dotContent
}

func BuildRCallMethodMap(clzs []models.JClassNode) map[string][]string {
	var methodMap = make(map[string][]string)
	for _, clz := range clzs {
		for _, method := range clz.Methods {
			var caller = clz.Package + "." + clz.Class + "." + method.Name
			for _, call := range method.MethodCalls {
				if call.Class != "" {
					callee := call.Package + "." + call.Class + "." + call.MethodName
					if len(methodMap[callee]) == 0 {
						methodMap[callee] = append(methodMap[callee], caller)
					}
					for _, cacheCaller := range methodMap[callee] {
						if cacheCaller != caller {
							methodMap[callee] = append(methodMap[callee], caller)
						}
					}
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
