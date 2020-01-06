package cmd

import (
	"path/filepath"
	"testing"
)

func TestEvaluate(t *testing.T) {
	abs, _ := filepath.Abs("../_fixtures/arch")
	abs = filepath.FromSlash(abs)

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