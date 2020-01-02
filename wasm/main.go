package main

import (
	"encoding/json"
	"syscall/js"
	"wasm/wadapter"
)

func registerCallbacks() {
	js.Global().Set("compileCode", js.FuncOf(compileCodeCallback))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}

func compileCodeCallback(value js.Value, args []js.Value) interface{} {
	callback := args[len(args)-1:][0]
	message := args[0].String()

	results := new(wadapter.WAnalysis).Analysis(message)

	identModel, _ := json.Marshal(results)
	callback.Invoke(js.Null(), string(identModel))
	return nil
}
