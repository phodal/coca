package cmd

import (
	"testing"
)

func TestArch(t *testing.T) {
	abs := "../_fixtures/arch"

	analysis := []CmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + abs,
		golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		name:   "arch",
		cmd:    "arch -P ",
		golden: "",
	}}
	RunTestCmd(t, tests)
}