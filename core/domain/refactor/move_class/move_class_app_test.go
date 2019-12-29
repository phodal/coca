package move_class

import (
	. "github.com/onsi/gomega"
	"os"
	"path/filepath"
	"testing"
)

func TestNewMoveClassApp(t *testing.T) {
	g := NewGomegaWithT(t)

	config := "../../../../_fixtures/refactor/move.config"
	path := "../../../../_fixtures/refactor/"
	abs_path, _ := filepath.Abs(path)
	app := NewMoveClassApp(config, abs_path + "/")
	app.Analysis()

	stat, _ := os.Stat(path + "/move/b/ImportForB.java")
	g.Expect(stat.Name()).To(Equal("ImportForB.java"))
}