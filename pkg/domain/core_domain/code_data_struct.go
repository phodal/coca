package core_domain

import "strings"

type CodeDataStruct struct {
	NodeName        string
	Type            string
	Package         string
	FilePath        string
	Fields          []CodeField
	Extend          string
	MultipleExtend  []string // for C++
	Implements      []string
	Functions       []CodeFunction
	InnerStructures []CodeDataStruct
	Annotations     []CodeAnnotation
	FunctionCalls   []CodeCall     // for field call
	InOutProperties []CodeProperty //for golang interface
	// Deprecated should get from code file
	Imports   []CodeImport
	Extension interface{}
}

func NewDataStruct() *CodeDataStruct {
	return &CodeDataStruct{}
}

func (j *CodeDataStruct) IsUtilClass() bool {
	return strings.Contains(strings.ToLower(j.NodeName), "util") || strings.Contains(strings.ToLower(j.NodeName), "utils")
}

func (j *CodeDataStruct) IsServiceClass() bool {
	return strings.Contains(strings.ToLower(j.NodeName), "service")
}

func (j *CodeDataStruct) SetMethodFromMap(methodMap map[string]CodeFunction) {
	var methodsArray []CodeFunction
	for _, value := range methodMap {
		methodsArray = append(methodsArray, value)
	}

	j.Functions = methodsArray
}

func (j *CodeDataStruct) BuildStringMethodMap(projectMethods map[string]string) {
	for _, method := range j.Functions {
		projectMethods[method.BuildFullMethodName(*j)] = method.BuildFullMethodName(*j)
	}
}

func (j *CodeDataStruct) IsNotEmpty() bool {
	return len(j.Functions) > 0 || len(j.FunctionCalls) > 0
}

func BuildCallMethodMap(deps []CodeDataStruct) map[string]CodeFunction {
	var callMethodMap = make(map[string]CodeFunction)
	for _, clz := range deps {
		for _, method := range clz.Functions {
			callMethodMap[method.BuildFullMethodName(clz)] = method
		}
	}
	return callMethodMap
}
