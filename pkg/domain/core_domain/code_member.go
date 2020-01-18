package core_domain

type CodeMember struct {
	ID            string
	AliasPackage  string
	Name          string
	Type          string
	Structures    []CodeDataStruct
	FunctionNodes []CodeFunction
	Namespace     []string
	FileID        string
	DataStructID  string
	Position      CodePosition
}

func NewCodeMember() *CodeMember {
	return &CodeMember{}
}

func (c *CodeMember) BuildMemberId() {
	IsDefaultFunction := c.Name == "default"
	if IsDefaultFunction {
		for _, function := range c.FunctionNodes {
			c.ID = c.AliasPackage + "::" + function.Name
		}
	}

	//if c.DataStructID
}
