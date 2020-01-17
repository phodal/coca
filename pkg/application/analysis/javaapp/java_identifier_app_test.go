package javaapp

import (
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func TestJavaIdentifierApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewJavaIdentifierApp()
	identifiers := identApp.AnalysisPath(filepath.FromSlash("../../../../_fixtures/call"))

	g.Expect(len(identifiers)).To(Equal(1))
	g.Expect(identifiers[0].NodeName).To(Equal("BookController"))
	g.Expect(identifiers[0].Functions[0].Name).To(Equal("BookController"))
	g.Expect(identifiers[0].Functions[1].Name).To(Equal("createBook"))

	g.Expect(identifiers[0].Annotations[0].Name).To(Equal("RestController"))
}

func TestPolymorphism_Method(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewJavaIdentifierApp()
	identifiers := identApp.AnalysisPath(filepath.FromSlash("../../../../_fixtures/suggest/polymorphism"))

	g.Expect(len(identifiers)).To(Equal(1))
	g.Expect(identifiers[0].NodeName).To(Equal("Overload"))
	g.Expect(len(identifiers[0].Functions)).To(Equal(3))
}

func TestPolymorphism_Constructor(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewJavaIdentifierApp()
	identifiers := identApp.AnalysisPath(filepath.FromSlash("../../../../_fixtures/suggest/factory"))

	g.Expect(len(identifiers)).To(Equal(2))
	g.Expect(identifiers[0].NodeName).To(Equal("Insect"))
	g.Expect(identifiers[1].NodeName).To(Equal("Bee"))
}

func TestAddReturnNull(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewJavaIdentifierApp()
	identifiers := identApp.AnalysisPath(filepath.FromSlash("../../../../_fixtures/evaluate/null"))

	var returNullCount = 0
	for _, method := range identifiers[0].Functions {
		if method.IsReturnNull {
			returNullCount++
		}
	}

	g.Expect(identifiers[0].Functions[1].IsReturnNull).To(Equal(true))
	g.Expect(identifiers[0].Functions[2].IsReturnNull).To(Equal(true))
	g.Expect(returNullCount).To(Equal(2))
}

func TestStaticMethod(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewJavaIdentifierApp()
	identifiers := identApp.AnalysisPath(filepath.FromSlash("../../../../_fixtures/suggest/static"))

	g.Expect(identifiers[0].Functions[0].Modifiers[0]).To(Equal("public"))
	g.Expect(identifiers[0].Functions[0].Modifiers[1]).To(Equal("static"))
}

func TestModifierLength(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewJavaIdentifierApp()
	identifiers := identApp.AnalysisPath(filepath.FromSlash("../../../../_fixtures/bs/ComplexIf.java"))

	g.Expect(len(identifiers[0].Functions[0].Modifiers)).To(Equal(1))
}
