package tbs

import (
	"github.com/phodal/coca/core/models"
)

type TbsApp struct {
}

func NewTbsApp() *TbsApp {
	return &TbsApp{}
}

type TestBadSmell struct {
	FileName    string
	Type        string
	Description string
	Line        int
}

func (a TbsApp) AnalysisPath(deps []models.JClassNode, identifiersMap map[string]models.JIdentifier) []TestBadSmell {
	var results []TestBadSmell = nil

	for _, clz := range deps {
		// TODO refactoring identify & annotation
		for _, method := range clz.Methods {
			checkIgnoreTest(clz.Path, method, &results)
			checkEmptyTest(clz.Path, method, &results)
			checkRedundantPrintTest(clz.Path, method, &results)
		}
	}

	return results
}

func checkRedundantPrintTest(path string, method models.JMethod, results *[]TestBadSmell) {
	for _, method := range method.MethodCalls {
		if method.Class == "System.out" && (method.MethodName == "println" || method.MethodName == "printf" || method.MethodName == "print") {
			tbs := *&TestBadSmell{
				FileName:    path,
				Type:        "RedundantPrintTest",
				Description: "",
				Line:        0,
			}

			*results = append(*results, tbs)
		}
	}
}

func checkEmptyTest(path string, method models.JMethod, results *[]TestBadSmell) {
	for _, annotation := range method.Annotations {
		if annotation.QualifiedName == "Test" {
			if len(method.MethodCalls) <= 1 {
				tbs := *&TestBadSmell{
					FileName:    path,
					Type:        "EmptyTest",
					Description: "",
					Line:        0,
				}

				*results = append(*results, tbs)
			}
		}
	}

}

func checkIgnoreTest(clzPath string, method models.JMethod, results *[]TestBadSmell) {
	for _, annotation := range method.Annotations {
		if annotation.QualifiedName == "Ignore" {
			tbs := *&TestBadSmell{
				FileName:    clzPath,
				Type:        "IgnoreTest",
				Description: "",
				Line:        0,
			}

			*results = append(*results, tbs)
		}
	}
}
