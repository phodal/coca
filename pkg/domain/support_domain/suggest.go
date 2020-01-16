package api_domain

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
	"strings"
)

type Suggest struct {
	File    string
	Package string
	Class   string
	Pattern string
	Reason  string
	Size    int
	Line    int
}

func NewSuggest(clz core_domain.CodeDataStruct, pattern, reason string) Suggest {
	return Suggest{
		File:    clz.FilePath,
		Package: clz.Package,
		Class:   clz.NodeName,
		Pattern: pattern,
		Reason:  reason,
	}
}

func MergeSuggest(clz core_domain.CodeDataStruct, currentSuggestList []Suggest) Suggest {
	var suggest = NewSuggest(clz, "", "")
	for _, s := range currentSuggestList {
		if !strings.Contains(suggest.Pattern, s.Pattern) {
			if suggest.Pattern != "" {
				suggest.Pattern = suggest.Pattern + ", " + s.Pattern
			} else {
				suggest.Pattern = s.Pattern
			}
		}

		if !strings.Contains(suggest.Reason, s.Reason) {
			if suggest.Reason != "" {
				suggest.Reason = suggest.Reason + ", " + s.Reason
			} else {
				suggest.Reason = s.Reason
			}
		}
	}
	return suggest
}
