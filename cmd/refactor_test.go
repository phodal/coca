package cmd

import (
	"testing"
)

func TestRefactorMove(t *testing.T) {
	tests := []CmdTestCase{{
		Name:   "refactor",
		Cmd:    "refactor -p . -m .",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}

func TestRefactorRename(t *testing.T) {
	tests := []CmdTestCase{{
		Name:   "refactor",
		Cmd:    "refactor -p . -R . -m .",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}