package cmd

import (
	"testing"
)

func TestCall(t *testing.T) {
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p .",
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "call",
		cmd:    "call",
		golden: "",
	}}
	runTestCmd(t, tests)
}