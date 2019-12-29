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
			if !isTest(method) {
				continue
			}

			var testType = ""
			for _, annotation := range method.Annotations {
				checkIgnoreTest(clz.Path, annotation, &results, &testType)
				checkEmptyTest(clz.Path, annotation, &results, method, &testType)
			}

			var methodCallMap = make(map[string][]models.JMethodCall)
			var hasAssert = false
			for index, methodCall := range method.MethodCalls {
				methodCallMap[methodCallName(methodCall)] = append(methodCallMap[methodCallName(methodCall)], methodCall)

				checkRedundantPrintTest(clz.Path, methodCall, &results, &testType)
				checkSleepyTest(clz.Path, methodCall, method, &results, &testType)

				if strings.Contains(methodCall.MethodName, "assert") {
					hasAssert = true
				}

				if index == len(method.MethodCalls)-1 {
					if !hasAssert {
						checkUnknownTest(clz, method, &results, &testType)
					}
				}
			}

			checkDuplicateAssertTest(clz, &results, methodCallMap, &testType)
		}
	}

	return results
}

func isTest(method models.JMethod) bool {
	var isTest = false
	for _, annotation := range method.Annotations {
		if annotation.QualifiedName == "Test" || annotation.QualifiedName == "Ignore" {
			isTest = true
		}
	}
	return isTest
}

func checkUnknownTest(clz models.JClassNode, method models.JMethod, results *[]TestBadSmell, testType *string) {
	*testType = "UnknownTest"
	tbs := *&TestBadSmell{
		FileName:    clz.Path,
		Type:        *testType,
		Description: "",
		Line:        method.StartLine,
	}

	*results = append(*results, tbs)
}

func checkDuplicateAssertTest(clz models.JClassNode, results *[]TestBadSmell, methodCallMap map[string][]models.JMethodCall, testType *string) {
	for _, methodCall := range methodCallMap {
		if len(methodCall) >= 3 {
			if strings.Contains(methodCall[len(methodCall)-1].MethodName, "assert") {

				*testType = "DuplicateAssertTest"
				tbs := *&TestBadSmell{
					FileName:    clz.Path,
					Type:        *testType,
					Description: "",
					Line:        methodCall[len(methodCall)-1].StartLine,
				}

				*results = append(*results, tbs)
			}
		}
	}
}

func methodCallName(methodCall models.JMethodCall) string {
	return methodCall.Package + "." + methodCall.Class + "." + methodCall.MethodName
}

func checkSleepyTest(path string, method models.JMethodCall, jMethod models.JMethod, results *[]TestBadSmell, testType *string) {
	if method.MethodName == "sleep" && method.Class == "Thread" {
		*testType = "SleepyTest"
		tbs := *&TestBadSmell{
			FileName:    path,
			Type:        *testType,
			Description: "",
			Line:        method.StartLine,
		}

		*results = append(*results, tbs)
	}
}

func checkRedundantPrintTest(path string, mCall models.JMethodCall, results *[]TestBadSmell, testType *string) {
	if mCall.Class == "System.out" && (mCall.MethodName == "println" || mCall.MethodName == "printf" || mCall.MethodName == "print") {
		*testType = "RedundantPrintTest"
		tbs := *&TestBadSmell{
			FileName:    path,
			Type:        *testType,
			Description: "",
			Line:        mCall.StartLine,
		}

		*results = append(*results, tbs)
	}
}

func checkEmptyTest(path string, annotation models.Annotation, results *[]TestBadSmell, method models.JMethod, testType *string) {
	if annotation.QualifiedName == "Test" {
		if len(method.MethodCalls) <= 1 {
			*testType = "EmptyTest"
			tbs := *&TestBadSmell{
				FileName:    path,
				Type:        *testType,
				Description: "",
				Line:        method.StartLine,
			}

			*results = append(*results, tbs)
		}
	}
}

func checkIgnoreTest(clzPath string, annotation models.Annotation, results *[]TestBadSmell, testType *string) {
	if annotation.QualifiedName == "Ignore" {
		*testType = "IgnoreTest"
		tbs := *&TestBadSmell{
			FileName:    clzPath,
			Type:        *testType,
			Description: "",
			Line:        0,
		}

		*results = append(*results, tbs)
	}
}
