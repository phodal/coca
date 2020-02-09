package arch

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cocatest"
	"github.com/phodal/coca/cocatest/testhelper"
	"github.com/phodal/coca/pkg/application/arch/tequila"
	"path/filepath"
	"testing"
)

func TestArch_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/grammar/java/arch/step2-java"
	callNodes, identifiersMap, _ := testhelper.BuildAnalysisDeps(codePath)

	app := NewArchApp()
	results := app.Analysis(callNodes, identifiersMap)

	g.Expect(len(results.RelationList)).To(Equal(16))
	g.Expect(len(results.NodeList)).To(Equal(13))

	g.Expect(results.RelationList["domain.AggregateRootA->domain.AggregateRoot"].From).To(Equal("domain.AggregateRootA"))
	g.Expect(results.RelationList["domain.AggregateRootA->domain.AggregateRoot"].To).To(Equal("domain.AggregateRoot"))

	graph := results.ToDot(".", func(key string) bool {
		return false
	})

	g.Expect(len(graph.Nodes.Lookup)).To(Equal(13))
	g.Expect(len(graph.SubGraphs.SubGraphs)).To(Equal(3))

	jsonContent, _ := json.MarshalIndent(results, "", "\t")
	content := cmd_util.ReadFile(filepath.FromSlash(codePath + "/" + "results.json"))

	g.Expect(cocatest.JSONBytesEqual(jsonContent, content, "")).To(Equal(true))
}

func TestArch_AnalysisWithFans(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/grammar/java/arch/step2-java"
	callNodes, identifiersMap, _ := testhelper.BuildAnalysisDeps(codePath)

	app := NewArchApp()
	result := app.Analysis(callNodes, identifiersMap)

	fans := result.SortedByFan(tequila.MergePackageFunc)

	g.Expect(len(fans)).To(Equal(3))
	var fanPackage []string;
	for _, fan := range fans {
		fanPackage = append(fanPackage, fan.Name)
	}

	g.Eventually(fanPackage).Should(ConsistOf("domain", "repositories", "gateways"))
	g.Expect(fans[0].Name).To(Equal("domain"))
	g.Eventually(fans[0].FanIn).Should(Equal(2))
	g.Eventually(fans[0].FanOut).Should(Equal(0))
}
