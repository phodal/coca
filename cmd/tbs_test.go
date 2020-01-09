package cmd

import (
	"testing"
)

func TestTbs(t *testing.T) {
	tests := []CmdTestCase{{
		name:   "tbs",
		cmd:    "tbs -p ../_fixtures/tbs/code -s ",
		golden: "testdata/tbs_normal.txt",
	}}
	RunTestCmd(t, tests)
}