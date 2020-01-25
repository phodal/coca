package bs_domain

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
)

type BSDataStruct struct {
	core_domain.CodeDataStruct

	Functions    []BSFunction
	DataStructBS ClassBadSmellInfo
}

type BSFunction struct {
	core_domain.CodeFunction

	FunctionBody string
	FunctionBS   FunctionBSInfo
}

type FunctionBSInfo struct {
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

func NewMethodBadSmellInfo() FunctionBSInfo {
	return FunctionBSInfo{
		IfSize:     0,
		SwitchSize: 0,
		IfInfo:     nil,
	}
}

type ClassBadSmellInfo struct {
	OverrideSize  int
	PublicVarSize int
}

func NewJFullClassNode() BSDataStruct {
	info := ClassBadSmellInfo{0, 0}
	return BSDataStruct{
		DataStructBS: info,
	}
}

func (b *BSDataStruct) HaveCallParent() bool {
	hasCallParentMethod := false
	for _, methodCall := range b.FunctionCalls {
		if methodCall.NodeName == b.Extend {
			hasCallParentMethod = true
		}
	}
	return hasCallParentMethod
}
