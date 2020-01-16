package evaluator

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
)

type Util struct {
}

func (Util) Evaluate(result *EvaluateModel, node core_domain.JClassNode) {

}

func (s Util) EvaluateList(evaluateModel *EvaluateModel, nodes []core_domain.JClassNode, nodeMap map[string]core_domain.JClassNode, identifiers []core_domain.JIdentifier) {

}
