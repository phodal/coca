package main

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/core/adapter/call"
	"github.com/phodal/coca/core/adapter/identifier"
)

func main()  {
	path := "_fixtures/abug"

	identifierApp := identifier.NewJavaIdentifierApp()
	iNodes := identifierApp.AnalysisPath(path)

	fmt.Println(iNodes)

	var classes []string = nil

	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := new(call.JavaCallApp)

	callNodes := callApp.AnalysisPath(path, classes, iNodes)
	cModel, _ := json.MarshalIndent(callNodes, "", "\t")

	fmt.Println(string(cModel))
}