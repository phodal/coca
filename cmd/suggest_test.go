package cmd

import (
	"testing"
)

func TestSuggest(t *testing.T) {
	analysis := []CmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p ../_fixtures/suggest",
		golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		name:   "suggest",
		cmd:    "suggest",
		golden: "testdata/suggest_normal.txt",
	}}
	RunTestCmd(t, tests)
}