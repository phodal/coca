package evaluator

import (
	"github.com/phodal/coca/core/domain"
)

type Empty struct {

}

func (Empty) Evaluate(*EvaluateModel, domain.JClassNode) {

}

func (Empty) EvaluateList(*EvaluateModel, []domain.JClassNode, map[string]domain.JClassNode, []domain.JIdentifier) {

}