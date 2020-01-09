package cmd

import (
	"github.com/spf13/cobra"
	"io"
)

var (
	output       io.Writer
	trialRootCmd = &cobra.Command{
		Use:   "cots",
		Short: "A generator for Cobra based Applications",
		Long:  `cots`,
	}
)

func NewTrialRootCmd(out io.Writer) *cobra.Command {
	output = out
	trialRootCmd.SetOut(out)
	return trialRootCmd
}
