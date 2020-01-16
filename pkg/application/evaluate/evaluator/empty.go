package evaluator

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
)

type Empty struct {

}

func (Empty) Evaluate(*EvaluateModel, core_domain.CodeDataStruct) {

}

func (Empty) EvaluateList(*EvaluateModel, []core_domain.CodeDataStruct, map[string]core_domain.CodeDataStruct, []core_domain.CodeDataStruct) {

}