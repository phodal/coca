package call

import (
	"github.com/phodal/coca/core/adapter/identifier"
	. "github.com/onsi/gomega"
	"testing"
)

func TestJavaCallApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/call"
	identifierApp := new(identifier.JavaIdentifierApp)
	iNodes := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := new(JavaCallApp)

	callNodes := callApp.AnalysisPath(codePath, classes, iNodes)

	g.Expect(len(callNodes)).To(Equal(1))
}