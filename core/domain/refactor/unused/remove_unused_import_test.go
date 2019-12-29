package unused

import (
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func TestRemoveUnusedImportApp_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)


	codePath := "../../../../_fixtures/refactor/unused"
	codePath = filepath.FromSlash(codePath)
	app := NewRemoveUnusedImportApp(codePath)
	results := app.Analysis()

	g.Expect(len(results)).To(Equal(1))

	errorLines := BuildErrorLines(results[0])
	g.Expect(errorLines).To(Equal([]int{3,4,5}))
}
