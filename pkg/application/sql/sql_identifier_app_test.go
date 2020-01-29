package sql

import (
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func TestJavaIdentifierApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	identApp := NewSqlIdentifierApp()
	results := identApp.AnalysisPath(filepath.FromSlash("../../../_fixtures/grammar/sql"))

	g.Expect(len(results)).To(Equal(1))
}
