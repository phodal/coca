package models

var methods []JMethod

type JIdentifier struct {
	Package     string
	ClassName   string
	Type        string
	ExtendsName string
	Extends     []JIdentifier
	Methods     []JMethod
}

func NewJIdentifier() *JIdentifier {
	identifier := &JIdentifier{"", "", "", "", nil, nil}
	methods = nil
	return identifier
}

func (identifier *JIdentifier) AddMethod(method JMethod) {
	methods = append(methods, method)
}

func (identifier *JIdentifier) GetMethods() []JMethod {
	return methods
}
