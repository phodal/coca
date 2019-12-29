package cmd

import (
	"testing"
)

func TestGit(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "git",
		cmd:    "git -a -f -t",
		golden: "",
	}}
	runTestCmd(t, tests)
}