package core_domain

import (
	"github.com/phodal/coca/pkg/infrastructure/constants"
	"strings"
)

type CodeCall struct {
	Package    string
	Type       string
	Class      string
	MethodName string
	Parameters []CodeProperty
	Position   CodePosition
}

func NewCodeMethodCall() CodeCall {
	return CodeCall{}
}

func (c *CodeCall) BuildFullMethodName() string {
	isConstructor := c.MethodName == ""
	if isConstructor {
		return c.Package + "." + c.Class
	}
	return c.Package + "." + c.Class + "." + c.MethodName
}

func (c *CodeCall) IsSystemOutput() bool {
	return c.Class == "System.out" && (c.MethodName == "println" || c.MethodName == "printf" || c.MethodName == "print")
}

func (c *CodeCall) IsThreadSleep() bool {
	return c.MethodName == "sleep" && c.Class == "Thread"
}

func (c *CodeCall) HasAssertion() bool {
	methodName := strings.ToLower(c.MethodName)
	for _, assertion := range constants.ASSERTION_LIST {
		if strings.HasPrefix(methodName, assertion) {
			return true
		}
	}

	return false
}
