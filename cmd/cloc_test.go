package cmd

import (
	"testing"
)

func TestCloc(t *testing.T) {
	analysis := []CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p .",
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		Name:   "cloc",
		Cmd:    "cloc",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}