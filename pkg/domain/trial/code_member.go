package trial

import "github.com/phodal/coca/pkg/domain"

type CodeMember struct {
	ID           string
	Name         string
	Type         string
	Properties   []CodeProperty
	ClassNodes   []domain.JClassNode
	Namespace    []string
	FileID       string
	DataStructID string
	Position     CodePosition
}

type CodeProperty struct {
	Modifiers []string
	Name      string
	TypeName  string
	TypeType  string
}
