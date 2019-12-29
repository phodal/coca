package identifier

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestJavaIdentifierApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewJavaIdentifierApp()
	identifiers := identApp.AnalysisPath("../../../_fixtures/call")

	g.Expect(len(identifiers)).To(Equal(1))
	g.Expect(identifiers[0].ClassName).To(Equal("BookController"))
	g.Expect(identifiers[0].Methods[0].Name).To(Equal("BookController"))
	g.Expect(identifiers[0].Methods[1].Name).To(Equal("createBook"))

	g.Expect(identifiers[0].Annotations[0].QualifiedName).To(Equal("RestController"))
}

func TestPolymorphism_Method(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewJavaIdentifierApp()
	identifiers := identApp.AnalysisPath("../../../_fixtures/suggest/polymorphism")

	g.Expect(len(identifiers)).To(Equal(1))
	g.Expect(identifiers[0].ClassName).To(Equal("Overload"))
	g.Expect(len(identifiers[0].Methods)).To(Equal(3))
}

func TestPolymorphism_Constructor(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewJavaIdentifierApp()
	identifiers := identApp.AnalysisPath("../../../_fixtures/suggest/factory")

	g.Expect(len(identifiers)).To(Equal(2))
	g.Expect(identifiers[0].ClassName).To(Equal("Insect"))
	g.Expect(identifiers[1].ClassName).To(Equal("Bee"))
}

func TestAddReturnNull(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewJavaIdentifierApp()
	identifiers := identApp.AnalysisPath("../../../_fixtures/evaluate/null")

	g.Expect(identifiers[0].Methods[1].IsReturnNull).To(Equal(true))
	g.Expect(identifiers[0].Methods[2].IsReturnNull).To(Equal(true))
}

func TestStaticMethod(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewJavaIdentifierApp()
	identifiers := identApp.AnalysisPath("../../../_fixtures/suggest/static")

	g.Expect(identifiers[0].Methods[0].Modifiers[0]).To(Equal("public"))
	g.Expect(identifiers[0].Methods[0].Modifiers[1]).To(Equal("static"))
}

func TestModifierLength(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewJavaIdentifierApp()
	identifiers := identApp.AnalysisPath("../../../_fixtures/bs/ComplexIf.java")

	g.Expect(len(identifiers[0].Methods[0].Modifiers)).To(Equal(1))
}