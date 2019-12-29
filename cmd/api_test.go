package cmd

import (
	"path/filepath"
	"testing"
)

func TestApi(t *testing.T) {
	abs, _ := filepath.Abs("../_fixtures/call")
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + abs,
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "api",
		cmd:    "api -c -f -p " + abs,
		golden: "",
	}}
	runTestCmd(t, tests)
}