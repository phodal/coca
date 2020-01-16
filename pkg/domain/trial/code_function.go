package trial

type CodeFunction struct {
	Name           string
	ReturnTypes    []CodeProperty
	Parameters     []CodeProperty
	MethodCalls    []CodeCall
	Override       bool
	Annotations    interface{}
	Modifiers      []string
	Creators       []CodeDataStruct
	InnerFunctions []CodeFunction
	CodePosition   CodePosition
}
