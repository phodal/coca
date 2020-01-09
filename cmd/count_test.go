package cmd

import (
	"testing"
)

func TestCount(t *testing.T) {
	abs := "../_fixtures/examples/rename"

	analysis := []CmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + abs,
		golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		name:   "count",
		cmd:    "count -t 1",
		golden: "testdata/count.txt",
	}}
	RunTestCmd(t, tests)
}
