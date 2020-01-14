package tcmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func Test_AnalysisGoCmd(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "go",
		Cmd:    "go -p ../tcmd",
		Golden: "",
	}}
	RunTrialTestCmd(t, tests)
}