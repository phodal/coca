package unusedclasses

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"path/filepath"
	"testing"
)

func TestRefactoring(t *testing.T) {
	g := NewGomegaWithT(t)


	var parsedDeps []core_domain.CodeDataStruct
	codePath := "../../../../_fixtures/count/call.json"
	codePath = filepath.FromSlash(codePath)
	file := cmd_util.ReadFile(codePath)
	_ = json.Unmarshal(file, &parsedDeps)

	results := Refactoring(parsedDeps)

	g.Expect(len(results)).To(Equal(1))
}
