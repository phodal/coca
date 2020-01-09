package moveclass

import (
	. "github.com/onsi/gomega"
	"os"
	"path/filepath"
	"testing"
)

func TestMoveClassApp(t *testing.T) {
	g := NewGomegaWithT(t)

	config := filepath.FromSlash("../../../../_fixtures/refactor/move.config")
	path := filepath.FromSlash("../../../../_fixtures/refactor")

	absPath, _ := filepath.Abs(path)
	app := NewMoveClassApp(config, filepath.FromSlash(absPath))
	app.Analysis()
	app.Refactoring()

	stat, _ := os.Stat(filepath.FromSlash(absPath + "/move/b/ImportForB.java"))
	g.Expect(stat.Name()).To(Equal("ImportForB.java"))

	g.Expect(true).To(Equal(true))
}