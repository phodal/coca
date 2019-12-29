package unused_classes

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"testing"
)

func TestRefactoring(t *testing.T) {
	g := NewGomegaWithT(t)


	var parsedDeps []models.JClassNode
	file := support.ReadFile("../../../../_fixtures/count/call.json")
	_ = json.Unmarshal(file, &parsedDeps)

	results := Refactoring(parsedDeps)

	g.Expect(len(results)).To(Equal(1))
}
