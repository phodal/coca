package cmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func TestCloc(t *testing.T) {
	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p .",
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "cloc",
		Cmd:    "cloc",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}