package visual

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/modernizing/coca/cmd/cmd_util"
	"github.com/modernizing/coca/pkg/application/analysis/javaapp"
	"github.com/modernizing/coca/pkg/domain/core_domain"
	"path/filepath"
	"testing"
)

func Test_BasicVisualSupport(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []core_domain.CodeDataStruct
	codePath := "../../../_fixtures/evaluate/service_long_parameters.json"
	codePath = filepath.FromSlash(codePath)
	file := cmd_util.ReadFile(codePath)
	_ = json.Unmarshal(file, &parsedDeps)

	ddata := FromDeps(parsedDeps)

	g.Expect(len(ddata.Links)).To(Equal(0))
	g.Expect(len(ddata.Nodes)).To(Equal(3))
}

func Test_LinkedVisual(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/grammar/java/arch/step2-java"
	codePath = filepath.FromSlash(codePath)

	identifierApp := new(javaapp.JavaIdentifierApp)
	iNodes := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.NodeName)
	}

	callApp := javaapp.NewJavaFullApp()
	callNodes := callApp.AnalysisPath(codePath, iNodes)

	ddata := FromDeps(callNodes)

	g.Expect(len(ddata.Links)).To(Equal(7))
	g.Expect(len(ddata.Nodes)).To(Equal(14))
}
