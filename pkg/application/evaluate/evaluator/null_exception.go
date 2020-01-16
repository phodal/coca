	package evaluator

import (
	"github.com/phodal/coca/pkg/domain/jdomain"
)

type NullPointException struct {
}

func (NullPointException) Evaluate(*EvaluateModel, jdomain.JClassNode) {

}

func (n NullPointException) EvaluateList(evaluateModel *EvaluateModel, nodes []jdomain.JClassNode, nodeMap map[string]jdomain.JClassNode, identifiers []jdomain.JIdentifier) {
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

func buildMethodPath(ident jdomain.JIdentifier, method jdomain.JMethod) string {
	return ident.Package + "." + ident.ClassName + "." + method.Name
}
