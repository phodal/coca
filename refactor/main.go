package main

import (
	. "./rename"
	. "./unused"
)

func main() {
	rename := NewRemoveMethodApp("configs/move.config", "examples/move-demo/src/")
	rename.Analysis()

	app2 := NewRemoveUnusedImportApp("configs/move.config", "examples/move-demo/src/")
	app2.Analysis()
}