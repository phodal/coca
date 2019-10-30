package models

type JClassNode struct {
	Package     string
	Class       string
	Type        string
	MethodCalls []JMethodCall
}
