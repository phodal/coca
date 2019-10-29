package main

import (
	"os"

	. "./adapter/call"
	. "./adapter/identifier"
)

func main() {
	//cmd.Execute()
	path := "/Users/fdhuang/test/mall"

	if len(os.Args) > 1 {
		path = os.Args[1:][0]
	}

	callApp := new(JavaCallApp)
	callApp.AnalysisPath(path)

	identifierApp := new(JavaIdentifierApp)
	identifierApp.AnalysisPath(path)
}
