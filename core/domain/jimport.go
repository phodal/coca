package domain

type JImport struct {
	source string
}

func NewJImport(str string) JImport {
	return *&JImport{
		source: str,
	}
}
