package models

type JFullClassNode struct {
	Package     string
	Class       string
	Type        string
	Path        string
	Methods     []JFullMethod
	MethodCalls []JFullMethodCall
	ClassBS     ClassBadSmellInfo
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
	MethodBody        string
	Parameters        []JFullParameter
	MethodBs          MethodBadSmellInfo
}

type MethodBadSmellInfo struct {
	SwitchSize int
	IfSize     int
}

type ClassBadSmellInfo struct {
	OverrideSize  int
	PublicVarSize int
}

type JFullParameter struct {
	Name string
	Type string
}


func NewJFullClassNode() *JFullClassNode {
	info := &ClassBadSmellInfo{0, 0};
	return &JFullClassNode{"", "", "", "", nil, nil, *info}
}
