package cmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func TestCall(t *testing.T) {
	abs := "../_fixtures/arch"

	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p " + abs,
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "call",
		Cmd:    "call -r com",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}