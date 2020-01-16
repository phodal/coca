package core_domain

type CodeMember struct {
	ID            string
	Name          string
	Type          string
	Structures    []CodeDataStruct
	FunctionNodes []CodeFunction
	Namespace     []string
	FileID        string
	DataStructID  string
	Position      CodePosition
}
