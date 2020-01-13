package trial

type CodeFunction struct {
	Name          string
	ReturnTypes   []CodeProperty
	Parameters    []CodeProperty
	MethodCalls   []CodeMethodCall
	Override      bool
	Annotations   []CodeAnnotation
	Modifiers     []string
	Creators      []CodeDataStruct
}

type CodeAnnotation struct {
}

type CodeMethodCall struct {
}

type CodeParameter struct {
}
