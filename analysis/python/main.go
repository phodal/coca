package main

import (
	"github.com/modernizing/coca/analysis/python/app"
	"os"
)

func main() {
	output := os.Stdout
	rootCmd := app.NewRootCmd(output)
	_ = rootCmd.Execute()
}
