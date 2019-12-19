package models

type JClassNode struct {
	Package     string
	Class       string
	Type        string
	Path        string
	Fields      []JAppField
	Methods     []JMethod
	MethodCalls []JMethodCall
}

type JAppField struct {
	Type  string
	Value string
}

func NewClassNode() *JClassNode {
	return &JClassNode{"", "", "", "", nil, nil, nil}
}

