package unused_classes

import (
	. "github.com/phodal/coca/core/domain"
	"sort"
	"strings"
)

var parsedDeps []JClassNode
var analysisPackage = ""

func Refactoring(parsedDeps []JClassNode) []string {
	sourceClasses := make(map[string]string)
	targetClasses := make(map[string]string)

	for _, node := range parsedDeps {
		if strings.Contains(node.Package, analysisPackage) {
			className := node.Package + "." + node.Class
			sourceClasses[className] = className
		}

		for _, method := range node.Methods {
			for _, methodCall := range method.MethodCalls {
				if strings.Contains(methodCall.Package, analysisPackage) {
					className := methodCall.Package + "." + methodCall.Class
					targetClasses[className] = className
				}
			}
		}
	}

	var excludePackage []string = nil
	for _, clz := range sourceClasses {
		if targetClasses[clz] != clz {
			excludePackage = append(excludePackage, clz)
		}
	}

	sort.Sort(sort.StringSlice(excludePackage))
	return excludePackage
}
