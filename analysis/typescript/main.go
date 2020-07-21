package main

import (
	"github.com/phodal/coca/analysis/typescript/app"
	"os"
)

func main() {
	output := os.Stdout
	rootCmd := app.NewRootCmd(output)
	_ = rootCmd.Execute()
}
