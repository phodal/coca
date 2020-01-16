	package evaluator

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
)

type NullPointException struct {
}

func (NullPointException) Evaluate(*EvaluateModel, core_domain.JClassNode) {

}

func (n NullPointException) EvaluateList(evaluateModel *EvaluateModel, nodes []core_domain.JClassNode, nodeMap map[string]core_domain.JClassNode, identifiers []core_domain.JIdentifier) {
	var nullableList []string = nil
	var nullableMap = make(map[string]string)
	for _, ident := range identifiers {
		for _, method := range ident.Methods {
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

func buildMethodPath(ident core_domain.JIdentifier, method core_domain.JMethod) string {
	return ident.Package + "." + ident.ClassName + "." + method.Name
}
