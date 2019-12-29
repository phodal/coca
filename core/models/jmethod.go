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
	Annotations       []Annotation
	IsConstructor     bool
	IsReturnNull      bool
	Modifiers         []string
	Creators          []JClassNode
}

type Annotation struct {
	QualifiedName string
	ValuePairs    []AnnotationKeyValue
}

type AnnotationKeyValue struct {
	Key   string
	Value string
}

func NewAnnotation() Annotation {
	return *&Annotation{
		QualifiedName: "",
		ValuePairs:    nil,
	}
}

// TODO support annnotation
func NewJMethod() JMethod {
	return *&JMethod{
		Name:              "",
		Type:              "",
		Annotations:       nil,
		StartLine:         0,
		StartLinePosition: 0,
		StopLine:          0,
		StopLinePosition:  0,
		Parameters:        nil,
		MethodCalls:       nil,
		IsConstructor:     false,
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
