package cmd

import (
	"testing"
)

func TestTbs(t *testing.T) {
	tests := []CmdTestCase{{
		Name:   "tbs",
		Cmd:    "tbs -p ../_fixtures/tbs/code -s ",
		Golden: "testdata/tbs_normal.txt",
	}}
	RunTestCmd(t, tests)
}