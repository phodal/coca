package evaluate

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"testing"
)

func TestAnalyser_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	analyser := NewEvaluateAnalyser()
	file := support.ReadFile("../../../_fixtures/evaluate/service.json")
	_ = json.Unmarshal(file, &parsedDeps)

	analyser.Analysis(parsedDeps)
	g.Expect(true).To(Equal(true))
}

func TestUpdate_Analyser_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	analyser := NewEvaluateAnalyser()
	file := support.ReadFile("../../../_fixtures/evaluate/service_lifecycle.json")
	_ = json.Unmarshal(file, &parsedDeps)

	analyser.Analysis(parsedDeps)
	g.Expect(true).To(Equal(true))
}