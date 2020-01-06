package cmd

import (
	"testing"
)

func TestCall(t *testing.T) {
	abs := "../_fixtures/arch"

	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + abs,
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