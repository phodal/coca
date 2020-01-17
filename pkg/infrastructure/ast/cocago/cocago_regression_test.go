package cocago

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cocatest"
	"testing"
)

func getRegressionFile(name string) string {
	return "testdata/regression/" + name
}

func Test_Regression1(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	filePath := getRegressionFile("coll_stack")
	results := testParser.ProcessFile(filePath + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, filePath+".json")).To(Equal(true))
}
