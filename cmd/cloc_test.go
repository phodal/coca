package cmd

import (
	"testing"
)

func TestCloc(t *testing.T) {
	analysis := []CmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p .",
		golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		name:   "cloc",
		cmd:    "cloc",
		golden: "",
	}}
	RunTestCmd(t, tests)
}