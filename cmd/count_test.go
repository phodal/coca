package cmd

import (
	"testing"
)

func TestCount(t *testing.T) {
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p .",
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "count",
		cmd:    "count",
		golden: "",
	}}
	runTestCmd(t, tests)
}