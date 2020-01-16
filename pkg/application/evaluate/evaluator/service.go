package evaluator

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/apriori"
	"github.com/phodal/coca/pkg/infrastructure/constants"
	"strings"
)

type Service struct {
}

var serviceNodeMap map[string]core_domain.CodeDataStruct
var returnTypeMap map[string][]string
var longParameterList []core_domain.CodeFunction

func (s Service) EvaluateList(evaluateModel *EvaluateModel, nodes []core_domain.CodeDataStruct, nodeMap map[string]core_domain.CodeDataStruct, identifiers []core_domain.CodeDataStruct) {
	serviceNodeMap = nodeMap
	longParameterList = nil
	returnTypeMap = make(map[string][]string)

	for _, node := range nodes {
		s.Evaluate(evaluateModel, node)
	}

	evaluateModel.ServiceSummary.ReturnTypeMap = returnTypeMap
	findRelatedMethodParameters(evaluateModel, longParameterList)
}

func findRelatedMethodParameters(model *EvaluateModel, list []core_domain.CodeFunction) {
	var dataset [][]string
	for _, method := range list {
		var paramTypeList []string
		for _, param := range method.Parameters {
			paramTypeList = append(paramTypeList, param.TypeValue)
		}
		dataset = append(dataset, paramTypeList)
	}

	var newOptions = apriori.NewOptions(0.8, 0.8, 0, 0)
	apriori := apriori.NewApriori(dataset)
	result := apriori.Calculate(newOptions)

	for _, res := range result {
		items := res.GetSupportRecord().GetItems()
		if len(items) >= 4 {
			model.ServiceSummary.RelatedMethod = items
		}
	}
}

func (s Service) Evaluate(result *EvaluateModel, node core_domain.CodeDataStruct) {
	var methodNameArray [][]string
	for _, method := range node.Functions {
		methodNameArray = append(methodNameArray, SplitCamelcase(method.Name))
	}

	if s.enableLifecycle() {
		lifecycleMap := s.buildLifecycle(methodNameArray)
		hasLifecycle := len(lifecycleMap) > 0
		if hasLifecycle {
			result.ServiceSummary.LifecycleMap = lifecycleMap
		}
	}

	// TODO: support for same end words
	//if s.enableSameBehavior() {
	//
	//}

	if s.enableAbstractParameters() {
		for _, method := range node.Functions {
			PARAMETERR_LENGTH_LIMIT := 4
			if len(method.Parameters) >= PARAMETERR_LENGTH_LIMIT {
				longParameterList = append(longParameterList, method)
			}
		}
	}

	if s.enableSameReturnType() {
		for _, method := range node.Functions {
			if !method.IsJavaLangReturnType() {
				methodType := method.ReturnType

				if _, ok := serviceNodeMap[methodType]; ok {
					returnTypeMap[methodType] = append(returnTypeMap[methodType], method.BuildFullMethodName(node))
				}
			}
		}
	}
}

func (s Service) buildLifecycle(methodNameArray [][]string) map[string][]string {
	var hadLifecycle = make(map[string][]string)
	var nameMap = make(map[string][]string)
	for _, nameArray := range methodNameArray {
		if len(nameArray) < 1 {
			continue
		}

		firstWord := nameArray[0]
		if !(IsTechStopWords(firstWord)) {
			nameMap[firstWord] = append(nameMap[firstWord], strings.Join(nameArray, ""))
		}
		if len(nameMap[firstWord]) > 1 {
			hadLifecycle[firstWord] = nameMap[firstWord]
		}
	}

	return hadLifecycle
}

func IsTechStopWords(firstWord string) bool {
	for _, word := range constants.TechStopWords {
		if word == firstWord {
			return true;
		}
	}

	return false;
}

func (s Service) enableLifecycle() bool {
	return true
}

func (s Service) enableAbstractParameters() bool {
	return true
}

func (s Service) enableSameReturnType() bool {
	return true
}
