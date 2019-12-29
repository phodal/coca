package cmd

import (
	"testing"
)

func TestApi(t *testing.T) {
	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p .",
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "api",
		cmd:    "api -c -r com",
		golden: "",
	}}
	runTestCmd(t, tests)
}