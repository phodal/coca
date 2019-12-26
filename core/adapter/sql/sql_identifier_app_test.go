package sql

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestJavaIdentifierApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewSqlIdentifierApp()
	results := identApp.AnalysisPath("../../../_fixtures/sql")

	g.Expect(len(results)).To(Equal(1))
}
