package cmd

import (
	"testing"
)

func TestRefactorMove(t *testing.T) {
	tests := []CmdTestCase{{
		name:   "refactor",
		cmd:    "refactor -p . -m .",
		golden: "",
	}}
	RunTestCmd(t, tests)
}

func TestRefactorRename(t *testing.T) {
	tests := []CmdTestCase{{
		name:   "refactor",
		cmd:    "refactor -p . -R . -m .",
		golden: "",
	}}
	RunTestCmd(t, tests)
}