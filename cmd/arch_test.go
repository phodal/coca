package cmd

import (
	"testing"
)

func TestArch(t *testing.T) {
	abs := "../_fixtures/arch"

	analysis := []CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p " + abs,
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		Name:   "arch",
		Cmd:    "arch -P ",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}