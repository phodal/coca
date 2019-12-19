package models

var methods []JMethod

type JIdentifier struct {
	Package     string
	ClassName   string
	Type        string
	ExtendsName string
	Extends     []string
	Implements     []string
	Methods     []JMethod
}

func NewJIdentifier() *JIdentifier {
	identifier := &JIdentifier{"", "", "", "", nil,  nil, nil}
	methods = nil
	return identifier
}

func (identifier *JIdentifier) AddMethod(method JMethod) {
	methods = append(methods, method)
}

func (identifier *JIdentifier) GetMethods() []JMethod {
	return methods
}
