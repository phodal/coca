package main

import (
	"os"

	. "./adapter/call"
)

func main() {
	//cmd.Execute()
	path := "/Users/fdhuang/learn/coca/poc/src/main/"

	if len(os.Args) > 1 {
		path = os.Args[1:][0]
	}

	app := new(JavaCallApp)
	app.AnalysisPath(path)
}
