package cmd

import (
	"testing"
)

func TestConcept(t *testing.T) {
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p .",
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "concept",
		cmd:    "concept",
		golden: "",
	}}
	runTestCmd(t, tests)
}