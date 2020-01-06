package cmd

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	abs := "../_fixtures/arch"

	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + abs,
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