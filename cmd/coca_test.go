package cmd

import (
	"github.com/modernizing/coca/cocatest"
	"github.com/modernizing/coca/cocatest/testcase"
	"testing"
)

func RunTestCmd(t *testing.T, tests []testcase.CmdTestCase) {
	cocatest.RunTestCaseWithCmd(t, tests, NewRootCmd)
}

