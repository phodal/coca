package concept

import (
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"encoding/json"
	"log"
	"testing"

	. "github.com/onsi/gomega"
)

func TestConceptAnalyser_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	analyser := NewConceptAnalyser()
	file := support.ReadFile("../../../_fixtures/call_api_test.json")
	if file == nil {
		log.Fatal("lost file")
	}

	_ = json.Unmarshal(file, &parsedDeps)

	counts := analyser.Analysis(&parsedDeps)

	g.Expect(counts.Len()).To(Equal(4))
}