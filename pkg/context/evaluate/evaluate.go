package evaluate

import (
	"github.com/phodal/coca/pkg/context/evaluate/evaluator"
	"github.com/phodal/coca/pkg/domain"
)

type Evaluator interface {
	Evaluate(result *evaluator.EvaluateModel, node domain.JClassNode)
	EvaluateList(evaluateModel *evaluator.EvaluateModel, nodes []domain.JClassNode, nodeMap map[string]domain.JClassNode, identifiers []domain.JIdentifier)
}

type Evaluation struct {
	Evaluator Evaluator
}

func (o *Evaluation) Evaluate(result *evaluator.EvaluateModel, node domain.JClassNode) {
	o.Evaluator.Evaluate(result, node)
}

func (o *Evaluation) EvaluateList(evaluateModel *evaluator.EvaluateModel, nodes []domain.JClassNode, nodeMap map[string]domain.JClassNode, identifiers []domain.JIdentifier) {
	o.Evaluator.EvaluateList(evaluateModel, nodes, nodeMap, identifiers)
}
