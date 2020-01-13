package domain

type CodePosition struct {
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
}

type CodeFile struct {
	FullName   string
	Imports    []string
	Members    []CodeMember
	ClassNodes []JClassNode
}
