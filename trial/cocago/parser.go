package cocago

import (
	"fmt"
	"github.com/phodal/coca/pkg/domain/trial"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"reflect"
)

type WalkFunc func(ast.Node) (ast.Node, bool)

func ProcessFile(fileName string) trial.CodeFile {
	content, _ := ioutil.ReadFile(fileName)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", string(content), 0)
	if err != nil {
		panic(err)
	}

	return Visitor(f, fset)
}

func Visitor(f *ast.File, fset *token.FileSet) trial.CodeFile {
	var currentStruct trial.CodeDataStruct
	var currentFile trial.CodeFile
	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.File:
			currentFile.PackageName = x.Name.String()
		case *ast.Ident:
			s = x.Name
		case *ast.TypeSpec:
			currentStruct = trial.CodeDataStruct{}
			currentStruct.Name = x.Name.String()
			currentStruct.ID = f.Name.String()
		case *ast.StructType:
			BuildStructType(currentStruct, x, &currentFile)
		}
		if s != "" {
			//fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}
		return true
	})

	return currentFile
}

func BuildStructType(currentStruct trial.CodeDataStruct, x *ast.StructType, currentFile *trial.CodeFile) {
	member := trial.CodeMember{
		DataStructID: currentStruct.ID,
		Type:         "struct",
	}
	for _, field := range x.Fields.List {
		typeName, typeType := BuildPropertyField(field)
		property := trial.CodeProperty{
			Modifiers: nil,
			Name:      field.Names[0].String(),
			TypeType:  typeType,
			TypeName:  typeName,
		}
		member.Properties = append(member.Properties, property)
	}
	currentFile.Members = append(currentFile.Members, member)
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
			typeName = typeX.X.(*ast.Ident).String() + "." + typeX.Sel.Name
		default:
			fmt.Println(reflect.TypeOf(x.Elt))
		}
	}
	return typeName, typeType
}
