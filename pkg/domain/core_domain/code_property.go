package core_domain

type CodeProperty struct {
	Modifiers   []string
	ParamName   string
	TypeValue   string
	TypeType    string
	ReturnTypes []CodeProperty
	Parameters  []CodeProperty
}

func NewCodeParameter(typeType string, typeValue string) CodeProperty {
	return CodeProperty{
		TypeValue: typeValue,
		TypeType:  typeType,
	}
}

