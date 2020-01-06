package cmd

import (
	"path/filepath"
	"testing"
)

func TestCount(t *testing.T) {
	abs, _ := filepath.Abs("../_fixtures/arch")
	abs = filepath.FromSlash(abs)

	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + abs,
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "count",
		cmd:    "count -t 1",
		golden: filepath.FromSlash("testdata/count.txt"),
	}}
	runTestCmd(t, tests)
}
