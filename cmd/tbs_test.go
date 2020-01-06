package cmd

import (
	"testing"
)

func TestTbs(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "tbs",
		cmd:    "tbs -p . -s",
		golden: "",
	}}
	runTestCmd(t, tests)
}