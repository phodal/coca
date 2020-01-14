package processor

import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_ProcessPackage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	debug := true
	results := ProcessPackage("../../../../pkg/domain", debug)
	g.Expect(len(results)).To(Equal(27))
}