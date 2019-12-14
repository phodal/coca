package support

import (
	"strings"
)
var relates []RefactorChangeRelate

func parseRelated(str string) *RefactorChangeRelate {
	result := &RefactorChangeRelate{"", ""}
	splitStr := strings.Split(str, " -> ")
	if len(splitStr) < 2 {
		return nil
	}

	result.NewObj = splitStr[1]
	result.OldObj = splitStr[0]
	return result
}

func ParseRelates(str string) []RefactorChangeRelate {
	relates = nil
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		related := parseRelated(line)
		relates = append(relates, *related)
	}

	return relates
}
