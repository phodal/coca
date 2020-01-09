package cmd

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	analysis := []CmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p ../_fixtures/arch",
		golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		name:   "evaluate",
		cmd:    "evaluate",
		golden: "testdata/evaluate.txt",
	}}
	RunTestCmd(t, tests)
}