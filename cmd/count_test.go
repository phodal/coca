package cmd

import (
	"path/filepath"
	"testing"
)

func TestCount(t *testing.T) {
	abs, _ := filepath.Abs("../_fixtures/arch")
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + abs,
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "count",
		cmd:    "count",
		golden: "",
	}}
	runTestCmd(t, tests)
}