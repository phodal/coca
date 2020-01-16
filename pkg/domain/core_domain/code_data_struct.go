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

type JIdentifier struct {
	NodeName    string
	Package     string
	Type        string
	Extend      string
	Implements  []string
	Functions   []CodeFunction
	Annotations []CodeAnnotation
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

func NewJIdentifier() *JIdentifier {
	return &JIdentifier{}
}

func (identifier *JIdentifier) GetClassFullName() string {
	return identifier.Package + "." + identifier.NodeName
}

func BuildIdentifierMap(identifiers []JIdentifier) map[string]JIdentifier {
	var identifiersMap = make(map[string]JIdentifier)

	for _, ident := range identifiers {
		identifiersMap[ident.Package+"."+ident.NodeName] = ident
	}
	return identifiersMap
}

func BuildDIMap(identifiers []JIdentifier, identifierMap map[string]JIdentifier) map[string]string {
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
