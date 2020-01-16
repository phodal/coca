	package evaluator

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
)

type NullPointException struct {
}

func (NullPointException) Evaluate(*EvaluateModel, core_domain.CodeDataStruct) {

}

func (n NullPointException) EvaluateList(evaluateModel *EvaluateModel, nodes []core_domain.CodeDataStruct, nodeMap map[string]core_domain.CodeDataStruct, identifiers []core_domain.CodeDataStruct) {
	var nullableList []string = nil
	var nullableMap = make(map[string]string)
	for _, ident := range identifiers {
		for _, method := range ident.Functions {
			methodName := buildMethodPath(ident, method)
			if method.IsReturnNull {
				nullableMap[methodName] = methodName
			} else {
				for _, annotation := range method.Annotations {
					if annotation.Name == "Nullable" || annotation.Name == "CheckForNull"  {
						nullableMap[methodName] = methodName
					}
				}
			}
		}
	}

	for _, value := range nullableMap {
		nullableList = append(nullableList, value)
	}

	evaluateModel.Nullable.Items = nullableList
}

func buildMethodPath(ident core_domain.CodeDataStruct, method core_domain.CodeFunction) string {
	return ident.Package + "." + ident.NodeName + "." + method.Name
}
