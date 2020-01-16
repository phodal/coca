package core_domain

import (
	"github.com/phodal/coca/pkg/infrastructure/string_helper"
	"strings"
)

type CodeFunction struct {
	Name            string
	ReturnTypes     []CodeProperty
	Parameters      []CodeProperty
	MethodCalls     []CodeCall
	Override        bool
	Annotations     interface{}
	Modifiers       []string
	InnerStructures []CodeDataStruct // InnerClass
	InnerFunctions  []CodeFunction
	Extension       interface{}
	Position        CodePosition
}

func (c *CodeFunction) BuildSingleReturnType(typeType string) *CodeProperty {
	return &CodeProperty{
		TypeType: typeType,
	}
}

type JMethod struct {
	Name              string
	Type              string
	Parameters        []CodeProperty
	MethodCalls       []CodeCall
	Override          bool
	Annotations       []CodeAnnotation
	IsConstructor     bool
	IsReturnNull      bool
	Modifiers         []string
	Creators          []JClassNode
	Position          CodePosition
}

func NewJMethod() JMethod {
	return JMethod{}
}

func (m *JMethod) IsJavaLangReturnType() bool {
	return m.Type == "String" || m.Type == "int" || m.Type == "float" || m.Type == "void" || m.Type == "char" || m.Type == "double"
}

func (m *JMethod) IsStatic() bool {
	return string_helper.StringArrayContains(m.Modifiers, "static")
}

func (m *JMethod) IsGetterSetter() bool {
	return strings.HasPrefix(m.Name, "set") || strings.HasPrefix(m.Name, "get")
}

func (m *JMethod) BuildFullMethodName(node JClassNode) string {
	return node.Package + "." + node.Class + "." + m.Name
}

func (m *JMethod) GetAllCallString() []string {
	var calls []string
	for _, call := range m.MethodCalls {
		if call.Class != "" {
			calls = append(calls, call.BuildFullMethodName())
		}
	}
	return calls
}

func (m *JMethod) IsJunitTest() bool {
	var isTest = false
	for _, annotation := range m.Annotations {
		if annotation.IsIgnoreOrTest() {
			isTest = true
		}
	}
	return isTest
}
