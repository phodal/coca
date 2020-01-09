package cmd

import (
	"bytes"
	"github.com/mattn/go-shellwords"
	"github.com/phodal/coca/cocatest/testcase"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func RunTestCmd(t *testing.T, tests []testcase.CmdTestCase) {
	RunTestCaseWithCmd(t, tests)
}

func RunTestCaseWithCmd(t *testing.T, tests []testcase.CmdTestCase) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			defer ResetEnv()()

			t.Log("running Cmd: ", tt.Cmd)
			_, output, err := executeActionCommandC(tt.Cmd)
			if (err != nil) != tt.WantError {
				t.Errorf("expected error, got '%v'", err)
			}
			if tt.Golden != "" {
				abs, _ := filepath.Abs(tt.Golden)
				slash := filepath.FromSlash(abs)
				AssertGoldenString(t, output, slash)
			}
		})
	}
}

func executeActionCommandC(cmd string) (*cobra.Command, string, error) {
	args, err := shellwords.Parse(cmd)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	command := NewRootCmd(buf)

	command.SetArgs(args)

	c, err := command.ExecuteC()

	return c, buf.String(), err
}

func ResetEnv() func() {
	origEnv := os.Environ()
	return func() {
		os.Clearenv()
		for _, pair := range origEnv {
			kv := strings.SplitN(pair, "=", 2)
			os.Setenv(kv[0], kv[1])
		}
	}
}
