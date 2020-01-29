package cmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func TestEvaluate(t *testing.T) {
	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p ../_fixtures/grammar/java/arch",
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "evaluate",
		Cmd:    "evaluate",
		Golden: "testdata/evaluate.txt",
	}}
	RunTestCmd(t, tests)
}