package cmd

import (
	"github.com/modernizing/coca/cocatest/testcase"
	"testing"
)

func TestRefactorMove(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "refactor",
		Cmd:    "refactor -p . -m .",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}

func TestRefactorRename(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "refactor",
		Cmd:    "refactor -p . -R . -m .",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}
