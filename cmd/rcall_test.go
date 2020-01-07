package cmd

import (
	"testing"
)

func TestRCall(t *testing.T) {
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p ../_fixtures/call",
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "rcall",
		cmd:    "rcall -r com -c com",
		golden: "testdata/rcall_normal.txt",
	}}
	runTestCmd(t, tests)
}