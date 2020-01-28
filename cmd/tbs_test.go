package cmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func TestTbs(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "tbs",
		Cmd:    "tbs -p ../_fixtures/tbs/usecases -s ",
		Golden: "testdata/tbs_normal.txt",
	}}
	RunTestCmd(t, tests)
}
