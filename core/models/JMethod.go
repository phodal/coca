package models

type JMethod struct {
	Name              string
	Type              string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
	Parameters        []JParameter
	MethodCalls       []JMethodCall
	Override          bool
	Annotation        []string
}

func NewJMethod() JMethod {
	return *&JMethod{
		Name:              "",
		Type:              "",
		Annotation:        nil,
		StartLine:         0,
		StartLinePosition: 0,
		StopLine:          0,
		StopLinePosition:  0,
		Parameters:        nil,
		MethodCalls:       nil,
	}
}

type JParameter struct {
	Name string
	Type string
}

type JMethodInfo struct {
	Name       string
	Type       string
	Parameters []JParameter
	Length     string
}
