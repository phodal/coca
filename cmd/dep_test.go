package cmd

import (
	"testing"
)

func Test_Dep_MavenFindUnused(t *testing.T) {
	tests := []CmdTestCase{{
		name:   "dep",
		cmd:    "deps -p ../_fixtures/deps/maven_sample",
		golden: "testdata/deps_maven.txt",
	}}
	RunTestCmd(t, tests)
}

func Test_Dep_GradleFindUnused(t *testing.T) {
	tests := []CmdTestCase{{
		name:   "dep",
		cmd:    "deps -p ../_fixtures/deps/gradle_sample",
		golden: "testdata/deps_gradle.txt",
	}}
	RunTestCmd(t, tests)
}