package main

import (
	"github.com/phodal/coca/trial/tcmd"
	"os"
)

func main() {
	output := os.Stdout
	rootCmd := tcmd.NewTrialRootCmd(output)
	_ = rootCmd.Execute()
}
