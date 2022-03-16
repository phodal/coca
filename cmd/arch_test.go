package cmd

import (
	"github.com/modernizing/coca/cocatest/testcase"
	"testing"
)

func TestArch(t *testing.T) {
	abs := "../_fixtures/arch"

	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p " + abs,
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "arch",
		Cmd:    "arch -P ",
		Golden: "",
	}}
	RunTestCmd(t, tests)

	header := []testcase.CmdTestCase{{
		Name:   "arch",
		Cmd:    "arch -H -x \"com.phodal\"",
		Golden: "",
	}}
	RunTestCmd(t, header)
}