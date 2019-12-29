package cmd

import (
	"testing"
)

func TestRefactorMove(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "refactor",
		cmd:    "refactor -p . -m .",
		golden: "",
	}}
	runTestCmd(t, tests)
}

func TestRefactorRename(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "refactor",
		cmd:    "refactor -p . -R . -m .",
		golden: "",
	}}
	runTestCmd(t, tests)
}