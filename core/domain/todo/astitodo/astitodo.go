package astitodo

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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

// Extract walks through an input path and extracts TODOs from all files it encounters
func Extract(path string, excludedPaths ...string) (todos TODOs, err error) {
	err = todos.extract(path, excludedPaths...)
	return
}

func (todos *TODOs) extract(path string, excludedPaths ...string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		// Process error
		if err != nil {
			return err
		}

		// Skip excluded paths
		for _, p := range excludedPaths {
			if p == path {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		// Skip vendor and all directories beginning with a .
		if info.IsDir() && (info.Name() == "vendor" || (len(info.Name()) > 1 && info.Name()[0] == '.')) {
			return filepath.SkipDir
		}

		// Only process go files
		if !info.IsDir() && filepath.Ext(path) != ".go" {
			return nil
		}

		// Everything is fine here, extract if path is a file
		if !info.IsDir() {
			if err = todos.extractFile(path); err != nil {
				return err
			}
		}
		return nil
	})
}

func (todos *TODOs) extractFile(filename string) (err error) {
	// Parse file and create the AST
	var fset = token.NewFileSet()
	var f *ast.File
	if f, err = parser.ParseFile(fset, filename, nil, parser.ParseComments); err != nil {
		return
	}

	// Loop in comment groups
	for _, cg := range f.Comments {
		// Loop in comments
		var todo *TODO
		var TODOFound bool
		for _, c := range cg.List {
			// Loop in lines
			for i, l := range strings.Split(c.Text, "\n") {
				// Init text
				var t = strings.TrimSpace(l)
				if strings.HasPrefix(t, "//") || strings.HasPrefix(t, "/*") || strings.HasPrefix(t, "*/") {
					t = strings.TrimSpace(t[2:])
				}

				// To do found
				if length, isTodo := isTodoIdentifier(t); isTodo {
					// Init to do
					todo = &TODO{Filename: filename, Line: fset.Position(c.Slash).Line + i}
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
					todo.Message = append(todo.Message, t)
					*todos = append(*todos, todo)
					TODOFound = true
				} else if TODOFound && len(t) > 0 {
					todo.Message = append(todo.Message, t)
				} else {
					TODOFound = false
				}
			}
		}
	}
	return
}

func isTodoIdentifier(s string) (int, bool) {
	for _, indent := range todoIdentifiers {
		if strings.HasPrefix(strings.ToUpper(s), indent) {
			return len(indent), true
		}
	}
	return 0, false
}

// AssignedTo returns TODOs which are assigned to the specified assignee
func (todos TODOs) AssignedTo(assignees ...string) (filteredTODOs TODOs) {
	for _, t := range todos {
		for _, assignee := range assignees {
			if assignee == t.Assignee {
				filteredTODOs = append(filteredTODOs, t)
			}
		}
	}

	return
}

// WriteText writes the TODOs as text to the specified writer
func (todos TODOs) WriteText(w io.Writer) (err error) {
	for _, t := range todos {
		if t.Assignee != "" {
			if _, err = io.WriteString(w, fmt.Sprintf("Assignee: %s\n", t.Assignee)); err != nil {
				return
			}
		}

		if _, err = io.WriteString(w, fmt.Sprintf("Message: %s\nFile:%s:%d\n\n", strings.Join(t.Message, "\n"), t.Filename, t.Line)); err != nil {
			return
		}
	}

	return
}

// WriteCSV writes the TODOs as CSV to the specified writer
// The columns are "Filename", "Line", "Assignee" and "Message" (which can contain newlines)
func (todos TODOs) WriteCSV(w io.Writer) (err error) {
	var c = csv.NewWriter(w)

	// Write the headings for the document
	if err = c.Write([]string{"Filename", "Line", "Assignee", "Message"}); err != nil {
		return
	}

	for _, t := range todos {
		err = c.Write([]string{
			t.Filename,
			strconv.Itoa(t.Line),
			t.Assignee,
			strings.Join(t.Message, "\n"),
		})

		if err != nil {
			return
		}
	}

	c.Flush()

	return
}

// WriteJSON writes the TODOs as JSON to the specified writer
func (todos TODOs) WriteJSON(w io.Writer) (err error) {
	enc := json.NewEncoder(w)
	err = enc.Encode(todos)
	return
}
