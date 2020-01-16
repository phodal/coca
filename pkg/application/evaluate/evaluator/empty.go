package evaluator

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
)

type Empty struct {

}

func (Empty) Evaluate(*EvaluateModel, core_domain.JClassNode) {

}

func (Empty) EvaluateList(*EvaluateModel, []core_domain.JClassNode, map[string]core_domain.JClassNode, []core_domain.JIdentifier) {

}