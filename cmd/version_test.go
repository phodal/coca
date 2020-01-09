package cmd

import (
	"testing"
)

func TestVersion(t *testing.T) {
	tests := []CmdTestCase{{
		Name:   "version",
		Cmd:    "version",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}