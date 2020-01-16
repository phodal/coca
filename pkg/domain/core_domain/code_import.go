package core_domain

type CodeImport struct {
	Source     string
	AsName     string
	ImportName string
	UsageName  []string
	Scope      string // function, method or class
}

func NewJImport(str string) CodeImport {
	return CodeImport{
		Source: str,
	}
}
