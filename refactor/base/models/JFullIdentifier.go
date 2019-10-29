package models

type JFullMethod struct {
	Name              string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
}

var methods []JFullMethod

type JFullIdentifier struct {
	Pkg  string
	Name string
	Type string
}

func NewJFullIdentifier() *JFullIdentifier {
	identifier := &JFullIdentifier{"", "", ""}
	methods = nil
	return identifier
}

func (identifier *JFullIdentifier) AddMethod(method JFullMethod) {
	methods = append(methods, method)
}

func (identifier *JFullIdentifier) GetMethods() []JFullMethod {
	return methods
}
