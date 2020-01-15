package trial

import "github.com/phodal/coca/pkg/domain"

type CodeFile struct {
	FullName       string
	PackageName    string
	Imports        []string // CodeImports
	Members        []*CodeMember
	DataStructures []CodeDataStruct
	// Deprecated: support for migration only
	ClassNodes []domain.JClassNode
}

type CodeImport struct {
	Source    string
	AsName    string
	UsageName []string
	Scope     string  // function, method or class
}
