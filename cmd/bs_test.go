package cmd

import (
	"testing"
)

func TestBadSmell(t *testing.T) {
	abs := "../_fixtures/bs"

	tests := []cmdTestCase{{
		name:   "bs",
		cmd:    "bs -s type -p " + abs,
		golden: "",
	}}
	runTestCmd(t, tests)
}