package models

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
