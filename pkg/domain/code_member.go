package domain

type CodeMember struct {
	ID           string
	Name         string
	Type         string
	Properties   []CodeProperty
	ClassNodes   []JClassNode
	Namespace    []string
	FileID       string
	DataStructID string
	Position     CodePosition
}

type CodeProperty struct {
	Modifiers []string
	Name      string
	Type      string
}
