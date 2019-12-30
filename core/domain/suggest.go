package domain

type Suggest struct {
	File    string
	Package string
	Class   string
	Pattern string
	Reason  string
	Size    int
	Line    int
}

func NewSuggest(clz JClassNode, pattern, reason string) Suggest {
	return *&Suggest{
		File:    clz.Path,
		Package: clz.Package,
		Class:   clz.Class,
		Pattern: pattern,
		Reason:  reason,
	}
}
