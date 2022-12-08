package astitodo

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"regexp"
	"strings"
)

var todoIdentifiers = []string{"TODO", "FIXME"}

type TODO struct {
	Assignee string
	Filename string
	Line     int
	Message  string
}

var (
	assignRegStr   = "^\\([\\w \\._\\+\\-@]+\\)"
	regexpAssignee = regexp.MustCompile(assignRegStr)
)

func ParseComment(token antlr.Token, filename string) *TODO {
	comment := token.GetText()

	var t = strings.TrimSpace(comment)
	// todo: add todo list
	if strings.HasPrefix(t, "//") || strings.HasPrefix(t, "/*") || strings.HasPrefix(t, "*/") || strings.HasPrefix(t, "#") {
		t = strings.TrimSpace(t[2:])
	}

	if length, isTodo := IsTodoIdentifier(t); isTodo {
		todo := &TODO{Filename: filename, Line: token.GetLine()}
		t = strings.TrimSpace(t[length:])
		if strings.HasPrefix(t, ":") {
			t = strings.TrimLeft(t, ":")
			t = strings.TrimSpace(t)
		}

		// Look for assignee
		if todo.Assignee = regexpAssignee.FindString(t); todo.Assignee != "" {
			t = strings.TrimSpace(t[len(todo.Assignee):])
			if strings.HasPrefix(t, ":") {
				t = strings.TrimLeft(t, ":")
				t = strings.TrimSpace(t)
			}
			todo.Assignee = todo.Assignee[1 : len(todo.Assignee)-1]
		}

		// Append text
		todo.Message = handleForMultipleLine(t)

		return todo
	}

	return nil
}

// todo: handle for letter
func handleForMultipleLine(t string) string {
	t = strings.ReplaceAll(t, "*/", " ")
	t = strings.ReplaceAll(t, "*", " ")
	t = strings.ReplaceAll(t, "\n", " ")
	return t
}

func IsTodoIdentifier(s string) (int, bool) {
	for _, indent := range todoIdentifiers {
		if strings.HasPrefix(strings.ToUpper(s), indent) {
			return len(indent), true
		}
	}
	return 0, false
}
