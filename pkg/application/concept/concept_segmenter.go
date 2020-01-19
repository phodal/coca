package concept

import (
	"github.com/iancoleman/strcase"
	"regexp"
	"strings"
)

type ConceptSegmenter struct {
}

var strMap map[string]int

func SegmentCamelcase(methodsName []string) map[string]int {
	strMap = make(map[string]int)
	for _, name := range methodsName {
		// get, set
		if (strings.HasSuffix("set", name) || strings.HasSuffix("get", name)) && len(name) > 3 {
			domainName := name[3:]
			if strMap[domainName] == 0 {
				strMap[domainName] = 1
			} else {
				strMap[domainName] = strMap[domainName] + 1
			}
		} else {
			delimited := strcase.ToDelimited(name, '.')
			split := strings.Split(delimited, ".")
			for _, word := range split {
				if FilterString(word) == "" {
					continue
				}
				if strMap[word] == 0 {
					strMap[word] = 1
				} else {
					strMap[word] = strMap[word] + 1
				}
			}
		}
	}

	return strMap
}

func FilterString(str string) string {
	var digitCheck = regexp.MustCompile(`^[0-9]+$`)

	if digitCheck.MatchString(str) {
		return ""
	}

	return strings.TrimSpace(str)
}
