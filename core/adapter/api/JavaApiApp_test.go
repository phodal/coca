package api


import (
	"github.com/phodal/coca/core/adapter"
	"github.com/phodal/coca/core/adapter/call"
	"github.com/phodal/coca/core/adapter/identifier"
	. "github.com/onsi/gomega"
	"testing"
)

func TestJavaCallApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/call"
	identifierApp := new(identifier.JavaIdentifierApp)
	identifiers := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := new(call.JavaCallApp)
	callNodes := callApp.AnalysisPath(codePath, classes, identifiers)

	identifiersMap := adapter.BuildIdentifierMap(identifiers)
	diMap := adapter.BuildDIMap(identifiers, identifiersMap)

	app := new(JavaApiApp)
	restApis := app.AnalysisPath(codePath, callNodes, identifiersMap, diMap)


	g.Expect(len(restApis)).To(Equal(4))
	g.Expect(restApis[0].HttpMethod).To(Equal("POST"))
	g.Expect(restApis[0].Uri).To(Equal("/books"))
}