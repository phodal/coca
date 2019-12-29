package cmd

import (
	"path/filepath"
	"testing"
)

func TestCall(t *testing.T) {
	abs, _ := filepath.Abs("../_fixtures/arch")
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + abs,
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "call",
		cmd:    "call",
		golden: "",
	}}
	runTestCmd(t, tests)
}