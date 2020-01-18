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

func (c *CodeMember) BuildMemberId() {
	IsDefaultFunction := c.Name == "default"
	if IsDefaultFunction {
		//c.ID = c.
	}
}
