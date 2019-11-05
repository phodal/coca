package main

import (
	"encoding/json"
	. "github.com/phodal/coca/bs"
	. "github.com/phodal/coca/utils"
)

func main()  {
		bsApp := new(BadSmellApp)
		bsList := bsApp.AnalysisPath("examples/")

		bsModel, _ := json.MarshalIndent(bsList, "", "\t")

		WriteToFile("bs.json", string(bsModel))
}
