package suggest

import "github.com/phodal/coca/core/models"

type Suggest struct {
	File    string
	Package string
	Class   string
	Pattern string
	Reason  string
}

func NewSuggest(clz models.JClassNode, pattern, reason string) Suggest {
	return *&Suggest{
		File:    clz.Path,
		Package: clz.Package,
		Class:   clz.Class,
		Pattern: pattern,
		Reason:  reason,
	}
}
