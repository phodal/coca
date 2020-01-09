package cmd

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	analysis := []CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p ../_fixtures/arch",
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		Name:   "evaluate",
		Cmd:    "evaluate",
		Golden: "testdata/evaluate.txt",
	}}
	RunTestCmd(t, tests)
}