package evaluate

import (
	"github.com/phodal/coca/pkg/application/evaluate/evaluator"
	"github.com/phodal/coca/pkg/domain/core_domain"
)

type Evaluator interface {
	Evaluate(result *evaluator.EvaluateModel, node core_domain.JClassNode)
	EvaluateList(evaluateModel *evaluator.EvaluateModel, nodes []core_domain.JClassNode, nodeMap map[string]core_domain.JClassNode, identifiers []core_domain.JIdentifier)
}

type Evaluation struct {
	Evaluator Evaluator
}

func (o *Evaluation) Evaluate(result *evaluator.EvaluateModel, node core_domain.JClassNode) {
	o.Evaluator.Evaluate(result, node)
}

func (o *Evaluation) EvaluateList(evaluateModel *evaluator.EvaluateModel, nodes []core_domain.JClassNode, nodeMap map[string]core_domain.JClassNode, identifiers []core_domain.JIdentifier) {
	o.Evaluator.EvaluateList(evaluateModel, nodes, nodeMap, identifiers)
}
