package jdomain

import (
	"github.com/phodal/coca/pkg/infrastructure/constants"
	"strings"
)

type JMethodCall struct {
	Package           string
	Type              string
	Class             string
	MethodName        string
	Parameters        []string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
}

func NewJMethodCall() JMethodCall {
	return JMethodCall{
		Package:           "",
		Type:              "",
		Class:             "",
		MethodName:        "",
		Parameters:        nil,
		StartLine:         0,
		StartLinePosition: 0,
		StopLine:          0,
		StopLinePosition:  0,
	}
}

func (c *JMethodCall) BuildFullMethodName() string {
	isConstructor := c.MethodName == ""
	if isConstructor {
		return c.Package + "." + c.Class
	}
	return c.Package + "." + c.Class + "." + c.MethodName
}

func (c *JMethodCall) IsSystemOutput() bool {
	return c.Class == "System.out" && (c.MethodName == "println" || c.MethodName == "printf" || c.MethodName == "print")
}

func (c *JMethodCall) IsThreadSleep() bool {
	return c.MethodName == "sleep" && c.Class == "Thread"
}

func (c *JMethodCall) HasAssertion() bool {
	methodName := strings.ToLower(c.MethodName)
	for _, assertion := range constants.ASSERTION_LIST {
		if strings.HasPrefix(methodName, assertion) {
			return true
		}
	}

	return false
}

