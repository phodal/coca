package unused

import (
	. "github.com/onsi/gomega"
	"github.com/modernizing/coca/cocatest/testhelper"
	"path/filepath"
	"testing"
)

func TestRemoveUnusedImportApp_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/refactor/unused"
	codePath = filepath.FromSlash(codePath)
	testhelper.ResetGitDir(codePath)

	deps1, _, _ := testhelper.BuildAnalysisDeps(codePath)
	g.Expect(len(deps1[0].Imports)).To(Equal(3))

	app := NewRemoveUnusedImportApp(codePath)
	results := app.Analysis()

	g.Expect(len(results)).To(Equal(1))

	errorLines := BuildErrorLines(results[0])
	g.Expect(errorLines).To(Equal([]int{3, 4, 5}))

	app.Refactoring(results)

	deps, _, _ := testhelper.BuildAnalysisDeps(codePath)
	g.Expect(len(deps[0].Imports)).To(Equal(0))
	testhelper.ResetGitDir(codePath)
}
