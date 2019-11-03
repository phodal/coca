package main

import (
	"encoding/json"
	"fmt"
	. "github.com/phodal/coca/adapter/call"
	. "github.com/phodal/coca/adapter/identifier"
)

func main() {
	identifierApp := new(JavaIdentifierApp)
	iNodes := identifierApp.AnalysisPath("examples/lambda/LambdaExample.java")

	var classes []string = nil

	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.Name)
	}

	callApp := new(JavaCallApp)
	callNodes := callApp.AnalysisPath("examples/lambda/LambdaExample.java", classes)

	cModel, _ := json.MarshalIndent(callNodes, "", "\t")

	fmt.Println(string(cModel))
}
