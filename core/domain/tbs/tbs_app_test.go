package tbs

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/adapter"
	"github.com/phodal/coca/core/adapter/call"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"testing"
)

func TestTbsApp_EmptyTest(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/tbs/code/EmptyTest.java"
	result := buildTbsResult(codePath)

	g.Expect(result[0].Type).To(Equal("EmptyTest"))
}

func TestTbsApp_IgnoreTest(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/tbs/code/IgnoreTest.java"
	result := buildTbsResult(codePath)

	g.Expect(result[0].Type).To(Equal("IgnoreTest"))
}

func buildTbsResult(codePath string) []TestBadSmell {
	files := support.GetJavaTestFiles(codePath)
	var identifiers []models.JIdentifier

	identifiers = adapter.LoadTestIdentify(files)
	identifiersMap := adapter.BuildIdentifierMap(identifiers)

	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	analysisApp := call.NewJavaCallApp()
	classNodes := analysisApp.AnalysisFiles(identifiers, files, classes)

	app := NewTbsApp()
	result := app.AnalysisPath(classNodes, identifiersMap)
	return result
}