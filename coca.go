package main

import (
	"github.com/phodal/coca/cmd"
)

func main() {
	//defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook)
	cmd.Execute()
}
