package main

import (
	"github.com/pkg/profile"
	"coca/cmd"
)

func main() {
	defer profile.Start().Stop()

	cmd.Execute()
}
