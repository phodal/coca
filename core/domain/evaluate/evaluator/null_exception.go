package evaluator

import (
	"github.com/phodal/coca/core/models"
)

type NullPointException struct {
}

func (NullPointException) Evaluate(EvaluateModel, models.JClassNode) {

}

func (n NullPointException) EvaluateList(evaluateModel *EvaluateModel, nodes []models.JClassNode, nodeMap map[string]models.JClassNode, identifiers []models.JIdentifier) {
	var nullableList []string
	for _, ident := range identifiers {
		for _, method := range ident.Methods {
			for _, annotation := range method.Annotations {
				if annotation.QualifiedName == "Nullable" {
					methodPath := ident.Package + "." + ident.ClassName + "." + method.Name
					nullableList = append(nullableList, methodPath)
				}
			}
		}
	}

	evaluateModel.Nullable.Items = nullableList
}
