package cmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func Test_ShouldOutputCount(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "todo",
		Cmd:    "todo -p ../_fixtures/todo",
		Golden: "testdata/todo_normal.txt",
	}}
	RunTestCmd(t, tests)
}

func Test_ShouldFilterTodo(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "todo",
		Cmd:    "todo -p ../_fixtures/todo --ext=.phodal",
		Golden: "testdata/todo_filter.txt",
	}}
	RunTestCmd(t, tests)
}

//TODO: update func for CI which clone depth = 1
func TestTodo(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "todo",
		Cmd:    "todo -p ../_fixtures/todo -g",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}
