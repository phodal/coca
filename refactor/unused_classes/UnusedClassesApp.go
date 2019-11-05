package main

import (
	"encoding/json"
	"fmt"
	. "github.com/phodal/coca/adapter/models"
	. "github.com/phodal/coca/utils"
	"sort"
	"strings"
)

var parsedDeps []JClassNode

func main()  {
	var analysisPackage = ""
	file := ReadFile("deps.json")
	if file == nil {
		return
	}

	_ = json.Unmarshal(file, &parsedDeps)
	sourceClasses := make(map[string]string)
	targetlasses := make(map[string]string)

	for _, node := range parsedDeps {
		if strings.Contains(node.Package, analysisPackage) {
			className := node.Package + "." + node.Class
			sourceClasses[className] = className
		}

		for _, methodCall := range node.MethodCalls {
			if strings.Contains(methodCall.Package, analysisPackage) {
				className := methodCall.Package + "." + methodCall.Class
				targetlasses[className] = className
			}
		}
	}

	var excludePackage []string = nil
	for _, clz := range sourceClasses {
		if targetlasses[clz] != clz {
			excludePackage = append(excludePackage, clz)
		}
	}

	sort.Sort(sort.StringSlice(excludePackage))
	for _, res := range excludePackage {
		fmt.Println(res)
	}
}
