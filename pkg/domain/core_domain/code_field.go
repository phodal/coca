package core_domain

type CodeField struct {
	TypeType  string
	TypeValue string
	Modifiers []string
}

func NewJField(typeType string, typeValue string, modifier string) CodeField {
	property := CodeField{
		TypeValue: typeValue,
		TypeType:  typeType,
	}

	if modifier != "" {
		property.Modifiers = append(property.Modifiers, modifier)
	}

	return property
}
