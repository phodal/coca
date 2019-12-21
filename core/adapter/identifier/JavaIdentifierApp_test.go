package identifier

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestJavaIdentifierApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := new(JavaIdentifierApp)
	identifiers := identApp.AnalysisPath("../../../examples/method-call")

	g.Expect(len(identifiers)).To(Equal(1))
	g.Expect(identifiers[0].ClassName).To(Equal("BlogRepositoryImpl"))
	g.Expect(identifiers[0].Methods[0].Name).To(Equal("save"))
}