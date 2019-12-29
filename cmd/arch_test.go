package cmd

import (
	"testing"
)

func TestArch(t *testing.T) {
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p .",
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "arch",
		cmd:    "arch",
		golden: "",
	}}
	runTestCmd(t, tests)
}