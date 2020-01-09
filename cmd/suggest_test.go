package cmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func TestSuggest(t *testing.T) {
	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p ../_fixtures/suggest",
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "suggest",
		Cmd:    "suggest",
		Golden: "testdata/suggest_normal.txt",
	}}
	RunTestCmd(t, tests)
}