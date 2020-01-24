package call

import (
	api_domain2 "github.com/phodal/coca/pkg/domain/api_domain"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/jpackage"
	"strings"
)

type CallGraph struct {
}

func NewCallGraph() CallGraph {
	return CallGraph{}
}

func (c CallGraph) Analysis(funcName string, clzs []core_domain.CodeDataStruct) string {
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
			if _, ok := diMap[jpackage.GetClassName(child)]; ok {
				child = diMap[jpackage.GetClassName(child)] + "." + jpackage.GetMethodName(child)
			}
			if len(methodMap[child]) > 0 {
				arrayResult = arrayResult + BuildCallChain(child, methodMap, diMap)
			}
			arrayResult = arrayResult + "\"" + escapeStr(funcName) + "\" -> \"" + escapeStr(child) + "\";\n"
		}

		return arrayResult

	}
	return "\n"
}

func (c CallGraph) AnalysisByFiles(restApis []api_domain2.RestAPI, deps []core_domain.CodeDataStruct, diMap map[string]string) (string, []api_domain2.CallAPI) {
	methodMap := BuildMethodMap(deps)
	var apiCallSCounts []api_domain2.CallAPI

	results := "digraph G { \n"

	for _, restApi := range restApis {
		caller := restApi.BuildFullMethodPath()

		loopCount = 0
		chain := "\"" + restApi.HttpMethod + " " + restApi.Uri + "\" -> \"" + escapeStr(caller) + "\";\n"
		apiCallChain := BuildCallChain(caller, methodMap, diMap)
		chain = chain + apiCallChain

		count := &api_domain2.CallAPI{
			HTTPMethod: restApi.HttpMethod,
			Caller:     caller,
			URI:        restApi.Uri,
			Size:       len(strings.Split(apiCallChain, " -> ")),
		}
		apiCallSCounts = append(apiCallSCounts, *count)

		results = results + "\n" + chain
	}

	return results + "}\n", apiCallSCounts
}

func escapeStr(caller string) string {
	return strings.ReplaceAll(caller, "\"", "\\\"")
}

func BuildMethodMap(clzs []core_domain.CodeDataStruct) map[string][]string {
	var methodMap = make(map[string][]string)
	for _, clz := range clzs {
		for _, method := range clz.Functions {
			methodName := method.BuildFullMethodName(clz)
			methodMap[methodName] = method.GetAllCallString()
		}
	}

	return methodMap
}
