package cmd

import (
	"testing"
)

func TestConcept(t *testing.T) {
	analysis := []CmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p ../_fixtures/examples/api",
		golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		name:   "concept",
		cmd:    "concept",
		golden: "",
	}}
	RunTestCmd(t, tests)
}