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

func (a Analyser) Analysis(classNodes []models.JClassNode, identifiers []models.JIdentifier) evaluator.EvaluateModel {
	var servicesNode []models.JClassNode = nil
	var evaluation Evaluation
	var result = evaluator.NewEvaluateModel()

	var nodeMap = make(map[string]models.JClassNode)
	for _, node := range classNodes {
		nodeMap[node.Class] = node
	}

	for _, node := range classNodes {
		if strings.Contains(strings.ToLower(node.Class), "service") {
			servicesNode = append(servicesNode, node)
		} else {
			evaluation = Evaluation{evaluator.Empty{}}
		}
	}

	evaluation = Evaluation{evaluator.Service{}}
	evaluation.EvaluateList(&result, servicesNode, nodeMap, identifiers)

	nullableEva := Evaluation{evaluator.NullPointException{}}
	nullableEva.EvaluateList(&result, servicesNode, nodeMap, identifiers)

	return result
}
