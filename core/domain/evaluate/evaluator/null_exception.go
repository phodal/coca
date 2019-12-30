	package evaluator

import (
	"github.com/phodal/coca/core/models"
)

type NullPointException struct {
}

func (NullPointException) Evaluate(*EvaluateModel, models.JClassNode) {

}

func (n NullPointException) EvaluateList(evaluateModel *EvaluateModel, nodes []models.JClassNode, nodeMap map[string]models.JClassNode, identifiers []models.JIdentifier) {
	var nullableList []string = nil
	var nullableMap = make(map[string]string)
	for _, ident := range identifiers {
		for _, method := range ident.Methods {
			methodName := buildMethodPath(ident, method)
			if method.IsReturnNull {
				nullableMap[methodName] = methodName
			} else {
				for _, annotation := range method.Annotations {
					if annotation.QualifiedName == "Nullable" {
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

func buildMethodPath(ident models.JIdentifier, method models.JMethod) string {
	return ident.Package + "." + ident.ClassName + "." + method.Name
}
