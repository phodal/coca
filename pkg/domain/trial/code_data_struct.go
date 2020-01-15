package trial

type CodeDataStruct struct {
	Name        string
	ID          string
	MemberIds   []string
	Properties  []CodeProperty
	Extend      string
	Implements  []string
	Annotations []interface{}
	Extension   interface{}
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
