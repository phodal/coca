package domain

type CodeFile struct {
	FullName    string
	PackageName string
	Imports     []string
	Members     []CodeMember
	ClassNodes  []JClassNode
}
