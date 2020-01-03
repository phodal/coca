	package evaluator

import (
	"github.com/phodal/coca/core/domain"
)

type NullPointException struct {
}

func (NullPointException) Evaluate(*EvaluateModel, domain.JClassNode) {

}

func (n NullPointException) EvaluateList(evaluateModel *EvaluateModel, nodes []domain.JClassNode, nodeMap map[string]domain.JClassNode, identifiers []domain.JIdentifier) {
	var nullableList []string = nil
	var nullableMap = make(map[string]string)
	for _, ident := range identifiers {
		for _, method := range ident.Methods {
			methodName := buildMethodPath(ident, method)
			if method.IsReturnNull {
				nullableMap[methodName] = methodName
			} else {
				for _, annotation := range method.Annotations {
					if annotation.QualifiedName == "Nullable" || annotation.QualifiedName == "CheckForNull"  {
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

func buildMethodPath(ident domain.JIdentifier, method domain.JMethod) string {
	return ident.Package + "." + ident.ClassName + "." + method.Name
}
