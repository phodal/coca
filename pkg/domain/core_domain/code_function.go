package core_domain

type CodeFunction struct {
	Name            string
	ReturnTypes     []CodeProperty
	Parameters      []CodeProperty
	MethodCalls     []CodeCall
	Override        bool
	Annotations     interface{}
	Modifiers       []string
	InnerStructures []CodeDataStruct // InnerClass
	InnerFunctions  []CodeFunction
	Extension       interface{}
	Position        CodePosition
}

func (c *CodeFunction) BuildSingleReturnType(typeType string) *CodeProperty {
	return &CodeProperty{
		TypeType: typeType,
	}
}
