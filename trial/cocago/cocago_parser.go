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
	currentFile.FullName = fileName
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
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
			AddFunction(currentStruct, x, &currentFile)
		}
		return true
	})

	return &currentFile
}

func AddFunction(currentStruct trial.CodeDataStruct, x *ast.FuncDecl, currentFile *trial.CodeFile) {
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
}

func BuildReceiver(x *ast.FuncDecl, recv string) string {
	for _, item := range x.Recv.List {
		switch x := item.Type.(type) {
		case *ast.StarExpr:
			recv = getStarExprName(*x)
		case *ast.Ident:
			recv = x.Name
		default:
			fmt.Println("AddFunction", reflect.TypeOf(x))
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
		fieldList := x.Type.Params.List
		properties := BuildFieldToProperty(fieldList)
		codeFunc.Parameters = append(codeFunc.Parameters, properties...)
	}

	if x.Type.Results != nil {
		fieldList := x.Type.Results.List
		properties := BuildFieldToProperty(fieldList)
		codeFunc.ReturnTypes = append(codeFunc.Parameters, properties...)
	}
	return codeFunc
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
			Name:     getFieldName(field),
			TypeType: typeType,
			TypeName: typeName,
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

func AddStructType(currentStruct trial.CodeDataStruct, x *ast.StructType, currentFile *trial.CodeFile) {
	member := trial.CodeMember{
		DataStructID: currentStruct.Name,
		Type:         "struct",
	}
	for _, field := range x.Fields.List {
		typeName, typeType := BuildPropertyField(field)
		property := trial.CodeProperty{
			Modifiers: nil,
			Name:      getFieldName(field),
			TypeType:  typeType,
			TypeName:  typeName,
		}
		member.FileID = currentFile.FullName
		currentStruct.Properties = append(currentStruct.Properties, property)
	}
	currentFile.Members = append(currentFile.Members, &member)
	currentFile.DataStructures = append(currentFile.DataStructures, currentStruct)
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
