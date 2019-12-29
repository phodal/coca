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
			for _, annotation := range method.Annotations {
				checkIgnoreTest(clz.Path, annotation, &results)
				checkEmptyTest(clz.Path, annotation, &results, method)
			}

			for _, methodCall := range method.MethodCalls {
				checkRedundantPrintTest(clz.Path, methodCall, &results)
				checkUnknownTest(clz.Path, methodCall, &results)
			}
		}
	}

	return results
}

func checkUnknownTest(path string, method models.JMethodCall, results *[]TestBadSmell) {
	if method.MethodName == "sleep" && method.Class == "Thread" {
		tbs := *&TestBadSmell{
			FileName:    path,
			Type:        "SleepyTest",
			Description: "",
			Line:        0,
		}

		*results = append(*results, tbs)
	}
}

func checkRedundantPrintTest(path string, mCall models.JMethodCall, results *[]TestBadSmell) {
	if mCall.Class == "System.out" && (mCall.MethodName == "println" || mCall.MethodName == "printf" || mCall.MethodName == "print") {
		tbs := *&TestBadSmell{
			FileName:    path,
			Type:        "RedundantPrintTest",
			Description: "",
			Line:        0,
		}

		*results = append(*results, tbs)
	}
}

func checkEmptyTest(path string, annotation models.Annotation, results *[]TestBadSmell, method models.JMethod) {
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

func checkIgnoreTest(clzPath string, annotation models.Annotation, results *[]TestBadSmell) {
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
