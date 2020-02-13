package core_domain

type CodePackage struct {
	Name      string
	ID        string
	CodeFiles []CodeContainer
	Extension interface{}
}

type GoCodePackage struct {
	Fields  []CodeField
}
