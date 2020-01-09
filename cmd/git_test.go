package cmd

import (
	"testing"
)

// Todo: fake it
func TestGit(t *testing.T) {
	tests := []CmdTestCase{{
		name:   "git",
		cmd:    "git -a -f -t -b -o -r com -s 10 -m",
		golden: "",
	}}
	RunTestCmd(t, tests)
}