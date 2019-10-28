package models

import "fmt"

type JMethodCall struct {
	Pkg         string
	Dlz         string
	MethodName  string
	//methodCalls map[string]string
	//tableOps    map[string]string
}

func (call *JMethodCall) AddMethodCall (targetType string, method string) {
	fmt.Println(targetType, "->", method)
}

