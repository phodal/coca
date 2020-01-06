package cmd

import (
	"testing"
)

func TestSuggest(t *testing.T) {
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p ../_fixtures/suggest",
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "suggest",
		cmd:    "suggest",
		golden: "testdata/suggest_normal.txt",
	}}
	runTestCmd(t, tests)
}