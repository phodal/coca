package core_domain

import (
	"github.com/phodal/coca/pkg/infrastructure/string_helper"
	"strings"
)

type CodeFunction struct {
	Name            string
	ReturnType      string
	MultipleReturns []CodeProperty
	Parameters      []CodeProperty
	FunctionCalls   []CodeCall
	Override        bool
	Annotations     []CodeAnnotation

	IsConstructor bool // todo: move to extension
	IsReturnNull  bool // todo: move to extension

	Modifiers       []string
	InnerStructures []CodeDataStruct
	InnerFunctions  []CodeFunction
	Extension       interface{}
	Position        CodePosition
}

func (m *CodeFunction) BuildSingleReturnType(typeType string) *CodeProperty {
	return &CodeProperty{
		TypeType: typeType,
	}
}

func NewJMethod() CodeFunction {
	return CodeFunction{}
}

func (m *CodeFunction) IsJavaLangReturnType() bool {
	return m.ReturnType == "String" || m.ReturnType == "int" || m.ReturnType == "float" || m.ReturnType == "void" || m.ReturnType == "char" || m.ReturnType == "double"
}

func (m *CodeFunction) IsStatic() bool {
	return string_helper.StringArrayContains(m.Modifiers, "static")
}

func (m *CodeFunction) IsGetterSetter() bool {
	return strings.HasPrefix(m.Name, "set") || strings.HasPrefix(m.Name, "get")
}

func (m *CodeFunction) BuildFullMethodName(node CodeDataStruct) string {
	return node.Package + "." + node.NodeName + "." + m.Name
}

func (m *CodeFunction) GetAllCallString() []string {
	var calls []string
	for _, call := range m.FunctionCalls {
		if call.NodeName != "" {
			calls = append(calls, call.BuildFullMethodName())
		}
	}
	return calls
}

func (m *CodeFunction) IsJunitTest() bool {
	var isTest = false
	for _, annotation := range m.Annotations {
		if annotation.IsIgnoreOrTest() {
			isTest = true
		}
	}
	return isTest
}
