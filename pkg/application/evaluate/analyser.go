package evaluate

import (
	"github.com/phodal/coca/pkg/application/evaluate/evaluator"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"gonum.org/v1/gonum/stat"
)

type Analyser struct {
}

func NewEvaluateAnalyser() Analyser {
	return Analyser{}
}

func (a Analyser) Analysis(classNodes []core_domain.CodeDataStruct, identifiers []core_domain.CodeDataStruct) evaluator.EvaluateModel {
	var servicesNode []core_domain.CodeDataStruct = nil
	var evaluation Evaluation
	var result = evaluator.NewEvaluateModel()

	var nodeMap = make(map[string]core_domain.CodeDataStruct)

	for _, node := range classNodes {
		nodeMap[node.NodeName] = node

		if node.IsUtilClass() {
			result.Summary.UtilsCount++

			evaluation = Evaluation{evaluator.Util{}}
			evaluation.Evaluate(&result, node)
		}

		if node.IsServiceClass() {
			servicesNode = append(servicesNode, node)
		} else {
			evaluation = Evaluation{evaluator.Empty{}}
		}
	}

	SummaryMethodIdentifier(identifiers, &result)

	evaluation = Evaluation{evaluator.Service{}}
	evaluation.EvaluateList(&result, servicesNode, nodeMap, identifiers)

	nullableEva := Evaluation{evaluator.NullPointException{}}
	nullableEva.EvaluateList(&result, servicesNode, nodeMap, identifiers)

	return result
}

func SummaryMethodIdentifier(identifiers []core_domain.CodeDataStruct, result *evaluator.EvaluateModel) {
	var methodLengthArray []float64
	var methodCountArray []float64
	for _, ident := range identifiers {
		result.Summary.ClassCount++

		methodCountArray = append(methodCountArray, float64(len(ident.Functions)))

		for _, method := range ident.Functions {
			result.Summary.MethodCount++

			if method.IsStatic() {
				result.Summary.StaticMethodCount++
			}

			if method.IsGetterSetter() {
				result.Summary.NormalMethodCount++
				methodLength := method.Position.StopLine - method.Position.StartLine + 1
				result.Summary.TotalMethodLength = result.Summary.TotalMethodLength + methodLength

				methodLengthArray = append(methodLengthArray, float64(methodLength))
			}
		}
	}

	result.Summary.MethodLengthStdDeviation = stat.StdDev(methodLengthArray, nil)
	result.Summary.MethodNumStdDeviation = stat.StdDev(methodCountArray, nil)
}
