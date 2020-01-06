package cmd

import (
	"testing"
)

func TestCount(t *testing.T) {
	abs := "../_fixtures/arch"

	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + abs,
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "count",
		cmd:    "count -t 1",
		golden: "testdata/count.txt",
	}}
	runTestCmd(t, tests)
}
