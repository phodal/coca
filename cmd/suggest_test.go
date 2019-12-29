package cmd

import (
	"testing"
)

func TestSuggest(t *testing.T) {
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p .",
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "suggest",
		cmd:    "suggest",
		golden: "",
	}}
	runTestCmd(t, tests)
}