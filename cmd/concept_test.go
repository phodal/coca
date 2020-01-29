package cmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func TestConcept(t *testing.T) {
	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p ../_fixtures/grammar/java/examples/api",
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "concept",
		Cmd:    "concept",
		Golden: "testdata/concept.txt",
	}}
	RunTestCmd(t, tests)
}