package cocago

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cocatest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

var testParser *CocagoParser

func setup() {
	testParser = NewCocagoParser()
}

func shutdown() {
	testParser = nil
}

func Test_DataStructProperty(t *testing.T) {
	g := NewGomegaWithT(t)

	results := testParser.ProcessFile("testdata/data_struct_property.code")
	g.Expect(len(results.Members)).To(Equal(1))
	properties := results.DataStructures[0].Properties

	g.Expect(len(properties)).To(Equal(5))
	g.Expect(properties[0].Name).To(Equal("FullName"))
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/data_struct_property.json")).To(Equal(true))
}

func Test_DataStructWithFuncType(t *testing.T) {
	g := NewGomegaWithT(t)

	results := testParser.ProcessFile("testdata/struct_with_func.code")
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/struct_with_func.json")).To(Equal(true))
}

func Test_DataStructWithFuncDecl(t *testing.T) {
	g := NewGomegaWithT(t)

	results := testParser.ProcessFile("testdata/struct_with_func_decl.code")
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/struct_with_func_decl.json")).To(Equal(true))
}

func Test_DataStructZero(t *testing.T) {
	g := NewGomegaWithT(t)

	results := testParser.ProcessFile("testdata/struct_type_zero.code")
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/struct_type_zero.json")).To(Equal(true))
}

func Test_Method(t *testing.T) {
	g := NewGomegaWithT(t)

	var test = "normal_method"
	results := testParser.ProcessFile("testdata/" + test + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/" + test + ".json")).To(Equal(true))
}
