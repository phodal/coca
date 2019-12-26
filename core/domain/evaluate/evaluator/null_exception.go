package evaluator

import (
	"github.com/phodal/coca/core/models"
)

type NullException struct {

}

func (NullException) EvaluateList([]models.JClassNode, map[string]models.JClassNode, []models.JIdentifier) {

}
