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

type JsonIdentifier struct {
	Package string
	Name    string
	Type    string
	Methods []JMethod
}

func NewClassNode() *JClassNode {
	return &JClassNode{"", "", "", "", nil, nil, nil}
}

func NewJsonIdentifier() *JsonIdentifier {
	return &JsonIdentifier{"", "", "", nil}
}
