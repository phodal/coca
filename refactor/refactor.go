package main

import (
	. "./base"
	"os"
)

func main() {
	//cmd.Execute()
	path := "examples/unused-import"

	if len(os.Args) > 1 {
		path = os.Args[1:][0]
	}

	callApp := new(JavaRefactorApp)
	callApp.AnalysisPath(path)
}
