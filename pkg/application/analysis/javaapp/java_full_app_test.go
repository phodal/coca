package javaapp

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"path/filepath"
	"testing"
)

func TestJavaCallApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/call"
	codePath = filepath.FromSlash(codePath)

	identifierApp := new(JavaIdentifierApp)
	iNodes := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.NodeName)
	}

	callApp := NewJavaFullApp()
	callNodes := callApp.AnalysisPath(codePath, iNodes)

	g.Expect(len(callNodes)).To(Equal(1))
}

func TestJavaCallListener_EnterConstructorDeclaration(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/suggest/factory"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)
	g.Expect(len(callNodes[0].Functions)).To(Equal(3))
}

func getCallNodes(codePath string) []core_domain.CodeDataStruct {
	identifierApp := NewJavaIdentifierApp()
	iNodes := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.NodeName)
	}

	callApp := NewJavaFullApp()

	callNodes := callApp.AnalysisPath(codePath, iNodes)
	return callNodes
}

func TestLambda_Express(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/grammar/java/lambda"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[1].Functions {
		methodMap[c.Name] = c
	}

	g.Expect(methodMap["save"].FunctionCalls[0].FunctionName).To(Equal("of"))
	g.Expect(methodMap["findById"].FunctionCalls[3].FunctionName).To(Equal("toDomainModel"))
}

func TestInterface(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/grammar/java/interface"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)

	for _, c := range callNodes[0].Functions {
		methodMap[c.Name] = c
	}

	g.Expect(len(callNodes[0].Functions)).To(Equal(6))
	g.Expect(methodMap["count"].Name).To(Equal("count"))
}

func TestAnnotation(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/grammar/java/regression/HostDependentDownloadableContribution.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[0].Functions {
		methodMap[c.Name] = c
	}

	g.Expect(len(callNodes[0].Annotations)).To(Equal(0))
	g.Expect(methodMap["macOsXPositiveTest"].Name).To(Equal("macOsXPositiveTest"))
}

func Test_ShouldHaveOnlyOneAnnotation(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/tbs/regression/CallAssertInClassTests.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[0].Functions {
		methodMap[c.Name] = c
	}

	g.Expect(len(methodMap["supportsEventType"].Annotations)).To(Equal(1))
	g.Expect(len(methodMap["genericListenerRawTypeTypeErasure"].Annotations)).To(Equal(1))
}

func Test_ShouldHaveOnlyOneAnnotationWithMultipleSame(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/tbs/regression/EnvironmentSystemIntegrationTests.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[0].Functions {
		methodMap[c.Name] = c
	}

	g.Expect(len(methodMap["setUp"].Annotations)).To(Equal(1))
	g.Expect(len(methodMap["annotationConfigApplicationContext_withProfileExpressionMatchOr"].Annotations)).To(Equal(1))
	g.Expect(len(methodMap["annotationConfigApplicationContext_withProfileExpressionMatchAnd"].Annotations)).To(Equal(1))
	g.Expect(len(methodMap["annotationConfigApplicationContext_withProfileExpressionNoMatchAnd"].Annotations)).To(Equal(1))
	g.Expect(len(methodMap["annotationConfigApplicationContext_withProfileExpressionNoMatchNone"].Annotations)).To(Equal(1))
}

func Test_CreatorAnnotation(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/grammar/java/regression/HostDependentDownloadableContribution.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[0].Functions {
		methodMap[c.Name] = c
	}

	g.Expect(len(methodMap["macOsXPositiveTest"].Annotations)).To(Equal(0))
}

func Test_ShouldGetMethodCreators(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/grammar/java/regression/HostDependentDownloadableContribution.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[0].Functions {
		methodMap[c.Name] = c
	}

	g.Expect(len(methodMap["macOsXPositiveTest"].InnerStructures)).To(Equal(2))
}

func Test_ShouldNotGetCreators(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/tbs/usecases/RedundantAssertionTest.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[0].Functions {
		methodMap[c.Name] = c
	}

	g.Expect(len(methodMap["testTrue"].InnerStructures)).To(Equal(0))
}

func Test_ShouldGetMethodCallParameters(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/tbs/usecases/RedundantAssertionTest.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodCallMap := make(map[string]core_domain.CodeCall)
	for _, method := range callNodes[0].Functions {
		for _, call := range method.FunctionCalls {
			methodCallMap[call.FunctionName] = call
		}
	}

	g.Expect(methodCallMap["assertEquals"].Parameters[0].TypeValue).To(Equal("true"))
	g.Expect(methodCallMap["assertEquals"].Parameters[1].TypeValue).To(Equal("true"))
}

func Test_BuilderCallSplitIssue(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/grammar/java/regression/BuilderCallSplitIssue.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodCallMap := make(map[string]core_domain.CodeCall)
	for _, method := range callNodes[0].Functions {
		for _, call := range method.FunctionCalls {
			methodCallMap[call.FunctionName] = call
		}
	}

	g.Expect(methodCallMap["assertThat"].NodeName).To(Equal(""))
	g.Expect(methodCallMap["isFalse"].NodeName).To(Equal("assertThat"))
}

func Test_InnerClass(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/grammar/java/regression/InnerClass.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	g.Expect(len(callNodes)).To(Equal(1))
	g.Expect(callNodes[0].NodeName).To(Equal("Outer"))
	g.Expect(callNodes[0].InnerStructures[0].NodeName).To(Equal("Inner"))
}

func Test_DoubleClass(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/grammar/java/regression/DoubleClass.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	g.Expect(len(callNodes)).To(Equal(2))
	g.Expect(callNodes[0].NodeName).To(Equal("ClassOne"))
	g.Expect(callNodes[1].NodeName).To(Equal("ClassTwo"))
}

func Test_InnerJavaClass(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/grammar/java/regression/JavaInner.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	g.Expect(len(callNodes)).To(Equal(1))
	g.Expect(callNodes[0].NodeName).To(Equal("Outer"))
	g.Expect(callNodes[0].InnerStructures[0].NodeName).To(Equal("Inner"))
}

func Test_FormalParameterCall(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/grammar/java/regression/FormalParameterCall.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	calls := callNodes[0].Functions[0].FunctionCalls
	g.Expect(len(calls)).To(Equal(1))
	g.Expect(calls[0].FunctionName).To(Equal("getIsbn"))
	g.Expect(calls[0].Package).To(Equal("hello"))
	g.Expect(calls[0].NodeName).To(Equal("CreateBookCommand"))
}

func Test_NormalChainCall(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/grammar/java/regression/NormalChainCall.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	calls := callNodes[0].Functions[0].FunctionCalls
	g.Expect(len(calls)).To(Equal(2))
	g.Expect(calls[1].NodeName).To(Equal("UriComponentsBuilder"))
	g.Expect(calls[1].FunctionName).To(Equal("buildAndExpand"))
}
