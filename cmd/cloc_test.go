package cmd

import (
	"testing"
)

func TestCloc(t *testing.T) {
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p .",
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "cloc",
		cmd:    "cloc",
		golden: "",
	}}
	runTestCmd(t, tests)
}