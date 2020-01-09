package cmd

import (
	"testing"
)

func TestBadSmell(t *testing.T) {
	abs := "../_fixtures/bs"

	tests := []CmdTestCase{{
		Name:   "bs",
		Cmd:    "bs -s type -p " + abs,
		Golden: "",
	}}
	RunTestCmd(t, tests)
}