package models

var methods []JMethod

type JIdentifier struct {
	Pkg  string
	Name string
	Type string
}

func NewJIdentifier() *JIdentifier {
	identifier := &JIdentifier{"", "", ""}
	methods = nil
	return identifier
}

func (identifier *JIdentifier) AddMethod(method JMethod) {
	methods = append(methods, method)
}

func (identifier *JIdentifier) GetMethods() []JMethod {
	return methods
}
