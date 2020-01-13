package domain

type CodeMember struct {
	ID           string
	Name         string
	ClassNodes   []JClassNode
	Namespace    []string
	FileID       string
	DataStructID string
	Position     CodePosition
}
