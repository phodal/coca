package core_domain

type CodeCall struct {
	Package    string
	Type       string
	Class      string
	MethodName string
	Parameters []CodeProperty
	Position   CodePosition
}
