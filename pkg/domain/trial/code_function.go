package trial

type CodeFunction struct {
	Name          string
	ReturnTypes   []CodeProperty
	Parameters    []CodeProperty
	MethodCalls   []CodeCall
	Override      bool
	Annotations   []CodeAnnotation
	Modifiers     []string
	Creators      []CodeDataStruct
}

type CodeAnnotation struct {
}
