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
	for _, ident := range identifiers {
		for _, method := range ident.Methods {
			if method.IsReturnNull {
				nullableList = append(nullableList, buildMethodPath(ident, method))
			} else {
				for _, annotation := range method.Annotations {
					if annotation.QualifiedName == "Nullable" {
						nullableList = append(nullableList, buildMethodPath(ident, method))
					}
				}
			}
		}
	}

	evaluateModel.Nullable.Items = nullableList
}

func buildMethodPath(ident models.JIdentifier, method models.JMethod) string {
	return ident.Package + "." + ident.ClassName + "." + method.Name
}
