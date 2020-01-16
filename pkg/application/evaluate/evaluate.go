package evaluate

import (
	"github.com/phodal/coca/pkg/application/evaluate/evaluator"
	"github.com/phodal/coca/pkg/domain/jdomain"
)

type Evaluator interface {
	Evaluate(result *evaluator.EvaluateModel, node jdomain.JClassNode)
	EvaluateList(evaluateModel *evaluator.EvaluateModel, nodes []jdomain.JClassNode, nodeMap map[string]jdomain.JClassNode, identifiers []jdomain.JIdentifier)
}

type Evaluation struct {
	Evaluator Evaluator
}

func (o *Evaluation) Evaluate(result *evaluator.EvaluateModel, node jdomain.JClassNode) {
	o.Evaluator.Evaluate(result, node)
}

func (o *Evaluation) EvaluateList(evaluateModel *evaluator.EvaluateModel, nodes []jdomain.JClassNode, nodeMap map[string]jdomain.JClassNode, identifiers []jdomain.JIdentifier) {
	o.Evaluator.EvaluateList(evaluateModel, nodes, nodeMap, identifiers)
}
