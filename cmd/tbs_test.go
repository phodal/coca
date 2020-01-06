package cmd

import (
	"testing"
)

func TestTbs(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "tbs",
		cmd:    "tbs -p ../_fixtures/tbs/code -s ",
		golden: "testdata/tbs_normal.txt",
	}}
	runTestCmd(t, tests)
}