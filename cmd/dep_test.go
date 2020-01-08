package cmd

import (
	"testing"
)

func TestDepFindUnused(t *testing.T) {
	tests := []cmdTestCase{{
		name:   "dep",
		cmd:    "deps -p ../_fixtures/deps/maven_sample",
		golden: "testdata/deps_maven.txt",
	}}
	runTestCmd(t, tests)
}