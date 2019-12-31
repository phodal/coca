package bs

import (
	"encoding/json"
	"github.com/phodal/coca/core/domain/bs_domain"
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
		// To be Defined number
		checkLazyElement(node, &badSmellList)

		onlyHaveGetterAndSetter := true
		// Long Method
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

	return badSmellList
}

func checkLazyElement(node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	if node.Type == "Class" && len(node.Methods) < 1 {
		*badSmellList = append(*badSmellList, *&bs_domain.BadSmellModel{node.Path, "", "lazyElement", "", 0})
	}
}

func checkLongMethod(method bs_domain.BsJMethod, node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	methodLength := method.StopLine - method.StartLine

	if methodLength > BS_METHOD_LENGTH {
		description := "method length: " + strconv.Itoa(methodLength)
		longMethod := &bs_domain.BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "longMethod", description, methodLength}
		*badSmellList = append(*badSmellList, *longMethod)
	}
}

func checkDataClass(onlyHaveGetterAndSetter bool, node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	if onlyHaveGetterAndSetter && node.Type == "Class" && len(node.Methods) > 0 {
		dataClass := &bs_domain.BadSmellModel{node.Path, "", "dataClass", "", len(node.Methods)}
		*badSmellList = append(*badSmellList, *dataClass)
	}
}

func checkRefusedBequest(node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	if node.Extends != "" {
		if node.HaveCallParent() {
			*badSmellList = append(*badSmellList, *&bs_domain.BadSmellModel{node.Path, "", "refusedBequest", "", 0})
		}
	}
}

func checkLargeClass(node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	normalClassLength := withOutGetterSetterClass(node.Methods)
	if node.Type == "Class" && normalClassLength >= BS_LARGE_LENGTH {
		description := "methods number (without getter/setter): " + strconv.Itoa(normalClassLength)
		*badSmellList = append(*badSmellList, *&bs_domain.BadSmellModel{node.Path, "", "largeClass", description, normalClassLength})
	}
}

func checkComplexIf(method bs_domain.BsJMethod, node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	for _, info := range method.MethodBs.IfInfo {
		if info.EndLine-info.StartLine >= BS_IF_LINES_LENGTH {
			longParams := &bs_domain.BadSmellModel{node.Path, strconv.Itoa(info.StartLine), "complexCondition", "complexCondition", 0}
			*badSmellList = append(*badSmellList, *longParams)
		}
	}
}

func checkRepeatedSwitches(method bs_domain.BsJMethod, node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	if method.MethodBs.IfSize >= BS_IF_SWITCH_LENGTH {
		longParams := &bs_domain.BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "repeatedSwitches", "ifSize", method.MethodBs.IfSize}
		*badSmellList = append(*badSmellList, *longParams)
	}

	// repeatedSwitches
	if method.MethodBs.SwitchSize >= BS_IF_SWITCH_LENGTH {
		longParams := &bs_domain.BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "repeatedSwitches", "switchSize", method.MethodBs.SwitchSize}
		*badSmellList = append(*badSmellList, *longParams)
	}
}

func checkLongParameterList(method bs_domain.BsJMethod, node bs_domain.BsJClass, badSmellList *[]bs_domain.BadSmellModel) {
	if len(method.Parameters) > BS_LONG_PARAS_LENGTH {
		paramsJson, _ := json.Marshal(method.Parameters)
		str := string(paramsJson[:])
		longParams := &bs_domain.BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "longParameterList", str, len(method.Parameters)}
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
