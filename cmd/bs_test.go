package cmd

import (
	"testing"
)

func TestBadSmell(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "bs",
		cmd:    "bs -s type",
		golden: "",
	}}
	runTestCmd(t, tests)
}