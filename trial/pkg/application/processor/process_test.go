package processor

import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_ProcessPackage(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	results := ProcessPackage("../../../../pkg/domain", true)
	g.Expect(len(results)).To(Equal(27))
}