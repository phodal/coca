package trial

type CodeDataStruct struct {
	Name       string
	Type       string
	ID         string
	MemberIds  []string
	Extend     string
	Implements []string

	Annotations     interface{}
	InOutProperties []CodeProperty
	Functions       []CodeFunction
	FunctionCalls   []CodeCall // for field call

	Extension interface{}
	Fields    []CodeField
}

type CodeField struct {
	CodeProperty
	TypeValue string
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
