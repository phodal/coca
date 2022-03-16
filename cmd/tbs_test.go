package cmd

import (
	"github.com/modernizing/coca/cocatest/testcase"
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
