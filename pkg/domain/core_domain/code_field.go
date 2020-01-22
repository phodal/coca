package core_domain

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
