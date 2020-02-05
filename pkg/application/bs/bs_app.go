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

	SMELL_GARPH_CONNECTED_CALL = "graphConnectedCall"
	SMELL_LAZY_ELEMENT         = "lazyElement"
	SMELL_LONG_METHOD          = "longMethod"
	SMELL_DATA_CLASS           = "dataClass"
	SMELL_REFUSED_BEQUEST      = "refusedBequest"
	SMELL_LARGE_CLASS          = "largeClass"
	SMELL_COMPLEX_CONDITION    = "complexCondition"
	SMELL_REPEATED_SWITCHES    = "repeatedSwitches"
	SMELL_LONG_PARAMETER_LIST  = "longParameterList"
)

func AnalysisBadSmell(nodes []bs_domain.BSDataStruct) []bs_domain.BadSmellModel {
	var badSmellList []bs_domain.BadSmellModel
	for _, node := range nodes {
		checkLazyElement(node, &badSmellList)

		onlyHaveGetterAndSetter := true
		for _, method := range node.Functions {
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

func checkConnectedGraphCall(nodes []bs_domain.BSDataStruct, badSmellList *[]bs_domain.BadSmellModel) {
	var classNodes = map[string][]string{}
	var classNodeMaps = map[string]bool{}
	for _, node := range nodes {
		classNodeMaps[node.GetClassFullName()] = true
	}
	for _, node := range nodes {
		classNodes[node.GetClassFullName()] = bs_domain.GetCalledClasses(node, classNodeMaps)
	}
	var badSmellGraphCall = graphcall.NewBadSmellGraphCall()
	var descriptions = badSmellGraphCall.AnalysisGraphCallPath(classNodes)
	for _, description := range descriptions {
		*badSmellList = append(*badSmellList, bs_domain.BadSmellModel{Bs: SMELL_GARPH_CONNECTED_CALL, Description: description})
	}
}

func checkLazyElement(node bs_domain.BSDataStruct, badSmellList *[]bs_domain.BadSmellModel) {
	if node.Type == "Class" && len(node.Functions) < 1 {
		*badSmellList = append(*badSmellList, bs_domain.BadSmellModel{File: node.FilePath, Bs: SMELL_LAZY_ELEMENT})
	}
}

func checkLongMethod(method bs_domain.BSFunction, node bs_domain.BSDataStruct, badSmellList *[]bs_domain.BadSmellModel) {
	methodLength := method.Position.StopLine - method.Position.StartLine

	if methodLength > BS_METHOD_LENGTH {
		description := "method length: " + strconv.Itoa(methodLength)
		longMethod := bs_domain.BadSmellModel{File: node.FilePath, Line: strconv.Itoa(method.Position.StartLine), Bs: SMELL_LONG_METHOD, Description: description, Size: methodLength}
		*badSmellList = append(*badSmellList, longMethod)
	}
}

func checkDataClass(onlyHaveGetterAndSetter bool, node bs_domain.BSDataStruct, badSmellList *[]bs_domain.BadSmellModel) {
	if onlyHaveGetterAndSetter && node.Type == "Class" && len(node.Functions) > 0 {
		dataClass := bs_domain.BadSmellModel{File: node.FilePath, Bs: SMELL_DATA_CLASS, Size: len(node.Functions)}
		*badSmellList = append(*badSmellList, dataClass)
	}
}

func checkRefusedBequest(node bs_domain.BSDataStruct, badSmellList *[]bs_domain.BadSmellModel) {
	if node.Extend != "" {
		if node.HasCallSuper() {
			*badSmellList = append(*badSmellList, bs_domain.BadSmellModel{File: node.FilePath, Bs: SMELL_REFUSED_BEQUEST})
		}
	}
}

func checkLargeClass(node bs_domain.BSDataStruct, badSmellList *[]bs_domain.BadSmellModel) {
	normalClassLength := bs_domain.WithoutGetterSetterClass(node.Functions)
	if node.Type == "Class" && normalClassLength >= BS_LARGE_LENGTH {
		description := "methods number (without getter/setter): " + strconv.Itoa(normalClassLength)
		*badSmellList = append(*badSmellList, bs_domain.BadSmellModel{File: node.FilePath, Bs: SMELL_LARGE_CLASS, Description: description, Size: normalClassLength})
	}
}

func checkComplexIf(method bs_domain.BSFunction, node bs_domain.BSDataStruct, badSmellList *[]bs_domain.BadSmellModel) {
	for _, info := range method.FunctionBS.IfInfo {
		if info.EndLine-info.StartLine >= BS_IF_LINES_LENGTH {
			longParams := bs_domain.BadSmellModel{File: node.FilePath, Line: strconv.Itoa(info.StartLine), Bs: SMELL_COMPLEX_CONDITION, Description: SMELL_COMPLEX_CONDITION}
			*badSmellList = append(*badSmellList, longParams)
		}
	}
}

func checkRepeatedSwitches(method bs_domain.BSFunction, node bs_domain.BSDataStruct, badSmellList *[]bs_domain.BadSmellModel) {
	if method.FunctionBS.IfSize >= BS_IF_SWITCH_LENGTH {
		longParams := bs_domain.BadSmellModel{File: node.FilePath, Line: strconv.Itoa(method.Position.StartLine), Bs: SMELL_REPEATED_SWITCHES, Description: "ifSize", Size: method.FunctionBS.IfSize}
		*badSmellList = append(*badSmellList, longParams)
	}

	if method.FunctionBS.SwitchSize >= BS_IF_SWITCH_LENGTH {
		longParams := bs_domain.BadSmellModel{File: node.FilePath, Line: strconv.Itoa(method.Position.StartLine), Bs: SMELL_REPEATED_SWITCHES, Description: "switchSize", Size: method.FunctionBS.SwitchSize}
		*badSmellList = append(*badSmellList, longParams)
	}
}

func checkLongParameterList(method bs_domain.BSFunction, node bs_domain.BSDataStruct, badSmellList *[]bs_domain.BadSmellModel) {
	if len(method.Parameters) > BS_LONG_PARAS_LENGTH {
		paramsJson, _ := json.Marshal(method.Parameters)
		str := string(paramsJson[:])
		longParams := bs_domain.BadSmellModel{File: node.FilePath, Line: strconv.Itoa(method.Position.StartLine), Bs: SMELL_LONG_PARAMETER_LIST, Description: str, Size: len(method.Parameters)}
		*badSmellList = append(*badSmellList, longParams)
	}
}
