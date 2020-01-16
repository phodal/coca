package jdomain

type JImport struct {
	Source string
}

func NewJImport(str string) JImport {
	return JImport{
		Source: str,
	}
}
