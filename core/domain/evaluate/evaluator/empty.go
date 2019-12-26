package evaluator

import (
	"github.com/phodal/coca/core/models"
)

type Empty struct {

}

func (Empty) Evaluate(EvaluateModel, models.JClassNode) {

}

func (Empty) EvaluateList(*EvaluateModel, []models.JClassNode, map[string]models.JClassNode, []models.JIdentifier) {

}