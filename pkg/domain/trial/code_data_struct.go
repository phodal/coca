package trial

type CodeDataStruct struct {
	Name        string
	ID          string
	MemberIds   []string
	Extend      string
	Implements  []string
	Annotations []interface{}
	Properties  []CodeProperty
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
