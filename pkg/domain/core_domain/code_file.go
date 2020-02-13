package core_domain

type CodeContainer struct {
	FullName       string
	PackageName    string
	Imports        []CodeImport
	Members        []CodeMember
	DataStructures []CodeDataStruct
	Fields         []CodeField
	Containers     []CodeContainer
}
