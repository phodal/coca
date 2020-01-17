package bs

import (
	"encoding/json"
	"github.com/huleTW/bad-smell-analysis/graphcall"
	"github.com/phodal/coca/pkg/domain/bs_domain"
	"strconv"
)

var (
	BS_LONG_PARAS_LENGTH = 5
	BS_IF_SWITCH_LENGTH  = 8
	BS_LARGE_LENGTH      = 20
	BS_METHOD_LENGTH     = 30
	BS_IF_LINES_LENGTH   = 3
)

func AnalysisBadSmell(nodes []bs_domain.BsJClass) []bs_domain.BadSmellModel {
	var badSmellList []bs_domain.BadSmellModel
	for _, node := range nodes {
		checkLazyElement(node, &badSmellList)

		onlyHaveGetterAndSetter := true
		for _, method := range node.Methods {
			checkLongMethod(method, node, &badSmellList)

			if !(method.IsGetterSetter()) {
				onlyHaveGetterAndSetter = false
			}

			checkLongParameterList(method, node, &badSmellList)
			checkRepeatedSwitches(method, node, &badSmellList)
			checkComplexIf(method, node, &badSmellList)
		}

		checkDataClass(onlyHaveGetterAndSetter, node, &badSmellList)
		checkRefusedBequest(node, &badSmellList)
		checkLargeClass(node, &badSmellList)
	}
	checkConnectedGraphCall(nodes, &badSmellList)
	return badSmellList
}

func checkConnectedGraphCall(nodes []bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	var classNodes = map[string][]string{}
	var classNodeMaps = map[string]bool{}
	for _, node := range nodes {
		classNodeMaps[node.ClassFullName()] = true
	}
	for _, node := range nodes {
		classNodes[node.ClassFullName()] = getCalledClasses(node, classNodeMaps)
	}
	var badSmellGraphCall = graphcall.NewBadSmellGraphCall()
	var descriptions = badSmellGraphCall.AnalysisGraphCallPath(classNodes)
	for _, description := range descriptions {
		*badSmellList = append(*badSmellList, bs_domain.BadSmellModel{Bs: "graphConnectedCall", Description: description})
	}
}

//fixme java lamda & recursive
func getCalledClasses(class bs_domain.BsJClass, maps map[string]bool) []string {
	var calledClassesMap = make(map[string]struct{})
	var calledClasses []string
	for _, methodCalled := range class.MethodCalls {
		if methodCalled.NodeName == "" || !maps[methodCalled.BuildClassFullName()] || class.ClassFullName() == methodCalled.BuildClassFullName() {
			continue
		}
		calledClassesMap[methodCalled.BuildClassFullName()] = struct{}{}
	}
	for key := range calledClassesMap {
		calledClasses = append(calledClasses, key)
	}
	return calledClasses
}

func checkLazyElement(node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	if node.Type == "NodeName" && len(node.Methods) < 1 {
		*badSmellList = append(*badSmellList, bs_domain.BadSmellModel{File: node.Path, Bs: "lazyElement"})
	}
}

func checkLongMethod(method bs_domain.BsJMethod, node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	methodLength := method.Position.StopLine - method.Position.StartLine

	if methodLength > BS_METHOD_LENGTH {
		description := "method length: " + strconv.Itoa(methodLength)
		longMethod := &bs_domain.BadSmellModel{File: node.Path, Line: strconv.Itoa(method.Position.StartLine), Bs: "longMethod", Description: description, Size: methodLength}
		*badSmellList = append(*badSmellList, *longMethod)
	}
}

func checkDataClass(onlyHaveGetterAndSetter bool, node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	if onlyHaveGetterAndSetter && node.Type == "NodeName" && len(node.Methods) > 0 {
		dataClass := &bs_domain.BadSmellModel{File: node.Path, Bs: "dataClass", Size: len(node.Methods)}
		*badSmellList = append(*badSmellList, *dataClass)
	}
}

func checkRefusedBequest(node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	if node.Extends != "" {
		if node.HaveCallParent() {
			*badSmellList = append(*badSmellList, bs_domain.BadSmellModel{File: node.Path, Bs: "refusedBequest"})
		}
	}
}

func checkLargeClass(node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	normalClassLength := withOutGetterSetterClass(node.Methods)
	if node.Type == "NodeName" && normalClassLength >= BS_LARGE_LENGTH {
		description := "methods number (without getter/setter): " + strconv.Itoa(normalClassLength)
		*badSmellList = append(*badSmellList, bs_domain.BadSmellModel{File: node.Path, Bs: "largeClass", Description: description, Size: normalClassLength})
	}
}

func checkComplexIf(method bs_domain.BsJMethod, node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	for _, info := range method.MethodBs.IfInfo {
		if info.EndLine-info.StartLine >= BS_IF_LINES_LENGTH {
			longParams := &bs_domain.BadSmellModel{File: node.Path, Line: strconv.Itoa(info.StartLine), Bs: "complexCondition", Description: "complexCondition"}
			*badSmellList = append(*badSmellList, *longParams)
		}
	}
}

func checkRepeatedSwitches(method bs_domain.BsJMethod, node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	if method.MethodBs.IfSize >= BS_IF_SWITCH_LENGTH {
		longParams := &bs_domain.BadSmellModel{File: node.Path, Line: strconv.Itoa(method.Position.StartLine), Bs: "repeatedSwitches", Description: "ifSize", Size: method.MethodBs.IfSize}
		*badSmellList = append(*badSmellList, *longParams)
	}

	if method.MethodBs.SwitchSize >= BS_IF_SWITCH_LENGTH {
		longParams := &bs_domain.BadSmellModel{File: node.Path, Line: strconv.Itoa(method.Position.StartLine), Bs: "repeatedSwitches", Description: "switchSize", Size: method.MethodBs.SwitchSize}
		*badSmellList = append(*badSmellList, *longParams)
	}
}

func checkLongParameterList(method bs_domain.BsJMethod, node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	if len(method.Parameters) > BS_LONG_PARAS_LENGTH {
		paramsJson, _ := json.Marshal(method.Parameters)
		str := string(paramsJson[:])
		longParams := &bs_domain.BadSmellModel{File: node.Path, Line: strconv.Itoa(method.Position.StartLine), Bs: "longParameterList", Description: str, Size: len(method.Parameters)}
		*badSmellList = append(*badSmellList, *longParams)
	}
}

func withOutGetterSetterClass(fullMethods []bs_domain.BsJMethod) int {
	var normalMethodSize = 0
	for _, method := range fullMethods {
		if !(method.IsGetterSetter()) {
			normalMethodSize++
		}
	}

	return normalMethodSize
}
