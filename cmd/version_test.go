package cmd

import (
	"github.com/modernizing/coca/cocatest/testcase"
	"testing"
)

func TestVersion(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "version",
		Cmd:    "version",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}