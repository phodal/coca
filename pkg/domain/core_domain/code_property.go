package core_domain

type CodeProperty struct {
	Modifiers   []string
	Name        string
	TypeName    string
	TypeType    string
	ReturnTypes []CodeProperty
	Parameters  []CodeProperty
}

func NewCodeParameter(typeName string, typeType string) CodeProperty {
	return CodeProperty{
		TypeName: typeName,
		TypeType: typeType,
	}
}

