package trial

type CodeDataStruct struct {
	Name            string
	Type            string
	ID              string
	MemberIds       []string
	Extend          string
	Implements      []string
	InnerStructures []CodeDataStruct
	Annotations     interface{}
	InOutProperties []CodeProperty
	Functions       []CodeFunction
	FunctionCalls   []CodeCall // for field call
	Fields          []CodeField
	Extension       interface{}
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
