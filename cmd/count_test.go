package cmd

import (
	"testing"
)

func TestCount(t *testing.T) {
	abs := "../_fixtures/examples/rename"

	analysis := []CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p " + abs,
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		Name:   "count",
		Cmd:    "count -t 1",
		Golden: "testdata/count.txt",
	}}
	RunTestCmd(t, tests)
}
