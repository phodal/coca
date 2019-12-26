package count

import "github.com/phodal/coca/core/models"

func BuildCallMap(parserDeps []models.JClassNode) map[string]int {
	var projectMethods = make(map[string]bool)
	for _, clz := range parserDeps {
		for _, method := range clz.Methods {
			projectMethods[clz.Package+"."+clz.Class+"."+method.Name] = true
		}
	}

	// TODO: support identify data class
	var callMap = make(map[string]int)
	for _, clz := range parserDeps {
		for _, call := range clz.MethodCalls {
			callMethod := call.Package + "." + call.Class + "." + call.MethodName
			if projectMethods[callMethod] {
				if callMap[callMethod] == 0 {
					callMap[callMethod] = 1
				} else {
					callMap[callMethod]++
				}
			}
		}
	}

	return callMap
}
