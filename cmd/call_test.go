package cmd

import (
	"testing"
)

func TestCall(t *testing.T) {
	abs := "../_fixtures/arch"

	analysis := []CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p " + abs,
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		Name:   "call",
		Cmd:    "call -r com",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}