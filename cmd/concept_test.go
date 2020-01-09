package cmd

import (
	"testing"
)

func TestConcept(t *testing.T) {
	analysis := []CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p ../_fixtures/examples/api",
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		Name:   "concept",
		Cmd:    "concept",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}