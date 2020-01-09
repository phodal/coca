package cmd

import (
	"testing"
)

func TestCall(t *testing.T) {
	abs := "../_fixtures/arch"

	analysis := []CmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + abs,
		golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		name:   "call",
		cmd:    "call -r com",
		golden: "",
	}}
	RunTestCmd(t, tests)
}