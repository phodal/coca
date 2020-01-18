package core_domain

type CodeFile struct {
	FullName       string
	PackageName    string
	Imports        []CodeImport
	Members        []CodeMember
	DataStructures []CodeDataStruct
	Fields         []CodeField
}
