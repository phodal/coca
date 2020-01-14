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

func TestCocagoParser_ProcessFile(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
	}{
		{
			"data_struct_property",
			"data_struct_property",
		},
		{
			"struct_with_func",
			"struct_with_func",
		},
		{
			"struct_with_func_decl",
			"struct_with_func_decl",
		},
		{
			"struct_type_zero",
			"struct_type_zero",
		},
		{
			"normal_method",
			"normal_method",
		},
		{
			"hello_world",
			"hello_world",
		},
		{
			"basic_interface",
			"basic_interface",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &CocagoParser{}
			filePath := getFilePath(tt.fileName)
			if got := n.ProcessFile(filePath + ".code"); !cocatest.JSONFileBytesEqual(got, filePath+".json") {
				t.Errorf("ProcessFile() = %v, want %v", got, tt.fileName)
			}
		})
	}
}

func getFilePath(name string) string {
	return "testdata/node_infos/" + name
}

// todo: support it
func Test_NestedMethod(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	filePath := getFilePath("nested_method")
	results := testParser.ProcessFile(filePath + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, filePath+".json")).To(Equal(true))
}
