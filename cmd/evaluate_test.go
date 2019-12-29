package cmd

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p .",
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "evaluate",
		cmd:    "evaluate",
		golden: "",
	}}
	runTestCmd(t, tests)
}