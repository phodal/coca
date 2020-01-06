package cmd

import (
	"testing"
)

func TestArch(t *testing.T) {
	abs := "../_fixtures/arch"

	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + abs,
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "arch",
		cmd:    "arch -P ",
		golden: "",
	}}
	runTestCmd(t, tests)
}