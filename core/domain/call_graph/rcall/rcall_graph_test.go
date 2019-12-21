package rcall

import (
	"encoding/json"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"log"
	"testing"

	. "github.com/onsi/gomega"
)

func TestRCallGraph_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	analyser := NewRCallGraph()
	file := support.ReadFile("../../../../_fixtures/call_api_test.json")
	if file == nil {
		log.Fatal("lost file")
	}

	_ = json.Unmarshal(file, &parsedDeps)

	content := analyser.Analysis("com.phodal.pholedge.book.BookFactory.create", *&parsedDeps)

	g.Expect(content).To(Equal(`digraph G {
rankdir = LR;
edge [dir="back"];

}
`))
}