package evaluate

import (
	"github.com/phodal/coca/core/domain/evaluate/evaluator"
	"github.com/phodal/coca/core/models"
)

type Evaluator interface {
	Evaluate(result evaluator.EvaluateModel, node models.JClassNode)
	EvaluateList(evaluateModel *evaluator.EvaluateModel, nodes []models.JClassNode, nodeMap map[string]models.JClassNode, identifiers []models.JIdentifier)
}

type Evaluation struct {
	Evaluator Evaluator
}

func (o *Evaluation) Evaluate(result evaluator.EvaluateModel, node models.JClassNode) {
	o.Evaluator.Evaluate(result, node)
}

func (o *Evaluation) EvaluateList(evaluateModel *evaluator.EvaluateModel, nodes []models.JClassNode, nodeMap map[string]models.JClassNode, identifiers []models.JIdentifier) {
	o.Evaluator.EvaluateList(evaluateModel, nodes, nodeMap, identifiers)
}
