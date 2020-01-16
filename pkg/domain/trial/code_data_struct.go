package trial

type CodeDataStruct struct {
	Name          string
	Type          string
	ID            string
	MemberIds     []string
	Extend        string
	Implements    []string
	Annotations   interface{}
	Properties    []CodeProperty
	Functions     []CodeFunction
	FunctionCalls []CodeCall
	Extension     interface{}
}

type JavaExtension struct {
	MethodCalls []CodeCall
	Fields      []CodeProperty
	Tag         []interface{}
}

type PythonAnnotation struct {
	Name       string
	Properties []CodeProperty
}

func NewDataStruct() *CodeDataStruct {
	return &CodeDataStruct{}
}

func (d *CodeDataStruct) IsNotEmpty() bool {
	return len(d.Functions) > 0 || len(d.FunctionCalls) > 0
}
