package core_domain

type CodeProperty struct {
	Modifiers   []string
	Name        string
	TypeValue   string
	TypeType    string
	ReturnTypes []CodeProperty
	Parameters  []CodeProperty
}

func NewCodeParameter(typeValue string, typeType string) CodeProperty {
	return CodeProperty{
		TypeValue: typeValue,
		TypeType:  typeType,
	}
}

type CodeField struct {
	TypeType  string
	TypeValue string
	Modifiers []string
}

func NewJField(typeValue string, typeType string, modifier string) CodeField {
	property := CodeField{
		TypeValue: typeValue,
		TypeType:  typeType,
	}
	property.Modifiers = append(property.Modifiers, modifier)

	return property
}
