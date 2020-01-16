package evaluator

import (
	"github.com/phodal/coca/pkg/domain/jdomain"
)

type Empty struct {

}

func (Empty) Evaluate(*EvaluateModel, jdomain.JClassNode) {

}

func (Empty) EvaluateList(*EvaluateModel, []jdomain.JClassNode, map[string]jdomain.JClassNode, []jdomain.JIdentifier) {

}