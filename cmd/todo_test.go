package cmd

import (
	"testing"
)

func Test_ShouldOutputCount(t *testing.T) {
	tests := []CmdTestCase{{
		name:   "todo",
		cmd:    "todo -p ../_fixtures/todo",
		golden: "testdata/todo_normal.txt",
	}}
	RunTestCmd(t, tests)
}

//TODO: update func for CI which clone depth = 1
func TestTodo(t *testing.T) {
	tests := []CmdTestCase{{
		name:   "todo",
		cmd:    "todo -p ../_fixtures/todo -g",
		golden: "",
	}}
	RunTestCmd(t, tests)
}
