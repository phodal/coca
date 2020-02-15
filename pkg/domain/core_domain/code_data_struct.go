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
	//Parameters      []CodeProperty
	Extension       interface{}
	Imports         []CodeImport // deprecated: should get from code file
}

func NewDataStruct() *CodeDataStruct {
	return &CodeDataStruct{}
}

func (d *CodeDataStruct) IsUtilClass() bool {
	return strings.Contains(strings.ToLower(d.NodeName), "util") || strings.Contains(strings.ToLower(d.NodeName), "utils")
}

func (d *CodeDataStruct) IsServiceClass() bool {
	return strings.Contains(strings.ToLower(d.NodeName), "service")
}

func (d *CodeDataStruct) SetMethodFromMap(methodMap map[string]CodeFunction) {
	var methodsArray []CodeFunction
	for _, value := range methodMap {
		methodsArray = append(methodsArray, value)
	}

	d.Functions = methodsArray
}

func (d *CodeDataStruct) BuildStringMethodMap(projectMethods map[string]string) {
	for _, method := range d.Functions {
		projectMethods[method.BuildFullMethodName(*d)] = method.BuildFullMethodName(*d)
	}
}

func (d *CodeDataStruct) IsNotEmpty() bool {
	return len(d.Functions) > 0 || len(d.FunctionCalls) > 0
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

func (d *CodeDataStruct) GetClassFullName() string {
	return d.Package + "." + d.NodeName
}

func BuildIdentifierMap(identifiers []CodeDataStruct) map[string]CodeDataStruct {
	var identifiersMap = make(map[string]CodeDataStruct)

	for _, ident := range identifiers {
		identifiersMap[ident.Package+"."+ident.NodeName] = ident
	}
	return identifiersMap
}

func BuildDIMap(identifiers []CodeDataStruct, identifierMap map[string]CodeDataStruct) map[string]string {
	var diMap = make(map[string]string)
	for _, clz := range identifiers {
		if len(clz.Annotations) > 0 {
			for _, annotation := range clz.Annotations {
				if (annotation.IsComponentOrRepository()) && len(clz.Implements) > 0 {
					superClz := identifierMap[clz.Implements[0]]
					diMap[superClz.GetClassFullName()] = superClz.GetClassFullName()
				}
			}
		}
	}

	return diMap
}
