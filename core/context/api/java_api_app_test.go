package api

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/context/analysis"
	"github.com/phodal/coca/core/domain"
	"path/filepath"
	"testing"
)

func TestJavaCallApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/call"
	codePath = filepath.FromSlash(codePath)

	identifierApp := new(analysis.JavaIdentifierApp)
	identifiers := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := analysis.NewJavaFullApp()
	callNodes := callApp.AnalysisPath(codePath, classes, identifiers)

	identifiersMap := domain.BuildIdentifierMap(identifiers)
	diMap := domain.BuildDIMap(identifiers, identifiersMap)

	app := new(JavaApiApp)
	restApis := app.AnalysisPath(codePath, callNodes, identifiersMap, diMap)


	g.Expect(len(restApis)).To(Equal(4))
	g.Expect(restApis[0].HttpMethod).To(Equal("POST"))
	g.Expect(restApis[0].Uri).To(Equal("/books"))
}