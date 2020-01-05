package tbs

import (
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure/constants"
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

func (a TbsApp) AnalysisPath(deps []domain.JClassNode, identifiersMap map[string]domain.JIdentifier) []TestBadSmell {
	var results []TestBadSmell = nil
	callMethodMap := domain.BuildCallMethodMap(deps)
	for _, clz := range deps {
		for _, method := range clz.Methods {
			if !method.IsJunitTest() {
				continue
			}

			currentMethodCalls := updateMethodCallsForSelfCall(method, clz, callMethodMap)
			method.MethodCalls = currentMethodCalls

			var testType = ""
			for _, annotation := range method.Annotations {
				checkIgnoreTest(clz.Path, annotation, &results, &testType)
				checkEmptyTest(clz.Path, annotation, &results, method, &testType)
			}

			var methodCallMap = make(map[string][]domain.JMethodCall)
			var hasAssert = false
			for index, methodCall := range currentMethodCalls {
				if methodCall.MethodName == "" {
					if index == len(currentMethodCalls)-1 {
						checkAssert(hasAssert, clz, method, &results, &testType)
					}
					continue
				}

				methodCallMap[methodCall.BuildFullMethodName()] = append(methodCallMap[methodCall.BuildFullMethodName()], methodCall)

				checkRedundantPrintTest(clz.Path, methodCall, &results, &testType)
				checkSleepyTest(clz.Path, methodCall, method, &results, &testType)
				checkRedundantAssertionTest(clz.Path, methodCall, method, &results, &testType)

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

func checkAssert(hasAssert bool, clz domain.JClassNode, method domain.JMethod, results *[]TestBadSmell, testType *string) {
	if !hasAssert {
		*testType = "UnknownTest"
		tbs := *&TestBadSmell{
			FileName:    clz.Path,
			Type:        *testType,
			Description: "",
			Line:        method.StartLine,
		}

		*results = append(*results, tbs)

	}
}

func updateMethodCallsForSelfCall(method domain.JMethod, clz domain.JClassNode, callMethodMap map[string]domain.JMethod) []domain.JMethodCall {
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

func checkRedundantAssertionTest(path string, call domain.JMethodCall, method domain.JMethod, results *[]TestBadSmell, testType *string) {
	TWO_PARAMETERS := 2
	if len(call.Parameters) == TWO_PARAMETERS {
		if call.Parameters[0] == call.Parameters[1] {
			*testType = "RedundantAssertionTest"
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

func checkDuplicateAssertTest(clz domain.JClassNode, results *[]TestBadSmell, methodCallMap map[string][]domain.JMethodCall, method domain.JMethod, testType *string) {
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
		tbs := *&TestBadSmell{
			FileName:    clz.Path,
			Type:        *testType,
			Description: "",
			Line:        method.StartLine,
		}

		*results = append(*results, tbs)
	}
}

func checkSleepyTest(path string, method domain.JMethodCall, jMethod domain.JMethod, results *[]TestBadSmell, testType *string) {
	if method.IsThreadSleep() {
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

func checkRedundantPrintTest(path string, mCall domain.JMethodCall, results *[]TestBadSmell, testType *string) {
	if mCall.IsSystemOutput() {
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

func checkEmptyTest(path string, annotation domain.Annotation, results *[]TestBadSmell, method domain.JMethod, testType *string) {
	if annotation.IsTest() {
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

func checkIgnoreTest(clzPath string, annotation domain.Annotation, results *[]TestBadSmell, testType *string) {
	if annotation.IsIgnoreTest() {
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
