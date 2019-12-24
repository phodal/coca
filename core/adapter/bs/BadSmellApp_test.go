package bs

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestBadSmellApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	bsApp := new(BadSmellApp)
	codePath := "../../../_fixtures/call"
	bsList := bsApp.AnalysisPath(codePath, nil)

	g.Expect(len(bsList)).To(Equal(1))
}
