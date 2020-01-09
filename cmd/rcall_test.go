package cmd

import (
	"testing"
)

func TestRCall(t *testing.T) {
	analysis := []CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p ../_fixtures/call",
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		Name:   "rcall",
		Cmd:    "rcall -r com -c com",
		Golden: "testdata/rcall_normal.txt",
	}}
	RunTestCmd(t, tests)
}