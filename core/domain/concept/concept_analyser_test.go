package concept

import (
	"encoding/json"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"log"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
)

func TestConceptAnalyser_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	analyser := NewConceptAnalyser()
	codePath := "../../../_fixtures/call_api_test.json"
	codePath = filepath.FromSlash(codePath)

	file := support.ReadFile(codePath)
	if file == nil {
		log.Fatal("lost file")
	}

	_ = json.Unmarshal(file, &parsedDeps)

	counts := analyser.Analysis(&parsedDeps)

	g.Expect(counts.Len()).To(Equal(4))
}