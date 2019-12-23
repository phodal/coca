package evaluate

import "github.com/phodal/coca/core/models"

type Evaluator interface {
	Evaluate(node models.JClassNode)
}

type Evaluation struct {
	Evaluator Evaluator
}

func (o *Evaluation) Evaluate(node models.JClassNode) {
	o.Evaluator.Evaluate(node)
}
