package cocago

import (
	"bytes"
	"fmt"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
)

var currentPackage *core_domain.CodePackage

type CocagoParser struct {
}

var debug = false
var output io.Writer

func NewCocagoParser() *CocagoParser {
	currentPackage = &core_domain.CodePackage{}
	output = os.Stdout
	return &CocagoParser{}
}

func (n *CocagoParser) SetOutput(isDebug bool) io.Writer {
	output = new(bytes.Buffer)
	debug = isDebug

	return output
}

func (n *CocagoParser) ProcessFile(fileName string) core_domain.CodeFile {
	absPath, _ := filepath.Abs(fileName)
	content, _ := ioutil.ReadFile(absPath)

	fmt.Fprintf(output, "process file %s\n", fileName)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, string(content), 0)
	if err != nil {
		panic(err)
	}

	codeFile := n.Visitor(f, fset, fileName)
	currentPackage.CodeFiles = append(currentPackage.CodeFiles, *codeFile)
	return *codeFile
}

func (n *CocagoParser) Visitor(f *ast.File, fset *token.FileSet, fileName string) *core_domain.CodeFile {
	var currentStruct core_domain.CodeDataStruct
	var currentFile core_domain.CodeFile
	var currentFunc *core_domain.CodeFunction

	currentFile.FullName = fileName
	var funcType = ""
	var lastIdent = ""

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:
			lastIdent = x.Name
		case *ast.File:
			currentFile.PackageName = x.Name.String()
		case *ast.ImportSpec:
			imp := BuildImport(x)
			currentFile.Imports = append(currentFile.Imports, *imp)
		case *ast.TypeSpec:
			currentStruct = core_domain.CodeDataStruct{}
			currentStruct.Name = x.Name.String()
		case *ast.StructType:
			AddStructType(currentStruct, x, &currentFile)
		case *ast.FuncDecl:
			funcType = "FuncDecl"
			currentFunc = AddFunctionDecl(currentStruct, x, &currentFile)
		case *ast.FuncType:
			if funcType != "FuncDecl" {
				AddNestedFunction(currentFunc, x)
			}

			funcType = ""
		case *ast.InterfaceType:
			AddInterface(x, lastIdent, &currentFile)
		default:
			if reflect.TypeOf(x) != nil {
				fmt.Fprintf(output, "Visitor case %s\n", reflect.TypeOf(x))
			}
		}
		return true
	})

	return &currentFile
}

func BuildImport(x *ast.ImportSpec) *core_domain.CodeImport {
	path := x.Path.Value
	cleanPath := path[1 : len(path)-1]
	asName := ""
	if x.Name != nil {
		asName = x.Name.String()
	}
	imp := &core_domain.CodeImport{
		Source:     cleanPath,
		AsName:     asName,
		ImportName: "",
		UsageName:  nil,
		Scope:      "",
	}
	return imp
}

func AddInterface(x *ast.InterfaceType, ident string, codeFile *core_domain.CodeFile) {
	properties := BuildFieldToProperty(x.Methods.List)

	dataStruct := core_domain.CodeDataStruct{
		Name:            ident,
		InOutProperties: properties,
	}

	member := core_domain.CodeMember{
		DataStructID: dataStruct.Name,
		Type:         "interface",
	}

	codeFile.Members = append(codeFile.Members, &member)
	codeFile.DataStructures = append(codeFile.DataStructures, dataStruct)
}

func AddNestedFunction(currentFunc *core_domain.CodeFunction, x *ast.FuncType) {

}

func AddFunctionDecl(currentStruct core_domain.CodeDataStruct, x *ast.FuncDecl, currentFile *core_domain.CodeFile) *core_domain.CodeFunction {
	recv := ""
	if x.Recv != nil {
		recv = BuildReceiver(x, recv)
	}
	codeFunc := BuildFunction(x)

	if recv != "" {
		member := GetMemberFromFile(*currentFile, recv)
		if member != nil {
			member.FunctionNodes = append(member.FunctionNodes, *codeFunc)
		} else {
			createMember(currentStruct)
			// todo
		}
	} else {
		member := GetMemberFromFile(*currentFile, "default")
		if member == nil {
			member = &core_domain.CodeMember{
				DataStructID: "default",
				Type:         "method",
			}
		}

		member.FunctionNodes = append(member.FunctionNodes, *codeFunc)
		currentFile.Members = append(currentFile.Members, member)
	}

	return codeFunc
}

func BuildReceiver(x *ast.FuncDecl, recv string) string {
	for _, item := range x.Recv.List {
		switch x := item.Type.(type) {
		case *ast.StarExpr:
			recv = getStarExprName(*x)
		case *ast.Ident:
			recv = x.Name
		default:
			fmt.Fprintf(output, "AddFunctionDecl %s\n", reflect.TypeOf(x))
		}
	}
	return recv
}

func BuildMethodCall(codeFunc *core_domain.CodeFunction, item ast.Stmt) {
	switch it := item.(type) {
	case *ast.ExprStmt:
		BuildMethodCallExprStmt(it, codeFunc)
	default:
		fmt.Fprintf(output, "methodCall %s\n", reflect.TypeOf(it))
	}
}

func BuildMethodCallExprStmt(it *ast.ExprStmt, codeFunc *core_domain.CodeFunction) {
	switch expr := it.X.(type) {
	case *ast.CallExpr:
		selector, selName := BuildExpr(expr.Fun.(ast.Expr))
		call := core_domain.CodeCall{
			Package:    "",
			Type:       "",
			Class:      selector,
			MethodName: selName,
		}

		for _, arg := range expr.Args {
			value, kind := BuildExpr(arg.(ast.Expr))
			property := &core_domain.CodeProperty{
				TypeValue: value,
				TypeType:  kind,
			}

			call.Parameters = append(call.Parameters, *property)
		}

		codeFunc.MethodCalls = append(codeFunc.MethodCalls, call)
	}
}

func BuildExpr(expr ast.Expr) (string, string) {
	switch x := expr.(type) {
	case *ast.SelectorExpr:
		selector := ""
		switch sele := x.X.(type) {
		case *ast.Ident:
			selector = sele.String()
		}

		selName := x.Sel.Name
		return selector, selName
	case *ast.BasicLit:
		return x.Value, x.Kind.String()
	case *ast.Ident:
		name := ""
		if x.Obj != nil {
			name = x.Obj.Kind.String()
		}
		return x.Name, name
	default:
		fmt.Fprintf(output, "BuildExpr %s\n", reflect.TypeOf(x))
	}
	return "", ""
}

func createMember(codeDataStruct core_domain.CodeDataStruct) {

}

func GetMemberFromFile(file core_domain.CodeFile, recv string) *core_domain.CodeMember {
	var identMember *core_domain.CodeMember
	for _, member := range file.Members {
		if member.DataStructID == recv {
			identMember = member
		}
	}
	return identMember
}

func getFieldName(field *ast.Field) string {
	if len(field.Names) < 1 {
		return ""
	}
	return field.Names[0].Name
}

func AddStructType(currentStruct core_domain.CodeDataStruct, x *ast.StructType, currentFile *core_domain.CodeFile) {
	member := core_domain.CodeMember{
		DataStructID: currentStruct.Name,
		Type:         "struct",
	}
	for _, field := range x.Fields.List {
		property := BuildPropertyField(getFieldName(field), field)
		member.FileID = currentFile.FullName
		currentStruct.InOutProperties = append(currentStruct.InOutProperties, *property)
	}
	currentFile.Members = append(currentFile.Members, &member)
	currentFile.DataStructures = append(currentFile.DataStructures, currentStruct)
}
