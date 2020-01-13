package cocago

import (
	"fmt"
	"github.com/phodal/coca/pkg/domain/trial"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"reflect"
)

var currentPackage *trial.CodePackage

type CocagoParser struct {
}

func NewCocagoParser() *CocagoParser {
	currentPackage = &trial.CodePackage{}
	return &CocagoParser{}
}

func (n *CocagoParser) ProcessFile(fileName string) trial.CodeFile {
	absPath, _ := filepath.Abs(fileName)
	content, _ := ioutil.ReadFile(absPath)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, string(content), 0)
	if err != nil {
		panic(err)
	}

	codeFile := n.Visitor(f, fset, fileName)
	currentPackage.CodeFiles = append(currentPackage.CodeFiles, *codeFile)
	return *codeFile
}

func (n *CocagoParser) Visitor(f *ast.File, fset *token.FileSet, fileName string) *trial.CodeFile {
	var currentStruct trial.CodeDataStruct
	var currentFile trial.CodeFile
	var currentFunc *trial.CodeFunction

	currentFile.FullName = fileName
	var funcType = ""

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:

		case *ast.File:
			currentFile.PackageName = x.Name.String()
		case *ast.ImportSpec:
			path := x.Path.Value
			cleanPath := path[1 : len(path)-1]
			currentFile.Imports = append(currentFile.Imports, cleanPath)
		case *ast.TypeSpec:
			currentStruct = trial.CodeDataStruct{}
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
		default:
			if reflect.TypeOf(x) != nil {
				fmt.Println("Visitor case ", reflect.TypeOf(x))
			}
		}
		return true
	})

	return &currentFile
}

func AddNestedFunction(currentFunc *trial.CodeFunction, x *ast.FuncType) {

}

func AddFunctionDecl(currentStruct trial.CodeDataStruct, x *ast.FuncDecl, currentFile *trial.CodeFile) *trial.CodeFunction {
	recv := ""
	if x.Recv != nil {
		recv = BuildReceiver(x, recv)
	}
	codeFunc := BuildFunction(x)

	if recv != "" {
		member := GetMemberFromFile(*currentFile, recv)
		if member != nil {
			member.MethodNodes = append(member.MethodNodes, *codeFunc)
		} else {
			createMember(currentStruct)
			// todo
		}
	} else {
		member := GetMemberFromFile(*currentFile, "default")
		if member == nil {
			member = &trial.CodeMember{
				DataStructID: "default",
				Type:         "method",
			}
		}

		member.MethodNodes = append(member.MethodNodes, *codeFunc)
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
			fmt.Println("AddFunctionDecl", reflect.TypeOf(x))
		}
	}
	return recv
}

func getStarExprName(starExpr ast.StarExpr) string {
	switch x := starExpr.X.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.SelectorExpr:
		return getSelectorName(*x)
	default:
		fmt.Println("getStarExprName", reflect.TypeOf(x))
		return ""
	}
}

func BuildFunction(x *ast.FuncDecl) *trial.CodeFunction {
	codeFunc := &trial.CodeFunction{
		Name: x.Name.String(),
	}

	if x.Type.Params != nil {
		codeFunc.Parameters = append(codeFunc.Parameters, BuildFieldToProperty(x.Type.Params.List)...)
	}

	if x.Type.Results != nil {
		codeFunc.ReturnTypes = append(codeFunc.Parameters, BuildFieldToProperty(x.Type.Results.List)...)
	}

	for _, item := range x.Body.List {
		BuildMethodCall(codeFunc, item)
	}
	return codeFunc
}

func BuildMethodCall(codeFunc *trial.CodeFunction, item ast.Stmt) {
	switch it := item.(type) {
	case *ast.ExprStmt:
		BuildMethodCallExprStmt(it, codeFunc)
	default:
		fmt.Println("methodCall", reflect.TypeOf(it))
	}
}

func BuildMethodCallExprStmt(it *ast.ExprStmt, codeFunc *trial.CodeFunction) {
	switch expr := it.X.(type) {
	case *ast.CallExpr:
		selector, selName := BuildExpr(expr.Fun.(ast.Expr))
		call := trial.CodeCall{
			Package:    "",
			Type:       "",
			Class:      selector,
			MethodName: selName,
		}

		for _, arg := range expr.Args {
			value, kind := BuildExpr(arg.(ast.Expr))
			property := &trial.CodeProperty{
				TypeName: value,
				TypeType: kind,
			}

			call.Parameters = append(call.Parameters, *property)
		}

		codeFunc.MethodCalls = append(codeFunc.MethodCalls, call)
	}
}

func BuildExpr(expr ast.Expr) (string, string) {
	switch x := expr.(type) {
	case *ast.SelectorExpr:
		selector := x.X.(*ast.Ident).String()
		selName := x.Sel.Name
		return selector, selName
	case *ast.BasicLit:
		return x.Value, x.Kind.String()
	case *ast.Ident:
		return x.Name, x.Obj.Kind.String()
	default:
		fmt.Println("BuildExpr", reflect.TypeOf(x))
	}
	return "", ""
}

func createMember(codeDataStruct trial.CodeDataStruct) {

}

func GetMemberFromFile(file trial.CodeFile, recv string) *trial.CodeMember {
	var identMember *trial.CodeMember
	for _, member := range file.Members {
		if member.DataStructID == recv {
			identMember = member
		}
	}
	return identMember
}

func BuildFieldToProperty(fieldList []*ast.Field) []trial.CodeProperty {
	var properties []trial.CodeProperty
	for _, field := range fieldList {
		typeName, typeType := BuildPropertyField(field)
		property := trial.CodeProperty{
			Modifiers: nil,
			Name:      getFieldName(field),
			TypeType:  typeType,
			TypeName:  typeName,
		}
		properties = append(properties, property)
	}
	return properties
}

func getFieldName(field *ast.Field) string {
	if len(field.Names) < 1 {
		return ""
	}
	return field.Names[0].Name
}

func BuildPropertyField(field *ast.Field) (string, string) {
	var typeName string
	var typeType string
	switch x := field.Type.(type) {
	case *ast.Ident:
		typeType = "Identify"
		typeName = x.String()
	case *ast.ArrayType:
		typeType = "ArrayType"
		switch typeX := x.Elt.(type) {
		case *ast.Ident:
			typeName = typeX.String()
		case *ast.SelectorExpr:
			typeName = getSelectorName(*typeX)
		default:
			fmt.Println("BuildPropertyField ArrayType", reflect.TypeOf(x.Elt))
		}
	case *ast.FuncType:
		typeType = "Function"
		typeName = "func"
	case *ast.StarExpr:
		typeName = getStarExprName(*x)
		typeType = "Star"
	case *ast.SelectorExpr:
		typeName = getSelectorName(*x)
	default:
		fmt.Println("BuildPropertyField", reflect.TypeOf(x))
	}
	return typeName, typeType
}

func getSelectorName(typeX ast.SelectorExpr) string {
	return typeX.X.(*ast.Ident).String() + "." + typeX.Sel.Name
}
