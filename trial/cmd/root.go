package cmd

import (
	"github.com/spf13/cobra"
	"io"
)

var (
	output       io.Writer
	trialRootCmd = &cobra.Command{
		Use:   "coca",
		Short: "A generator for Cobra based Applications",
		Long:  `coca`,
	}
)

func NewTrialRootCmd(out io.Writer) *cobra.Command {
	output = out
	trialRootCmd.SetOut(out)
	return trialRootCmd
}
