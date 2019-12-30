package api

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/ast"
	"github.com/phodal/coca/core/ast/full"
	"github.com/phodal/coca/core/ast/identifier"
	"path/filepath"
	"testing"
)

func TestJavaCallApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/call"
	codePath = filepath.FromSlash(codePath)

	identifierApp := new(identifier.JavaIdentifierApp)
	identifiers := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := full.NewJavaFullApp()
	callNodes := callApp.AnalysisPath(codePath, classes, identifiers)

	identifiersMap := ast.BuildIdentifierMap(identifiers)
	diMap := ast.BuildDIMap(identifiers, identifiersMap)

	app := new(JavaApiApp)
	restApis := app.AnalysisPath(codePath, callNodes, identifiersMap, diMap)


	g.Expect(len(restApis)).To(Equal(4))
	g.Expect(restApis[0].HttpMethod).To(Equal("POST"))
	g.Expect(restApis[0].Uri).To(Equal("/books"))
}