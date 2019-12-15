package domain

import (
	"coca/src/adapter/models"
)

type CallGraph struct {
}

func NewCallGraph() CallGraph {
	return *&CallGraph{}
}

func (c CallGraph) Analysis(funcName string, clzs []models.JClassNode) string {
	methodMap := c.BuildMethodMap(clzs)

	BuildCallChain(funcName, methodMap)
	chain := BuildCallChain(funcName, methodMap)
	dotContent := toGraphviz(chain)
	return dotContent
}

func toGraphviz(chain string) string {
	var result = "digraph G { \n"
	result = result + chain
	result = result + "}\n"
	return result
}

func BuildCallChain(funcName string, methodMap map[string][]string) string {
	if len(methodMap[funcName]) > 0 {
		var arrayResult = ""
		for _, child := range methodMap[funcName] {
			if len(methodMap[child]) > 0 {
				arrayResult = arrayResult + BuildCallChain(child, methodMap)
			}
			arrayResult = arrayResult + "\"" + funcName + "\" -> \"" + child + "\";\n"
		}

		return arrayResult

	}
	return "\n"
}

func (c CallGraph) BuildMethodMap(clzs []models.JClassNode) map[string][]string {
	var methodMap = make(map[string][]string)
	for _, clz := range clzs {
		for _, method := range clz.Methods {
			var methodName = clz.Package + "." + clz.Class + "." + method.Name
			var calls []string
			for _, call := range method.MethodCalls {
				if call.Class != "" {
					calls = append(calls, call.Package+"."+call.Class+"."+call.MethodName)
				}
			}
			methodMap[methodName] = calls
		}
	}

	return methodMap
}
