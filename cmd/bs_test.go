package cmd

import (
	"testing"
)

func TestBadSmell(t *testing.T) {
	abs := "../_fixtures/bs"

	tests := []CmdTestCase{{
		name:   "bs",
		cmd:    "bs -s type -p " + abs,
		golden: "",
	}}
	RunTestCmd(t, tests)
}