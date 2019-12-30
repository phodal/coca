package domain

import "strings"

type JClassNode struct {
	Package     string
	Class       string
	Type        string
	Path        string
	Fields      []JAppField
	Methods     []JMethod
	MethodCalls []JMethodCall
	Extend      string
	Implements  []string
	Annotations []Annotation
}

type JAppField struct {
	Type  string
	Value string
}

func NewClassNode() *JClassNode {
	return &JClassNode{"", "", "", "", nil, nil, nil, "", nil, nil}
}

func (j *JClassNode) IsUtilClass() bool {
	return strings.Contains(strings.ToLower(j.Class), "util") || strings.Contains(strings.ToLower(j.Class), "utils")
}

func (j *JClassNode) IsServiceClass() bool {
	return strings.Contains(strings.ToLower(j.Class), "service")
}

func (j *JClassNode) SetMethodFromMap(methodMap map[string]JMethod) {
	var methodsArray []JMethod
	for _, value := range methodMap {
		methodsArray = append(methodsArray, value)
	}

	j.Methods = methodsArray
}

func (j *JClassNode) BuildStringMethodMap(projectMethods map[string]string) {
	for _, method := range j.Methods {
		projectMethods[method.GetFullMethodName(*j)] = method.GetFullMethodName(*j)
	}
}
