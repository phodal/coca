package cmd

import (
	"testing"
)

func TestSuggest(t *testing.T) {
	analysis := []CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p ../_fixtures/suggest",
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		Name:   "suggest",
		Cmd:    "suggest",
		Golden: "testdata/suggest_normal.txt",
	}}
	RunTestCmd(t, tests)
}