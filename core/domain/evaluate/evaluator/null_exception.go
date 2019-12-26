package evaluator

import (
	"github.com/phodal/coca/core/models"
)

type NullException struct {
}

func (NullException) Evaluate(EvaluateModel, models.JClassNode) {

}

func (n NullException) EvaluateList(evaluateModel EvaluateModel, nodes []models.JClassNode, classNodeMap map[string]models.JClassNode, identifiers []models.JIdentifier) {
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
