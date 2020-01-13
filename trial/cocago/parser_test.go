package cocago

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cocatest"
	"path/filepath"
	"testing"
)

func Test_DataStructProperty(t *testing.T) {
	g := NewGomegaWithT(t)

	abs, _ := filepath.Abs("testdata/data_struct_property.code")
	results := ProcessFile(abs)
	g.Expect(len(results.Members)).To(Equal(1))
	properties := results.Members[0].Properties

	g.Expect(len(properties)).To(Equal(5))
	g.Expect(properties[0].Name).To(Equal("FullName"))
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/data_struct_property.json")).To(Equal(true))
}
