package concept

import (
	"encoding/json"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/core/domain"
	"log"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
)

func TestConceptAnalyser_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []domain.JClassNode
	analyser := NewConceptAnalyser()
	codePath := "../../../_fixtures/call/call_api_test.json"
	codePath = filepath.FromSlash(codePath)

	file := cmd_util.ReadFile(codePath)
	if file == nil {
		log.Fatal("lost file")
	}

	_ = json.Unmarshal(file, &parsedDeps)

	counts := analyser.Analysis(&parsedDeps)

	g.Expect(counts.Len()).To(Equal(4))
}
