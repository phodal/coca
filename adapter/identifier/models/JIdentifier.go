package models

var methods []JMethod

type JIdentifier struct {
	Pkg  string
	Name string
	Type string
}

func (identifier *JIdentifier) AddMethod(method JMethod) {
	methods = append(methods, method)
}
