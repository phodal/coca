package tcmd

import (
	"github.com/phodal/coca/cocatest"
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func RunTrialTestCmd(t *testing.T, tests []testcase.CmdTestCase) {
	cocatest.RunTestCaseWithCmd(t, tests, NewTrialRootCmd)
}

