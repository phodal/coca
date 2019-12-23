package evaluate

import (
	"github.com/phodal/coca/core/domain/evaluate/evaluator"
	"github.com/phodal/coca/core/models"
)

type EvaluateAnalyser struct {

}


func NewEvaluateAnalyser() EvaluateAnalyser {
	return *&EvaluateAnalyser{}
}


func (a EvaluateAnalyser) Analysis(nodes *[]models.JClassNode) {
	evaluation := Evaluation{evaluator.Service{}}
	match := evaluation.IsMatch()
	if match {
		evaluation.Evaluate()
	}

}