package cmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func TestVersion(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "version",
		Cmd:    "version",
		Golden: "",
	}}
	RunTrialTestCmd(t, tests)
}