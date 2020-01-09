package cocatest

import (
	"bytes"
	"github.com/mattn/go-shellwords"
	"github.com/phodal/coca/cocatest/testcase"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func RunTestCaseWithCmd(t *testing.T, tests []testcase.CmdTestCase, rootCmd func(out io.Writer) *cobra.Command) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			defer ResetEnv()()

			t.Log("running Cmd: ", tt.Cmd)
			_, output, err := executeActionCommandC(tt.Cmd, rootCmd)
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

func executeActionCommandC(cmd string, rootCmd func(out io.Writer) *cobra.Command) (*cobra.Command, string, error) {
	args, err := shellwords.Parse(cmd)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	command := rootCmd(buf)

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
