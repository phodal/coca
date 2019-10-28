package main

import (
	//"./cmd"
	"os"

	. "./app"
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
