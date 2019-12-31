package bs

import (
	"encoding/json"
	"github.com/phodal/coca/core/domain/bs_domain"
	"strconv"
)

func AnalysisBadSmell(nodes []bs_domain.BsJClass) []bs_domain.BadSmellModel {
	var badSmellList []bs_domain.BadSmellModel
	for _, node := range nodes {
		// To be Defined number
		if node.Type == "Class" && len(node.Methods) < 1 {
			badSmellList = append(badSmellList, *&bs_domain.BadSmellModel{node.Path, "", "lazyElement", "", 0})
		}

		onlyHaveGetterAndSetter := true
		// Long Method
		for _, method := range node.Methods {
			methodLength := method.StopLine - method.StartLine
			if methodLength > 30 {
				description := "method length: " + strconv.Itoa(methodLength)
				longMethod := &bs_domain.BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "longMethod", description, methodLength}
				badSmellList = append(badSmellList, *longMethod)
			}

			if !(method.IsGetterSetter()) {
				onlyHaveGetterAndSetter = false
			}

			// longParameterList
			if len(method.Parameters) > 5 {
				paramsJson, _ := json.Marshal(method.Parameters)
				str := string(paramsJson[:])
				longParams := &bs_domain.BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "longParameterList", str, len(method.Parameters)}
				badSmellList = append(badSmellList, *longParams)
			}

			// repeatedSwitches
			if method.MethodBs.IfSize > 8 {
				longParams := &bs_domain.BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "repeatedSwitches", "ifSize", method.MethodBs.IfSize}
				badSmellList = append(badSmellList, *longParams)
			}

			// repeatedSwitches
			if method.MethodBs.SwitchSize > 8 {
				longParams := &bs_domain.BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "repeatedSwitches", "switchSize", method.MethodBs.SwitchSize}
				badSmellList = append(badSmellList, *longParams)
			}

			// complex if
			for _, info := range method.MethodBs.IfInfo {
				if info.EndLine-info.StartLine >= 2 {
					longParams := &bs_domain.BadSmellModel{node.Path, strconv.Itoa(info.StartLine), "complexCondition", "complexCondition", 0}
					badSmellList = append(badSmellList, *longParams)
				}
			}
		}

		// dataClass
		if onlyHaveGetterAndSetter && node.Type == "Class" && len(node.Methods) > 0 {
			dataClass := &bs_domain.BadSmellModel{node.Path, "", "dataClass", "", len(node.Methods)}
			badSmellList = append(badSmellList, *dataClass)
		}

		//Refused Bequest
		if node.Extends != "" {
			if node.HaveCallParent() {
				badSmellList = append(badSmellList, *&bs_domain.BadSmellModel{node.Path, "", "refusedBequest", "", 0})
			}
		}

		// LargeClass
		normalClassLength := withOutGetterSetterClass(node.Methods)
		if node.Type == "Class" && normalClassLength > 20 {
			description := "methods number (without getter/setter): " + strconv.Itoa(normalClassLength)
			badSmellList = append(badSmellList, *&bs_domain.BadSmellModel{node.Path, "", "largeClass", description, normalClassLength})
		}
	}

	return badSmellList
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
