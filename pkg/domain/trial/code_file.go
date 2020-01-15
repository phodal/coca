package trial

import "github.com/phodal/coca/pkg/domain"

type CodeFile struct {
	FullName       string
	PackageName    string
	Imports        []CodeImport
	Members        []*CodeMember
	DataStructures []CodeDataStruct
	// Deprecated: support for migration only
	ClassNodes []domain.JClassNode
}

type CodeImport struct {
	Source     string
	AsName     string
	ImportName string
	UsageName  []string
	Scope      string // function, method or class
}
