package cmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func Test_Analysis_Go(t *testing.T) {
	path := "config"

	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -f -l go -p " + path,
		Golden: "testdata/analysis_go.txt",
	}}
	RunTestCmd(t, analysis)
}
