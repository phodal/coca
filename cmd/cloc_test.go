package cmd

import (
	"github.com/modernizing/coca/cocatest/testcase"
	"testing"
)

func TestCloc(t *testing.T) {
	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p .",
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "cloc",
		Cmd:    "cloc",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}

func TestClocByDirectory(t *testing.T) {
	abs := "../_fixtures/cloc/normal"

	tests := []testcase.CmdTestCase{{
		Name:   "cloc",
		Cmd:    "cloc " + abs + " --by-directory --include-ext=java,kt",
		Golden: "testdata/cloc_directory.txt",
	}}
	RunTestCmd(t, tests)
}

func TestShouldReturnNullWhenIgnoreDir(t *testing.T) {
	abs := "../_fixtures/cloc/someignore"

	tests := []testcase.CmdTestCase{{
		Name:   "cloc",
		Cmd:    "cloc " + abs + " --by-directory",
		Golden: "testdata/cloc_ignore.txt",
	}}
	RunTestCmd(t, tests)
}

func TestShouldByFileSize(t *testing.T) {
	abs := "../_fixtures/suggest"

	tests := []testcase.CmdTestCase{{
		Name:   "cloc",
		Cmd:    "cloc " + abs + " --top-size=10 --top-file",
		Golden: "testdata/top_file.txt",
	}}
	RunTestCmd(t, tests)
}
