package cmd

import (
	"testing"
)

func TestApi(t *testing.T) {
	path := "../_fixtures/call"

	analysis := []CmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + path,
		golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		name:   "api",
		cmd:    "api -c -f -p " + path,
		golden: "testdata/api.txt",
	}}
	RunTestCmd(t, tests)
}

func Test_ApiWithSortRemove(t *testing.T) {
	path := "../_fixtures/call"

	analysis := []CmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + path,
		golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []CmdTestCase{{
		name:   "api",
		cmd:    "api -c -s -r com.phodal.pholedge.book. -p" + path,
		golden: "testdata/api_sort_remove.txt",
	}}
	RunTestCmd(t, tests)
}