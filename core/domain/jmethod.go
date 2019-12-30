package domain

import (
	"github.com/phodal/coca/core/infrastructure"
	"strings"
)

type JMethod struct {
	Name              string
	Type              string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
	Parameters        []JParameter
	MethodCalls       []JMethodCall
	Override          bool
	Annotations       []Annotation
	IsConstructor     bool
	IsReturnNull      bool
	Modifiers         []string
	Creators          []JClassNode
}

func NewJMethod() JMethod {
	return *&JMethod{
		Name:              "",
		Type:              "",
		Annotations:       nil,
		StartLine:         0,
		StartLinePosition: 0,
		StopLine:          0,
		StopLinePosition:  0,
		Parameters:        nil,
		MethodCalls:       nil,
		IsConstructor:     false,
	}
}

type JMethodInfo struct {
	Name       string
	Type       string
	Parameters []JParameter
	Length     string
}

func (m *JMethod) IsJavaLangReturnType() bool {
	return m.Type == "String" || m.Type == "int" || m.Type == "float" || m.Type == "void" || m.Type == "char" || m.Type == "double"
}

func (m *JMethod) IsStatic() bool {
	return infrastructure.StringArrayContains(m.Modifiers, "static")
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
			calls = append(calls, call.BuilFullMethodName())
		}
	}
	return calls
}
