package cmd

import (
	"testing"
)

func TestVersion(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "version",
		cmd:    "version",
		golden: "",
	}}
	runTestCmd(t, tests)
}