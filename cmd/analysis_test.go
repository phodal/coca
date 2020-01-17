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

func Test_Analysis_Python(t *testing.T) {
	path := "../_fixtures/grammar/python"

	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -f -l py -p " + path,
		Golden: "testdata/analysis_python.txt",
	}}
	RunTestCmd(t, analysis)
}

func Test_Analysis_TypeScript(t *testing.T) {
	path := "../_fixtures/grammar/typescript"

	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -l ts -p " + path,
		Golden: "testdata/analysis_typescript.txt",
	}}
	RunTestCmd(t, analysis)
}
