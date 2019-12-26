package unused

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestRemoveUnusedImportApp_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)


	codePath := "../../../../_fixtures/refactor/unused"
	app := NewRemoveUnusedImportApp(codePath)
	app.Analysis()


	g.Expect(true).To(Equal(true))
}
