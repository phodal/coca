package cmd

import (
	"path/filepath"
	"testing"
)

func TestArch(t *testing.T) {
	abs, _ := filepath.Abs("../_fixtures/arch")
	abs = filepath.FromSlash(abs)

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