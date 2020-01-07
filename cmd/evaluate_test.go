package cmd

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p ../_fixtures/arch",
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "evaluate",
		cmd:    "evaluate",
		golden: "testdata/evaluate.txt",
	}}
	runTestCmd(t, tests)
}