package cmd

import (
	"github.com/spf13/cobra"
	"io"
)

var (
	output io.Writer
	rootCmd = &cobra.Command{
		Use:   "coca",
		Short: "A generator for Cobra based Applications",
		Long:  `coca`,
	}
)

func NewRootCmd(out io.Writer) *cobra.Command {
	output = out
	rootCmd.SetOut(out)
	return rootCmd
}
