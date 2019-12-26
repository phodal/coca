package evaluator

import (
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"github.com/phodal/coca/core/support/apriori"
	"strings"
)

type Service struct {
}

var serviceNodeMap map[string]models.JClassNode
var returnTypeMap map[string][]string
var longParameterList []models.JMethod

func (s Service) EvaluateList(evaluateModel *EvaluateModel, nodes []models.JClassNode, nodeMap map[string]models.JClassNode, identifiers []models.JIdentifier) {
	serviceNodeMap = nodeMap
	longParameterList = nil
	returnTypeMap = make(map[string][]string)

	for _, node := range nodes {
		s.Evaluate(evaluateModel, node)
	}

	evaluateModel.ServiceSummary.ReturnTypeMap = returnTypeMap
	findRelatedMethodParameter(evaluateModel, longParameterList)
}

func findRelatedMethodParameter(model *EvaluateModel, list []models.JMethod) {
	var dataset [][]string
	for _, method := range list {
		var methodlist []string
		for _, param := range method.Parameters {
			methodlist = append(methodlist, param.Type)
		}
		dataset = append(dataset, methodlist)
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

func (s Service) Evaluate(result *EvaluateModel, node models.JClassNode) {
	var methodNameArray [][]string
	for _, method := range node.Methods {
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
	if s.enableSameBehavior() {

	}

	if s.enableAbstractParameters() {
		for _, method := range node.Methods {
			if len(method.Parameters) >= 4 {
				longParameterList = append(longParameterList, method)
			}
		}
	}

	if s.enableSameReturnType() {
		for _, method := range node.Methods {
			if !s.isJavaType(method) {
				methodType := method.Type

				if _, ok := serviceNodeMap[methodType]; ok {
					fullMethodName := node.Package + "." + node.Class + "." + method.Name
					returnTypeMap[methodType] = append(returnTypeMap[methodType], fullMethodName)
				}
			}
		}
	}
}

func (s Service) isJavaType(method models.JMethod) bool {
	return method.Type == "String" || method.Type == "int"
}

func (s Service) buildLifecycle(methodNameArray [][]string) map[string][]string {
	var hadLifecycle = make(map[string][]string)
	var nameMap = make(map[string][]string)
	for _, nameArray := range methodNameArray {
		if len(nameArray) < 1 {
			continue
		}

		firstWord := nameArray[0]
		if !(support.IsTechStopWords(firstWord)) {
			nameMap[firstWord] = append(nameMap[firstWord], strings.Join(nameArray, ""))
		}
		if len(nameMap[firstWord]) > 1 {
			hadLifecycle[firstWord] = nameMap[firstWord]
		}
	}

	return hadLifecycle
}

func (s Service) enableLifecycle() bool {
	return true
}

func (s Service) enableSameBehavior() bool {
	return true
}

func (s Service) enableAbstractParameters() bool {
	return true
}

func (s Service) enableSameReturnType() bool {
	return true
}
