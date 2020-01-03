package main

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/core/context/concept"
	"github.com/phodal/coca/core/domain"
	"syscall/js"
	"wasm/wadapter"
)

func registerCallbacks() {
	js.Global().Set("compileCode", js.FuncOf(compileCodeCallback))
	js.Global().Set("analysisBadsmell", js.FuncOf(badSmellCallback))
	js.Global().Set("analysisConcept", js.FuncOf(conceptCallback))
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

func badSmellCallback(value js.Value, args []js.Value) interface{} {
	callback := args[len(args)-1:][0]
	message := args[0].String()

	results := new(wadapter.WBadSmell).Analysis(message)

	identModel, _ := json.Marshal(results)
	callback.Invoke(js.Null(), string(identModel))
	return nil
}

func conceptCallback(value js.Value, args []js.Value) interface{} {
	callback := args[len(args)-1:][0]
	message := args[0].String()

	var parsedDeps []domain.JClassNode
	_ = json.Unmarshal([]byte(message), &parsedDeps)

	fmt.Println(parsedDeps)

	wordCounts := concept.NewConceptAnalyser().Analysis(&parsedDeps)

	identModel, _ := json.Marshal(wordCounts)
	callback.Invoke(js.Null(), string(identModel))
	return nil
}
