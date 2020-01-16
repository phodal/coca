package analysis

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"path/filepath"
	"testing"
)

func TestJavaCallApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/call"
	codePath = filepath.FromSlash(codePath)

	identifierApp := new(JavaIdentifierApp)
	iNodes := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := NewJavaFullApp()
	callNodes := callApp.AnalysisPath(codePath, classes, iNodes)

	g.Expect(len(callNodes)).To(Equal(1))
}

func TestJavaCallListener_EnterConstructorDeclaration(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/suggest/factory"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)
	g.Expect(len(callNodes[0].Methods)).To(Equal(3))
}

func getCallNodes(codePath string) []core_domain.JClassNode {
	identifierApp := NewJavaIdentifierApp()
	iNodes := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := NewJavaFullApp()

	callNodes := callApp.AnalysisPath(codePath, classes, iNodes)
	return callNodes
}

func TestLambda_Express(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/lambda"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[1].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(methodMap["save"].MethodCalls[0].MethodName).To(Equal("of"))
	g.Expect(methodMap["findById"].MethodCalls[3].MethodName).To(Equal("toDomainModel"))
}

func TestInterface(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/grammar/interface"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)

	for _, c := range callNodes[0].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(len(callNodes[0].Methods)).To(Equal(6))
	g.Expect(methodMap["count"].Name).To(Equal("count"))
}

func TestAnnotation(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/grammar/HostDependentDownloadableContribution.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[0].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(len(callNodes[0].Annotations)).To(Equal(0))
	g.Expect(methodMap["macOsXPositiveTest"].Name).To(Equal("macOsXPositiveTest"))
}

func Test_ShouldHaveOnlyOneAnnotation(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/tbs/regression/CallAssertInClassTests.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[0].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(len(methodMap["supportsEventType"].Annotations)).To(Equal(1))
	g.Expect(len(methodMap["genericListenerRawTypeTypeErasure"].Annotations)).To(Equal(1))
}

func Test_ShouldHaveOnlyOneAnnotationWithMultipleSame(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/tbs/regression/EnvironmentSystemIntegrationTests.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[0].Methods {
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

	codePath := "../../../_fixtures/grammar/HostDependentDownloadableContribution.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[0].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(len(methodMap["macOsXPositiveTest"].Annotations)).To(Equal(0))
}

func Test_ShouldGetMethodCreators(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/grammar/HostDependentDownloadableContribution.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[0].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(len(methodMap["macOsXPositiveTest"].InnerStructures)).To(Equal(2))
}

func Test_ShouldNotGetCreators(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/tbs/code/RedundantAssertionTest.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]core_domain.CodeFunction)
	for _, c := range callNodes[0].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(len(methodMap["testTrue"].InnerStructures)).To(Equal(0))
}

func Test_ShouldGetMethodCallParameters(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/tbs/code/RedundantAssertionTest.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodCallMap := make(map[string]core_domain.CodeCall)
	for _, method := range callNodes[0].Methods {
		for _, call := range method.MethodCalls {
			methodCallMap[call.MethodName] = call
		}
	}

	g.Expect(methodCallMap["assertEquals"].Parameters[0].TypeValue).To(Equal("true"))
	g.Expect(methodCallMap["assertEquals"].Parameters[1].TypeValue).To(Equal("true"))
}

func Test_BuilderCallSplitIssue(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/abug/BuilderCallSplitIssue.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	methodCallMap := make(map[string]core_domain.CodeCall)
	for _, method := range callNodes[0].Methods {
		for _, call := range method.MethodCalls {
			methodCallMap[call.MethodName] = call
		}
	}

	g.Expect(methodCallMap["assertThat"].Class).To(Equal(""))
	g.Expect(methodCallMap["isFalse"].Class).To(Equal("assertThat"))
}

func Test_InnerClass(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/abug/InnerClass.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	g.Expect(len(callNodes)).To(Equal(1))
	g.Expect(callNodes[0].Class).To(Equal("Outer"))
	g.Expect(callNodes[0].InnerClass[0].Class).To(Equal("Inner"))
}

func Test_DoubleClass(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/abug/DoubleClass.java"
	codePath = filepath.FromSlash(codePath)

	callNodes := getCallNodes(codePath)

	g.Expect(len(callNodes)).To(Equal(2))
	g.Expect(callNodes[0].Class).To(Equal("ClassOne"))
	g.Expect(callNodes[1].Class).To(Equal("ClassTwo"))
}
