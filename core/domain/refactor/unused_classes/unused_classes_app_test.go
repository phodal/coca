package unused_classes

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"path/filepath"
	"testing"
)

func TestRefactoring(t *testing.T) {
	g := NewGomegaWithT(t)


	var parsedDeps []models.JClassNode
	codePath := "../../../../_fixtures/count/call.json"
	codePath = filepath.FromSlash(codePath)
	file := support.ReadFile(codePath)
	_ = json.Unmarshal(file, &parsedDeps)

	results := Refactoring(parsedDeps)

	g.Expect(len(results)).To(Equal(1))
}
