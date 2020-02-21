package ast_go

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cocatest"
	"io/ioutil"
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
	//buf := new(bytes.Buffer)
	//testParser.SetOutput(buf)
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
			"multiple_method_call",
			"multiple_method_call",
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

func Test_Method_Call(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
	}{
		{
			"local_var_method_call",
			"local_var_method_call",
		},
		{
			"param_method_call",
			"param_method_call",
		},
		{
			"var_inside_method_with_call",
			"var_inside_method_with_call",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := "testdata/method_call/" + tt.fileName
			if got := testParser.ProcessFile(filePath + ".code"); !cocatest.JSONFileBytesEqual(got, filePath+".json") {
				t.Errorf("ProcessFile() = %v, want %v", got, tt.fileName)
			}
		})
	}
}

func getFilePath(name string) string {
	return "testdata/node_infos/" + name
}

func Test_MemberFunctionNodesForTwoMethod(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	filePath := getFilePath("normal_method")
	results := testParser.ProcessFile(filePath + ".code")
	g.Expect(len(results.Members)).To(Equal(2))
	g.Expect(len(results.Members[0].FunctionNodes)).To(Equal(1))
	g.Expect(len(results.Members[1].FunctionNodes)).To(Equal(1))
}

func Test_basic_interface(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	filePath := getFilePath("basic_interface")
	results := testParser.ProcessFile(filePath + ".code")
	g.Expect(cocatest.JSONFileBytesEqual(results, filePath+".json")).To(Equal(true))
}

func Test_IdentFuncMember(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	code, _ := ioutil.ReadFile("testdata/node_infos/normal_method.code")
	results := testParser.IdentAnalysis(string(code), "core_domain:CodeDataStruct")
	g.Expect(results.Members[0].ID).To(Equal("testdata:ProcessTsString"))
}

// todo: support it
//func Test_NestedMethod(t *testing.T) {
//	t.Parallel()
//	g := NewGomegaWithT(t)
//
//	filePath := getFilePath("nested_method")
//	results := testParser.ProcessFile(filePath + ".code")
//	g.Expect(cocatest.JSONFileBytesEqual(results, filePath+".json")).To(Equal(true))
//}

func Test_RelatedImport(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	results := testParser.ProcessString(`
package goapp

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/cocago"
)

type GoIdentApp struct {
	Extensions interface{}
}

func (g *GoIdentApp) Analysis(code string, fileName string) core_domain.CodeFile {
	parser := cocago.NewCocagoParser()
	var imports []core_domain.CodeImport
	if g.Extensions != nil {
		imports = g.Extensions.([]core_domain.CodeImport)
	}
	return 	*parser.ProcessString(code, fileName, imports)
}
`, "call", nil)

	g.Expect(len(results.DataStructures)).To(Equal(1))
}
func Test_ShowShowSelfMethodCall(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	results := testParser.ProcessString(`
package node_infos

import "fmt"

func ShowChangeLogSummary() {
	changeMap := BuildChangeMap(commits)
	fmt.Println(changeMap)
}

func BuildChangeMap() {
	UpdateMessageForChange()
}

`, "self_method_call.go", nil)

	g.Expect(len(results.Members)).To(Equal(2))
	g.Expect(len(results.Members[0].FunctionNodes[0].FunctionCalls)).To(Equal(1))
}
