package api

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cocatest"
	"github.com/phodal/coca/pkg/domain"
	"testing"
)

func TestJavaCallApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/call"
	callNodes, identifiersMap, identifiers := cocatest.BuildAnalysisDeps(codePath)
	diMap := domain.BuildDIMap(identifiers, identifiersMap)

	app := new(JavaApiApp)
	restApis := app.AnalysisPath(codePath, callNodes, identifiersMap, diMap)


	g.Expect(len(restApis)).To(Equal(4))
	g.Expect(restApis[0].HttpMethod).To(Equal("POST"))
	g.Expect(restApis[0].Uri).To(Equal("/books"))
}