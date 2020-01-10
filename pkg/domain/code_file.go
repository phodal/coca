package domain

type CodePosition struct {
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
}

type CodeMember struct {
	ID           string
	Name         string
	ClassNodes   []JClassNode
	Namespace    []string
	FileID       string
	DataStructID string
	Position     CodePosition
}

type CodeFile struct {
	FullName string
	Imports  string
	Members  []CodeMember
}
