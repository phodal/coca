package evaluate

import "github.com/phodal/coca/core/models"

type Evaluator interface {
	Evaluate(node models.JClassNode)
	EvaluateList(nodes []models.JClassNode, nodeMap map[string]models.JClassNode, identifiers []models.JIdentifier)
}

type Evaluation struct {
	Evaluator Evaluator
}

func (o *Evaluation) Evaluate(node models.JClassNode) {
	o.Evaluator.Evaluate(node)
}

func (o *Evaluation) EvaluateList(nodes []models.JClassNode, nodeMap map[string]models.JClassNode, identifiers []models.JIdentifier) {
	o.Evaluator.EvaluateList(nodes, nodeMap, nil)
}
