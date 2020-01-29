package cmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func TestCount(t *testing.T) {
	abs := "../_fixtures/grammar/java/examples/rename"

	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p " + abs,
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "count",
		Cmd:    "count -t 1",
		Golden: "testdata/count.txt",
	}}
	RunTestCmd(t, tests)
}
