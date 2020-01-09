package cmd

import (
	"testing"
)

func Test_ShouldOutputCount(t *testing.T) {
	tests := []CmdTestCase{{
		Name:   "todo",
		Cmd:    "todo -p ../_fixtures/todo",
		Golden: "testdata/todo_normal.txt",
	}}
	RunTestCmd(t, tests)
}

//TODO: update func for CI which clone depth = 1
func TestTodo(t *testing.T) {
	tests := []CmdTestCase{{
		Name:   "todo",
		Cmd:    "todo -p ../_fixtures/todo -g",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}
