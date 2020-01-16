package jdomain

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
	"strings"
)

type JClassNode struct {
	Package     string
	Class       string
	Type        string
	FilePath    string
	Fields      []core_domain.CodeField
	Methods     []core_domain.JMethod
	MethodCalls []core_domain.CodeCall
	Extend      string
	Implements  []string
	Annotations []core_domain.CodeAnnotation
	InnerClass  []JClassNode
	Imports     []core_domain.CodeImport
}

func NewClassNode() *JClassNode {
	return &JClassNode{}
}

func (j *JClassNode) IsUtilClass() bool {
	return strings.Contains(strings.ToLower(j.Class), "util") || strings.Contains(strings.ToLower(j.Class), "utils")
}

func (j *JClassNode) IsServiceClass() bool {
	return strings.Contains(strings.ToLower(j.Class), "service")
}

func (j *JClassNode) SetMethodFromMap(methodMap map[string]core_domain.JMethod) {
	var methodsArray []core_domain.JMethod
	for _, value := range methodMap {
		methodsArray = append(methodsArray, value)
	}

	j.Methods = methodsArray
}

func (j *JClassNode) BuildStringMethodMap(projectMethods map[string]string) {
	for _, method := range j.Methods {
		projectMethods[method.BuildFullMethodName(*j)] = method.BuildFullMethodName(*j)
	}
}

func (j *JClassNode) IsNotEmpty() bool {
	return len(j.Methods) > 0 || len(j.MethodCalls) > 0
}

func BuildCallMethodMap(deps []JClassNode) map[string]core_domain.JMethod {
	var callMethodMap = make(map[string]core_domain.JMethod)
	for _, clz := range deps {
		for _, method := range clz.Methods {
			callMethodMap[method.BuildFullMethodName(clz)] = method
		}
	}
	return callMethodMap
}
