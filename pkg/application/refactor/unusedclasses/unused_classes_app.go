package unusedclasses

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
	"sort"
	"strings"
)

var analysisPackage = ""

func Refactoring(parsedDeps []core_domain.CodeDataStruct) []string {
	sourceClasses := make(map[string]string)
	targetClasses := make(map[string]string)

	for _, node := range parsedDeps {
		if strings.Contains(node.Package, analysisPackage) {
			className := node.Package + "." + node.NodeName
			sourceClasses[className] = className
		}

		for _, method := range node.Functions {
			for _, methodCall := range method.FunctionCalls {
				if strings.Contains(methodCall.Package, analysisPackage) {
					className := methodCall.Package + "." + methodCall.NodeName
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

	sort.Strings(excludePackage)
	return excludePackage
}
