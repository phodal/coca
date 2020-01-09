package cmd

import (
	"testing"
)

func TestRCall(t *testing.T) {
	analysis := []CmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p ../_fixtures/call",
		golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		name:   "rcall",
		cmd:    "rcall -r com -c com",
		golden: "testdata/rcall_normal.txt",
	}}
	RunTestCmd(t, tests)
}