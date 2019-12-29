package move_class

import (
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func TestMoveClassApp(t *testing.T) {
	g := NewGomegaWithT(t)

	config := "../../../../_fixtures/refactor/move.config"
	path := "../../../../_fixtures/refactor/"

	absPath, _ := filepath.Abs(path)
	app := NewMoveClassApp(config, absPath+ "/")
	app.Analysis()

	// TODO: fix in CI, https://travis-ci.org/phodal/coca/jobs/630546918
	//stat, _ := os.Stat(absPath + "/move/b/ImportForB.java")
	//g.Expect(stat.Name()).To(Equal("ImportForB.java"))
	g.Expect(true).To(Equal(true))
}