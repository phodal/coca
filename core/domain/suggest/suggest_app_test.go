package suggest

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
	analyser := NewSuggestApp()
	codePath := "../../../_fixtures/suggest/factory/factory_suggest.json"
	codePath = filepath.FromSlash(codePath)
	file := support.ReadFile(codePath)
	if file == nil {
		log.Fatal("lost file")
	}

	_ = json.Unmarshal(file, &parsedDeps)

	suggests := analyser.AnalysisPath(parsedDeps)

	g.Expect(len(suggests)).To(Equal(2))
	g.Expect(suggests[0].Pattern).To(Equal("factory"))
	g.Expect(suggests[0].Reason).To(Equal("too many constructor"))
	g.Expect(suggests[1].Pattern).To(Equal("factory, builder"))
	g.Expect(suggests[1].Reason).To(Equal("complex constructor, too many constructor, too many parameters"))
}