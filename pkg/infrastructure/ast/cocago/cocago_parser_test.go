package cocago

import (
	"bytes"
	"fmt"
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
	buf := new(bytes.Buffer)
	testParser.SetOutput(buf)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := getFilePath(tt.fileName)
			if got := testParser.ProcessFile(filePath + ".code"); !cocatest.JSONFileBytesEqual(got, filePath+".json") {
				t.Errorf("ProcessFile() = %v, want %v", got, tt.fileName)
			}
		})
	}
}

func getFilePath(name string) string {
	return "testdata/node_infos/" + name
}

func Test_basic_interface(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	filePath := getFilePath("basic_interface")
	results := testParser.ProcessFile(filePath + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, filePath+".json")).To(Equal(true))
}

// todo: support it
func Test_NestedMethod(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	filePath := getFilePath("nested_method")
	results := testParser.ProcessFile(filePath + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, filePath+".json")).To(Equal(true))
}

// var call
func Test_VarMethodCall(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	results := testParser.ProcessString(`

package main
 
import (
	"fmt"
	"sync"
)

var l *sync.Mutex
 
func main() {
	l = new(sync.Mutex)
	l.Lock()
	defer l.Unlock()
	fmt.Println("1")
}
`, "call")
	calls := results.Members[0].FunctionNodes[0].MethodCalls
	fmt.Println(calls)
	g.Expect(len(results.Fields)).To(Equal(1))
	g.Expect(len(calls)).To(Equal(2))
}
