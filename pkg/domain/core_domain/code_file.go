package core_domain

type CodeFile struct {
	FullName       string
	PackageName    string
	Imports        []CodeImport
	Members        []*CodeMember
	DataStructures []CodeDataStruct
}

type CodeImport struct {
	Source     string
	AsName     string
	ImportName string
	UsageName  []string
	Scope      string // function, method or class
}
