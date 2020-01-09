package cmd

import (
	"testing"
)

func TestVersion(t *testing.T) {
	tests := []CmdTestCase{{
		name:   "version",
		cmd:    "version",
		golden: "",
	}}
	RunTestCmd(t, tests)
}