package main

import (
	"github.com/phodal/coca/cmd"
	"os"
)

func main() {
	//defer profile.Start().Stop()
	output := os.Stdout
	rootCmd := cmd.NewRootCmd(output)
	_ = rootCmd.Execute()
}
