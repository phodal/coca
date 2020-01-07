package cmd

import (
	"testing"
)

func TestApi(t *testing.T) {
	path := "../_fixtures/call"

	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + path,
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "api",
		cmd:    "api -c -f -p " + path,
		golden: "testdata/api.txt",
	}}
	runTestCmd(t, tests)
}

func Test_ApiWithSortRemove(t *testing.T) {
	path := "../_fixtures/call"

	analysis := []cmdTestCase{{
		name:   "analysis",
		cmd:    "analysis -p " + path,
		golden: "",
	}}
	runTestCmd(t, analysis)

	tests := []cmdTestCase{{
		name:   "api",
		cmd:    "api -c -s -r com.phodal.pholedge.book. -p" + path,
		golden: "testdata/api_sort_remove.txt",
	}}
	runTestCmd(t, tests)
}