package move_class

import (
	"fmt"
	. "github.com/onsi/gomega"
	"os"
	"os/exec"
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

	// debug
	cmd := exec.Command("tree", absPath)
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))

	stat, _ := os.Stat(absPath + "/move/b/ImportForB.java")
	g.Expect(stat.Name()).To(Equal("ImportForB.java"))
}