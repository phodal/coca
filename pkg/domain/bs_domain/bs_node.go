package bs_domain

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
	"strings"
)

type BsJClass struct {
	Package       string
	Type          string
	NodeName      string
	FilePath      string
	Extend        string
	Implements    []string
	Functions     []BsJMethod
	FunctionCalls []core_domain.CodeCall
	ClassBS       ClassBadSmellInfo
}

type BsJMethod struct {
	Name       string
	ReturnType string
	Modifier   []string
	Parameters []core_domain.CodeProperty
	Position   core_domain.CodePosition

	MethodBody string
	MethodBs   MethodBadSmellInfo
}

type MethodBadSmellInfo struct {
	IfSize     int
	SwitchSize int
	IfInfo     []IfParInfo
}

type IfParInfo struct {
	StartLine int
	EndLine   int
}

func NewIfPairInfo() IfParInfo {
	return IfParInfo{
		StartLine: 0,
		EndLine:   0,
	}
}

func NewMethodBadSmellInfo() MethodBadSmellInfo {
	return MethodBadSmellInfo{
		IfSize:     0,
		SwitchSize: 0,
		IfInfo:     nil,
	}
}

type ClassBadSmellInfo struct {
	OverrideSize  int
	PublicVarSize int
}

func NewJFullClassNode() BsJClass {
	info := &ClassBadSmellInfo{0, 0}
	return BsJClass{
		"",
		"",
		"",
		"",
		"",
		nil,
		nil,
		nil,
		*info}
}

func (b *BsJMethod) IsGetterSetter() bool {
	return strings.HasPrefix(b.Name, "set") || strings.HasPrefix(b.Name, "get")
}

func (b *BsJClass) HaveCallParent() bool {
	hasCallParentMethod := false
	for _, methodCall := range b.FunctionCalls {
		if methodCall.NodeName == b.Extend {
			hasCallParentMethod = true
		}
	}
	return hasCallParentMethod
}

func (b *BsJClass) ClassFullName() string {
	return b.Package + "." + b.NodeName
}
