package core_domain

type CodePackage struct {
	Name      string
	ID        string
	CodeFiles []CodeFile
	Extension interface{}
}

type GoCodePackage struct {
	Fields  []CodeField
}
