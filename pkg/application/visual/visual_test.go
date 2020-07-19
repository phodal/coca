package visual

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"path/filepath"
	"testing"
)

func TestNewTodoApp(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []core_domain.CodeDataStruct
	codePath := "../../../_fixtures/evaluate/service_long_parameters.json"
	codePath = filepath.FromSlash(codePath)
	file := cmd_util.ReadFile(codePath)
	_ = json.Unmarshal(file, &parsedDeps)

	ddata := FromDeps(parsedDeps)

	g.Expect(len(ddata.Links)).To(Equal(0))
	g.Expect(len(ddata.Nodes)).To(Equal(3))
}
