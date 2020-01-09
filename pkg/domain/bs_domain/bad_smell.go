package bs_domain

import (
	"sort"
	"strings"
)

type BsJClass struct {
	Package     string
	Class       string
	Type        string
	Path        string
	Extends     string
	Implements  []string
	Methods     []BsJMethod
	MethodCalls []BsJMethodCall
	ClassBS     ClassBadSmellInfo
}

type BsJMethodCall struct {
	Package           string
	Type              string
	Class             string
	MethodName        string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
}

type BsJMethod struct {
	Name              string
	Type              string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
	MethodBody        string
	Modifier          string
	Parameters        []JFullParameter
	MethodBs          MethodBadSmellInfo
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
	return *&IfParInfo{
		StartLine: 0,
		EndLine:   0,
	}
}

func NewMethodBadSmellInfo() MethodBadSmellInfo {
	return *&MethodBadSmellInfo{
		IfSize:     0,
		SwitchSize: 0,
		IfInfo:     nil,
	}
}

type ClassBadSmellInfo struct {
	OverrideSize  int
	PublicVarSize int
}

type JFullParameter struct {
	Name string
	Type string
}

func NewJFullClassNode() BsJClass {
	info := &ClassBadSmellInfo{0, 0}
	return *&BsJClass{
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

type BadSmellModel struct {
	File        string `json:"EntityName,omitempty"`
	Line        string `json:"Line,omitempty"`
	Bs          string `json:"BS,omitempty"`
	Description string `json:"Description,omitempty"`
	Size        int    `size:"Description,omitempty"`
}

func (b *BsJMethod) IsGetterSetter() bool {
	return strings.HasPrefix(b.Name, "set") || strings.HasPrefix(b.Name, "get")
}

func (b *BsJClass) HaveCallParent() bool {
	hasCallParentMethod := false
	for _, methodCall := range b.MethodCalls {
		if methodCall.Class == b.Extends {
			hasCallParentMethod = true
		}
	}
	return hasCallParentMethod
}

func (b *BsJClass) ClassFullName() string {
	return b.Package + "." + b.Class
}

func (c *BsJMethodCall) ClassFullName() string {
	return c.Package + "." + c.Class
}

func SortSmellByType(models []BadSmellModel, filterFunc func(key string) bool) map[string][]BadSmellModel {
	sortSmells := make(map[string][]BadSmellModel)
	for _, model := range models {
		sortSmells[model.Bs] = append(sortSmells[model.Bs], model)
	}

	for key, smells := range sortSmells {
		if filterFunc(key) {
			sort.Slice(smells, func(i, j int) bool {
				return smells[i].Size > (smells[j].Size)
			})

			sortSmells[key] = smells
		}
	}

	return sortSmells
}

func FilterBadSmellList(models []BadSmellModel, ignoreRules map[string]bool) []BadSmellModel {
	var results []BadSmellModel
	for _, model := range models {
		if !ignoreRules[model.Bs] {
			results = append(results, model)
		}
	}
	return results
}
