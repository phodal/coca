package cmd

import (
	"path/filepath"
	"testing"
)

func TestBadSmell(t *testing.T) {
	abs, _ := filepath.Abs("../_fixtures/bs")
	tests := []cmdTestCase{{
		name:   "bs",
		cmd:    "bs -s type -p " + abs,
		golden: "",
	}}
	runTestCmd(t, tests)
}