package cmd

import (
	"testing"
)

func TestTodo(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "todo",
		cmd:    "todo -p ../_fixtures/todo -g",
		golden: "testdata/todo_normal.txt",
	}}
	runTestCmd(t, tests)
}