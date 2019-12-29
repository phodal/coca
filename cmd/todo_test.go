package cmd

import (
	"testing"
)

func TestTodo(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "todo",
		cmd:    "todo -p . -g",
		golden: "",
	}}
	runTestCmd(t, tests)
}