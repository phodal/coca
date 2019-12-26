package evaluator

import (
	"github.com/phodal/coca/core/models"
)

type Empty struct {

}

func (Empty) Evaluate(models.JClassNode) {

}

func (Empty) EvaluateList([]models.JClassNode, map[string]models.JClassNode, []models.JIdentifier) {

}