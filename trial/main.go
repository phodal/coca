package main

import (
	"github.com/phodal/coca/trial/cmd"
	"os"
)

func main() {
	output := os.Stdout
	rootCmd := cmd.NewTrialRootCmd(output)
	_ = rootCmd.Execute()
}
