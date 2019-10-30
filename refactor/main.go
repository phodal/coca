package main

import (
	. "./rename"
	. "./unused"
)

func main() {
	rename := RenameMethodApp("configs/move.coca", "examples/move-demo/src/")
	rename.Analysis()

	app2 := NewRemoveUnusedImportApp("configs/move.coca", "examples/move-demo/src/")
	app2.Analysis()
}