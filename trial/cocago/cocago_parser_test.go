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
	t.Parallel()
	g := NewGomegaWithT(t)

	var test = "data_struct_property"
	results := testParser.ProcessFile("testdata/" + test + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/" + test + ".json")).To(Equal(true))
}

func Test_DataStructWithFuncType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var test = "struct_with_func"
	results := testParser.ProcessFile("testdata/" + test + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/" + test + ".json")).To(Equal(true))
}

func Test_DataStructWithFuncDecl(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var test = "struct_with_func_decl"
	results := testParser.ProcessFile("testdata/" + test + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/" + test + ".json")).To(Equal(true))
}

func Test_DataStructZero(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var test = "struct_type_zero"
	results := testParser.ProcessFile("testdata/" + test + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/" + test + ".json")).To(Equal(true))
}

func Test_NormalMethod(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var test = "normal_method"
	results := testParser.ProcessFile("testdata/" + test + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/" + test + ".json")).To(Equal(true))
}

func Test_MethodCallWithHelloWorld(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var test = "hello_world"
	results := testParser.ProcessFile("testdata/" + test + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/" + test + ".json")).To(Equal(true))
}

// todo: support it
func Test_NestedMethod(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var test = "nested_method"
	results := testParser.ProcessFile("testdata/" + test + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/" + test + ".json")).To(Equal(true))
}

func Test_BasicInterface(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var test = "basic_interface"
	results := testParser.ProcessFile("testdata/" + test + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/" + test + ".json")).To(Equal(true))
}

