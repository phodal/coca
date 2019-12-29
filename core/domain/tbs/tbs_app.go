package tbs

import (
	"github.com/phodal/coca/core/models"
	"strings"
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

			var methodCallMap = make(map[string][]models.JMethodCall)
			var hasAssert = false
			for _, methodCall := range method.MethodCalls {
				checkRedundantPrintTest(clz.Path, methodCall, &results)
				checkSleepyTest(clz.Path, methodCall, &results)
				checkDuplicateAssertTest(methodCall, clz, &results, methodCallMap)

				if strings.Contains(methodCall.MethodName, "assert") {
					hasAssert = true
				}
			}

			checkUnknownTest(clz, &results, hasAssert)
		}
	}

	return results
}

func checkUnknownTest(clz models.JClassNode, results *[]TestBadSmell, hasAssert bool) {
	if !hasAssert {
		tbs := *&TestBadSmell{
			FileName:    clz.Path,
			Type:        "UnknownTest",
			Description: "",
			Line:        0,
		}

		*results = append(*results, tbs)
	}
}

func checkDuplicateAssertTest(methodCall models.JMethodCall, clz models.JClassNode, results *[]TestBadSmell, methodCallMap map[string][]models.JMethodCall) {
	methodCallMap[methodCallName(methodCall)] = append(methodCallMap[methodCallName(methodCall)], methodCall)
	if len(methodCallMap[methodCallName(methodCall)]) >= 3 {
		tbs := *&TestBadSmell{
			FileName:    clz.Path,
			Type:        "DuplicateAssertTest",
			Description: "",
			Line:        0,
		}

		*results = append(*results, tbs)
	}
}

func methodCallName(methodCall models.JMethodCall) string {
	return methodCall.Package + "." + methodCall.Class + "." + methodCall.MethodName
}

func checkSleepyTest(path string, method models.JMethodCall, results *[]TestBadSmell) {
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
