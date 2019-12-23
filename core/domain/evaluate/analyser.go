package evaluate

import (
	"github.com/phodal/coca/core/domain/evaluate/evaluator"
	"github.com/phodal/coca/core/models"
	"strings"
)

type Analyser struct {
}

func NewEvaluateAnalyser() Analyser {
	return *&Analyser{}
}

func (a Analyser) Analysis(nodes []models.JClassNode) {
	for _, node := range nodes {
		var evaluation Evaluation
		if strings.Contains(strings.ToLower(node.Class), "service") {
			evaluation = Evaluation{evaluator.Service{}}
		} else {
			evaluation = Evaluation{evaluator.Empty{}}
		}

		evaluation.Evaluate(node)
	}
}
