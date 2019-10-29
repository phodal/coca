package main

import (
	. "./unused"
)

func main() {
	app2 := NewRemoveUnusedImportApp("configs/move.config", "examples/move-demo/src/")
	app2.Analysis()
}