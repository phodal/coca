package call_graph

import (
	"coca/core/adapter/api"
	"coca/core/models"
	"sort"
	"strings"
)

type CallGraph struct {
}

func NewCallGraph() CallGraph {
	return *&CallGraph{}
}

func (c CallGraph) Analysis(funcName string, clzs []models.JClassNode) string {
	methodMap := BuildMethodMap(clzs)
	chain := BuildCallChain(funcName, methodMap, nil)
	dotContent := ToGraphviz(chain)
	return dotContent
}

// TODO: be a utils
func ToGraphviz(chain string) string {
	//rankdir = LR;
	var result = "digraph G {\n"
	result = result + chain
	result = result + "}\n"
	return result
}

var loopCount = 0

func BuildCallChain(funcName string, methodMap map[string][]string, diMap map[string]string) string {
	if loopCount > 6 {
		return "\n"
	}
	loopCount++

	if len(methodMap[funcName]) > 0 {
		var arrayResult = ""
		for _, child := range methodMap[funcName] {
			if _, ok := diMap[getClz(child)]; ok {
				child = diMap[getClz(child)] + "." + getMethodName(child)
			}
			if len(methodMap[child]) > 0 {
				arrayResult = arrayResult + BuildCallChain(child, methodMap, diMap)
			}
			arrayResult = arrayResult + "\"" + funcName + "\" -> \"" + child + "\";\n"
		}

		return arrayResult

	}
	return "\n"
}

func getClz(child string) string {
	split := strings.Split(child, ".")
	return strings.Join(split[:len(split)-1], ".")
}

func getMethodName(child string) string {
	split := strings.Split(child, ".")
	return strings.Join(split[len(split)-1:], ".")
}

func (c CallGraph) AnalysisByFiles(restApis []api.RestApi, deps []models.JClassNode, diMap map[string]string) (string, []CallApiCount) {
	methodMap := BuildMethodMap(deps)
	var apiCallSCounts []CallApiCount

	results := "digraph G { \n"

	for _, restApi := range restApis {
		caller := restApi.PackageName + "." + restApi.ClassName + "." + restApi.MethodName

		loopCount = 0
		chain := "\"" + restApi.HttpMethod + " " + restApi.Uri + "\" -> \"" + caller + "\";\n"
		apiCallChain := BuildCallChain(caller, methodMap, diMap)
		chain = chain + apiCallChain

		count := &CallApiCount{
			HttpMethod: restApi.HttpMethod,
			Caller:     caller,
			Uri:        restApi.Uri,
			Size:       len(strings.Split(apiCallChain, " -> ")),
		}
		apiCallSCounts = append(apiCallSCounts, *count)

		results = results + "\n" + chain
	}

	sort.Slice(apiCallSCounts, func(i, j int) bool {
		return apiCallSCounts[i].Size < apiCallSCounts[j].Size
	})

	return results + "}\n", apiCallSCounts
}

func BuildMethodMap(clzs []models.JClassNode) map[string][]string {
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
