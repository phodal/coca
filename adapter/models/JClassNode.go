package models

type JClassNode struct {
	Package     string
	Class       string
	Type        string
	Methods     []JMethod
	MethodCalls []JMethodCall
}

type JsonIdentifier struct {
	Package string
	Name    string
	Type    string
	Methods []JMethod
}

func NewClassNode() *JClassNode {
	return &JClassNode{"", "", "", nil, nil}
}

func NewJsonIdentifier() *JsonIdentifier {
	return &JsonIdentifier{"", "", "", nil}
}
