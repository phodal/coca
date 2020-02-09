package tbs

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/application/analysis/javaapp"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"path/filepath"
	"testing"
)

func TestTbsApp_EmptyTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/usecases/EmptyTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(result[0].FileName).To(Equal(filepath.FromSlash("../../../_fixtures/tbs/usecases/EmptyTest.java")))
	g.Expect(result[0].Line).To(Equal(8))
	g.Expect(result[0].Type).To(Equal("EmptyTest"))
}

func TestTbsApp_IgnoreTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/usecases/IgnoreTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(len(result)).To(Equal(1))
	g.Expect(result[0].Line).To(Equal(0))
	g.Expect(result[0].Type).To(Equal("IgnoreTest"))
}

func TestTbsApp_RedundantPrintTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/usecases/RedundantPrintTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(result[0].Line).To(Equal(9))
	g.Expect(result[0].Type).To(Equal("RedundantPrintTest"))
}

func TestTbsApp_SleepyTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/usecases/SleepyTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(result[0].Line).To(Equal(8))
	g.Expect(result[0].Type).To(Equal("SleepyTest"))
}

func TestTbsApp_DuplicateAssertTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/usecases/DuplicateAssertTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(len(result)).To(Equal(1))
	g.Expect(result[0].Line).To(Equal(9))
	g.Expect(result[0].Type).To(Equal("DuplicateAssertTest"))
}

func TestTbsApp_UnknownTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/usecases/UnknownTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(result[0].Type).To(Equal("EmptyTest"))
	g.Expect(result[0].Line).To(Equal(7))
	g.Expect(result[1].Type).To(Equal("UnknownTest"))
}

func TestTbsApp_RedundantAssertionTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/usecases/RedundantAssertionTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(len(result)).To(Equal(1))
	g.Expect(result[0].Type).To(Equal("RedundantAssertionTest"))
}

func TestTbsApp_CreatorNotUnknownTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/regression/CreatorNotUnknownTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(len(result)).To(Equal(0))
}

func TestTbsApp_CallAssertInClassTests(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/regression/CallAssertInClassTests.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(len(result)).To(Equal(0))
}

func TestTbsApp_ShouldReturnEmptyForFunc(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/regression/EnvironmentSystemIntegrationTests.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(len(result)).To(Equal(0))
}

func TestTbsApp_ShouldReturnMultipleResult(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/regression/I18NTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(len(result)).To(Equal(4))
}

func buildTbsResult(codePath string) []TestBadSmell {
	identifiersMap, classNodes := BuildTestAnalysisResultsByPath(codePath)

	app := NewTbsApp()
	result := app.AnalysisPath(classNodes, identifiersMap)
	return result
}

func BuildTestAnalysisResultsByPath(codePath string) (map[string]core_domain.CodeDataStruct, []core_domain.CodeDataStruct) {
	files := cocafile.GetJavaTestFiles(codePath)

	identifiers := cmd_util.LoadTestIdentify(files)
	identifiersMap := core_domain.BuildIdentifierMap(identifiers)

	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.NodeName)
	}

	analysisApp := javaapp.NewJavaFullApp()
	classNodes := analysisApp.AnalysisFiles(identifiers, files)
	return identifiersMap, classNodes
}
