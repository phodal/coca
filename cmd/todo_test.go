package cmd

import (
	"testing"
)

func Test_ShouldOutputCount(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "todo",
		cmd:    "todo -p ../_fixtures/todo",
		golden: "testdata/todo_normal.txt",
	}}
	runTestCmd(t, tests)
}

//TODO: update func for CI which clone depth = 1
func TestTodo(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "todo",
		cmd:    "todo -p ../_fixtures/todo -g",
		golden: "",
	}}
	runTestCmd(t, tests)
}
