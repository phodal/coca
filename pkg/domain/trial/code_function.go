package trial

type CodeFunction struct {
	Name           string
	ReturnTypes    []CodeProperty
	Parameters     []CodeProperty
	MethodCalls    []CodeCall
	Override       bool
	Annotations    []CodeAnnotation
	Modifiers      []string
	Creators       []CodeDataStruct
	InnerFunctions []CodeFunction
}

type CodeAnnotation struct {
}
