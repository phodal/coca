package core_domain

type CodeProperty struct {
	Modifiers   []string
	Name        string
	TypeName    string
	TypeType    string
	ReturnTypes []CodeProperty
	Parameters  []CodeProperty
}
