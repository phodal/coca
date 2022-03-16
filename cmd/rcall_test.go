package cmd

import (
	"github.com/modernizing/coca/cocatest/testcase"
	"testing"
)

func TestRCall(t *testing.T) {
	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p ../_fixtures/call",
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "rcall",
		Cmd:    "rcall -r com -c com",
		Golden: "testdata/rcall_normal.txt",
	}}
	RunTestCmd(t, tests)
}
