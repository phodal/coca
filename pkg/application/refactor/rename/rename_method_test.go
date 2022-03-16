package unused

import (
	"github.com/modernizing/coca/cmd/cmd_util"
	"github.com/modernizing/coca/cocatest/testhelper"
	"github.com/modernizing/coca/pkg/application/analysis/javaapp"
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func TestRenameMethodApp(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../../_fixtures/refactor/unused"
	configPath := "../../../../_fixtures/refactor/rename.config"
	codePath = filepath.FromSlash(codePath)
	configPath = filepath.FromSlash(configPath)

	identifierApp := new(javaapp.JavaIdentifierApp)
	identifiers := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.NodeName)
	}

	callApp := javaapp.NewJavaFullApp()
	callNodes := callApp.AnalysisPath(codePath, identifiers)

	configBytes := cmd_util.ReadFile(configPath)
	RenameMethodApp(callNodes).Refactoring(string(configBytes))

	newnodes := callApp.AnalysisPath(codePath, identifiers)
	g.Expect(newnodes[0].Functions[0].Name).To(Equal("demo"))

	testhelper.ResetGitDir(codePath)
}
