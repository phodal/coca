package support

import (
	"github.com/iancoleman/strcase"
	"strings"
)


type ConceptSegmenter struct {
}

var strMap map[string]int

func SegmentConceptCamelcase(methodsName []string) map[string]int {
	strMap = make(map[string]int)
	for _, name := range methodsName {
		delimited := strcase.ToDelimited(name, '.')
		split := strings.Split(delimited, ".")
		for _, word := range split {
			if strMap[word] == 0 {
				strMap[word] = 1
			} else {
				strMap[word] = strMap[word] + 1
			}
		}
	}

	return strMap
}

