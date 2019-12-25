package identifier

import (
	"fmt"
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

	fmt.Println(identifiers)
	g.Expect(len(identifiers)).To(Equal(1))
	g.Expect(identifiers[0].ClassName).To(Equal("Overload"))
	g.Expect(len(identifiers[0].Methods)).To(Equal(3))
}

func TestPolymorphism_Constructor(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := new(JavaIdentifierApp)
	identifiers := identApp.AnalysisPath("../../../_fixtures/suggest/factory")

	fmt.Println(identifiers)
	g.Expect(len(identifiers)).To(Equal(2))
	g.Expect(identifiers[0].ClassName).To(Equal("Insect"))
	g.Expect(identifiers[1].ClassName).To(Equal("Bee"))
}