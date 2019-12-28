package astitodo

import (
	"regexp"
)

// Vars
var (
	regexpAssignee = regexp.MustCompile("^\\([\\w \\._\\+\\-@]+\\)")
	todoIdentifiers = []string{"TODO", "FIXME"}
)

// TODOs represents a set of todos
type TODOs []*TODO

// TODO represents a todo
type TODO struct {
	Assignee string
	Filename string
	Line     int
	Message  []string
}

func Extract(path string, excludedPaths ...string) (todos TODOs, err error) {
	return nil, nil
}
