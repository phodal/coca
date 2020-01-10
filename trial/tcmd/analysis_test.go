package tcmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func Test_AnalysisCmd(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p ../../_fixtures/ts/grammar",
		Golden: "",
	}}
	RunTrialTestCmd(t, tests)
}