package models

type JFullClassNode struct {
	Package     string
	Class       string
	Type        string
	Path        string
	Methods     []JFullMethod
	MethodCalls []JFullMethodCall
}

type JFullMethodCall struct {
	Package           string
	Type              string
	Class             string
	MethodName        string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
}

type JFullMethod struct {
	Name              string
	Type              string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
	Parameters        []JFullParameter
}

type JFullParameter struct {
	Name string
	Type string
}

func NewJFullClassNode() *JFullClassNode {
	return &JFullClassNode{"", "", "",  "", nil, nil}
}
