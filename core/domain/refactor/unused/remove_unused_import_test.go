package unused

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestRemoveUnusedImportApp_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)


	codePath := "../../../../_fixtures/refactor/unused"
	app := NewRemoveUnusedImportApp(codePath)
	results := app.Analysis()

	g.Expect(len(results)).To(Equal(1))

	errorLines := BuildErrorLines(results[0])
	g.Expect(errorLines).To(Equal([]int{3,4,5}))
}
