package evaluate

import (
	"github.com/phodal/coca/core/domain/evaluate/evaluator"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"gonum.org/v1/gonum/stat"
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

		if strings.Contains(strings.ToLower(node.Class), "util") || strings.Contains(strings.ToLower(node.Class), "utils") {
			result.Summary.UtilsCount++

			evaluation = Evaluation{evaluator.Util{}}
			evaluation.Evaluate(&result, node)
		}

		if strings.Contains(strings.ToLower(node.Class), "service") {
			servicesNode = append(servicesNode, node)
		} else {
			evaluation = Evaluation{evaluator.Empty{}}
		}
	}

	var methodLengthArray []float64
	var methodCountArray []float64
	for _, ident := range identifiers {
		result.Summary.ClassCount++

		methodCountArray = append(methodCountArray, float64(len(ident.Methods)))

		for _, method := range ident.Methods {
			result.Summary.MethodCount++

			if support.Contains(method.Modifiers, "static") {
				result.Summary.StaticMethodCount++
			}

			if !strings.HasPrefix(method.Name, "set") && !strings.HasPrefix(method.Name, "get") {
				result.Summary.NormalMethodCount++
				methodLength := method.StopLine - method.StartLine + 1
				result.Summary.TotalMethodLength = result.Summary.TotalMethodLength + methodLength

				methodLengthArray = append(methodLengthArray, float64(methodLength))
			}
		}
	}

	result.Summary.MethodLengthStdDeviation = stat.StdDev(methodLengthArray, nil)
	result.Summary.MethodNumStdDeviation = stat.StdDev(methodCountArray, nil)

	evaluation = Evaluation{evaluator.Service{}}
	evaluation.EvaluateList(&result, servicesNode, nodeMap, identifiers)

	nullableEva := Evaluation{evaluator.NullPointException{}}
	nullableEva.EvaluateList(&result, servicesNode, nodeMap, identifiers)

	return result
}
