package cmd

import (
	"github.com/phodal/coca/cocatest"
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func RunTestCmd(t *testing.T, tests []testcase.CmdTestCase) {
	cocatest.RunTestCaseWithCmd(t, tests, NewRootCmd)
}

