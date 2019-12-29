package cmd

import (
	"testing"
)

func TestRCall(t *testing.T) {
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p .",
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "rcall",
		cmd:    "rcall -c com",
		golden: "",
	}}
	runTestCmd(t, tests)
}