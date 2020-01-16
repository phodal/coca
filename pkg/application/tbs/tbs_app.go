package tbs

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/domain/jdomain"
	"github.com/phodal/coca/pkg/infrastructure/constants"
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

func (a TbsApp) AnalysisPath(deps []jdomain.JClassNode, identifiersMap map[string]jdomain.JIdentifier) []TestBadSmell {
	var results []TestBadSmell = nil
	callMethodMap := jdomain.BuildCallMethodMap(deps)
	for _, clz := range deps {
		for _, method := range clz.Methods {
			if !method.IsJunitTest() {
				continue
			}

			currentMethodCalls := updateMethodCallsForSelfCall(method, clz, callMethodMap)
			method.MethodCalls = currentMethodCalls

			var testType = ""
			for _, annotation := range method.Annotations {
				checkIgnoreTest(clz.FilePath, annotation, &results, &testType)
				checkEmptyTest(clz.FilePath, annotation, &results, method, &testType)
			}

			var methodCallMap = make(map[string][]jdomain.JMethodCall)
			var hasAssert = false
			for index, methodCall := range currentMethodCalls {
				if methodCall.MethodName == "" {
					if index == len(currentMethodCalls)-1 {
						checkAssert(hasAssert, clz, method, &results, &testType)
					}
					continue
				}

				methodCallMap[methodCall.BuildFullMethodName()] = append(methodCallMap[methodCall.BuildFullMethodName()], methodCall)

				checkRedundantPrintTest(clz.FilePath, methodCall, &results, &testType)
				checkSleepyTest(clz.FilePath, methodCall, method, &results, &testType)
				checkRedundantAssertionTest(clz.FilePath, methodCall, method, &results, &testType)

				if methodCall.HasAssertion() {
					hasAssert = true
				}

				if index == len(currentMethodCalls)-1 {
					checkAssert(hasAssert, clz, method, &results, &testType)
				}
			}

			checkDuplicateAssertTest(clz, &results, methodCallMap, method, &testType)
		}
	}

	return results
}

func checkAssert(hasAssert bool, clz jdomain.JClassNode, method jdomain.JMethod, results *[]TestBadSmell, testType *string) {
	if !hasAssert {
		*testType = "UnknownTest"
		tbs := TestBadSmell{
			FileName:    clz.FilePath,
			Type:        *testType,
			Description: "",
			Line:        method.StartLine,
		}

		*results = append(*results, tbs)

	}
}

func updateMethodCallsForSelfCall(method jdomain.JMethod, clz jdomain.JClassNode, callMethodMap map[string]jdomain.JMethod) []jdomain.JMethodCall {
	currentMethodCalls := method.MethodCalls
	for _, methodCall := range currentMethodCalls {
		if methodCall.Class == clz.Class {
			jMethod := callMethodMap[methodCall.BuildFullMethodName()]
			if jMethod.Name != "" {
				currentMethodCalls = append(currentMethodCalls, jMethod.MethodCalls...)
			}
		}
	}
	return currentMethodCalls
}

func checkRedundantAssertionTest(path string, call jdomain.JMethodCall, method jdomain.JMethod, results *[]TestBadSmell, testType *string) {
	TWO_PARAMETERS := 2
	if len(call.Parameters) == TWO_PARAMETERS {
		if call.Parameters[0] == call.Parameters[1] {
			*testType = "RedundantAssertionTest"
			tbs := TestBadSmell{
				FileName:    path,
				Type:        *testType,
				Description: "",
				Line:        method.StartLine,
			}

			*results = append(*results, tbs)
		}
	}
}

func checkDuplicateAssertTest(clz jdomain.JClassNode, results *[]TestBadSmell, methodCallMap map[string][]jdomain.JMethodCall, method jdomain.JMethod, testType *string) {
	var isDuplicateAssert = false
	for _, methodCall := range methodCallMap {
		if len(methodCall) >= constants.DuplicatedAssertionLimitLength {
			if methodCall[len(methodCall)-1].HasAssertion() {
				isDuplicateAssert = true
			}
		}
	}

	if isDuplicateAssert {
		*testType = "DuplicateAssertTest"
		tbs := TestBadSmell{
			FileName:    clz.FilePath,
			Type:        *testType,
			Description: "",
			Line:        method.StartLine,
		}

		*results = append(*results, tbs)
	}
}

func checkSleepyTest(path string, method jdomain.JMethodCall, jMethod jdomain.JMethod, results *[]TestBadSmell, testType *string) {
	if method.IsThreadSleep() {
		*testType = "SleepyTest"
		tbs := TestBadSmell{
			FileName:    path,
			Type:        *testType,
			Description: "",
			Line:        method.StartLine,
		}

		*results = append(*results, tbs)
	}
}

func checkRedundantPrintTest(path string, mCall jdomain.JMethodCall, results *[]TestBadSmell, testType *string) {
	if mCall.IsSystemOutput() {
		*testType = "RedundantPrintTest"
		tbs := TestBadSmell{
			FileName:    path,
			Type:        *testType,
			Description: "",
			Line:        mCall.StartLine,
		}

		*results = append(*results, tbs)
	}
}

func checkEmptyTest(path string, annotation core_domain.CodeAnnotation, results *[]TestBadSmell, method jdomain.JMethod, testType *string) {
	if annotation.IsTest() {
		if len(method.MethodCalls) <= 1 {
			*testType = "EmptyTest"
			tbs := TestBadSmell{
				FileName:    path,
				Type:        *testType,
				Description: "",
				Line:        method.StartLine,
			}

			*results = append(*results, tbs)
		}
	}
}

func checkIgnoreTest(clzPath string, annotation core_domain.CodeAnnotation, results *[]TestBadSmell, testType *string) {
	if annotation.IsIgnoreTest() {
		*testType = "IgnoreTest"
		tbs := TestBadSmell{
			FileName:    clzPath,
			Type:        *testType,
			Description: "",
			Line:        0,
		}

		*results = append(*results, tbs)
	}
}
