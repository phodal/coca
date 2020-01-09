package cmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func Test_Dep_MavenFindUnused(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "dep",
		Cmd:    "deps -p ../_fixtures/deps/maven_sample",
		Golden: "testdata/deps_maven.txt",
	}}
	RunTestCmd(t, tests)
}

func Test_Dep_GradleFindUnused(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "dep",
		Cmd:    "deps -p ../_fixtures/deps/gradle_sample",
		Golden: "testdata/deps_gradle.txt",
	}}
	RunTestCmd(t, tests)
}