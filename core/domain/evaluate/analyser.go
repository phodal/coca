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
	var servicesNode []models.JClassNode = nil
	var evaluation Evaluation

	for _, node := range nodes {
		if strings.Contains(strings.ToLower(node.Class), "service") {
			servicesNode = append(servicesNode, node)
		} else {
			evaluation = Evaluation{evaluator.Empty{}}
		}
	}

	evaluation = Evaluation{evaluator.Service{}}
	evaluation.EvaluateList(servicesNode)
}
