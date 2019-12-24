package evaluator

import (
	"fmt"
	"github.com/phodal/coca/core/models"
	"strings"
)

type Service struct {
}

var nodeMap map[string]models.JClassNode
var returnTypeMap map[string][]string

func (s Service) EvaluateList(nodes []models.JClassNode, classNodeMap map[string]models.JClassNode) {
	nodeMap = classNodeMap
	returnTypeMap = make(map[string][]string)
	for _, node := range nodes {
		s.Evaluate(node)
	}

	fmt.Println(returnTypeMap)
}

func (s Service) Evaluate(node models.JClassNode) {
	var methodNameArray [][]string
	for _, method := range node.Methods {
		methodNameArray = append(methodNameArray, SplitCamelcase(method.Name))
	}

	lifecycleMap := s.buildLifecycle(methodNameArray)
	hasLifecycle := len(lifecycleMap) > 0
	if hasLifecycle {
		for key, value := range lifecycleMap {
			fmt.Println(key, value)
		}
	}

	if s.hasSameBehavior() {

	}

	if s.hasAbstractParameters() {
		// parameters
	}

	if s.hasSameReturnType() {
		for _, method := range node.Methods {
			fmt.Println(method)

			if s.isJavaType(method) {
				fullPackage := s.getReturnTypeFullPackage(method)
				if _, ok := nodeMap[fullPackage]; ok {
					fullMethodName := node.Package + "." + node.Class + "." + method.Name
					returnTypeMap[fullPackage] = append(returnTypeMap[fullPackage], fullMethodName)
				}
			}
		}
	}
}

func (s Service) isJavaType(method models.JMethod) bool {
	return method.Type == "String" || method.Type == "int"
}

func (s Service) getReturnTypeFullPackage(method models.JMethod) string {
	for _, call := range method.MethodCalls {
		if call.Class == method.Type {
			return call.Class
		}
	}
	return ""
}

func (s Service) buildLifecycle(methodNameArray [][]string) map[string][]string {
	var hadLifecycle = make(map[string][]string)
	var nameMap = make(map[string][]string)
	for _, nameArray := range methodNameArray {
		firstWord := nameArray[0]
		if !(firstWord == "set" || firstWord == "get") {
			nameMap[firstWord] = append(nameMap[firstWord], strings.Join(nameArray, ""))
		}
		if len(nameMap[firstWord]) > 1 {
			hadLifecycle[firstWord] = nameMap[firstWord]
		}
	}

	return hadLifecycle
}

func (s Service) hasSameBehavior() bool {
	return false
}

func (s Service) hasAbstractParameters() bool {
	return false
}

func (s Service) hasSameReturnType() bool {
	return true
}
