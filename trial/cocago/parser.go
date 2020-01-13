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

type WalkFunc func(ast.Node) (ast.Node, bool)

func ProcessFile(fileName string) trial.CodeFile {

	absPath, _ := filepath.Abs(fileName)
	content, _ := ioutil.ReadFile(absPath)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, string(content), 0)
	if err != nil {
		panic(err)
	}

	return Visitor(f, fset, fileName)
}

func Visitor(f *ast.File, fset *token.FileSet, fileName string) trial.CodeFile {
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
			BuildStructType(currentStruct, x, &currentFile)

		case *ast.FuncDecl:
			funcName := x.Name.String()
			recv := ""
			for _, item := range x.Recv.List {
				switch x := item.Type.(type) {
				case *ast.StarExpr:
					recv = x.X.(*ast.Ident).Name
				}
			}

			if recv != "" {

			} else {
				member := trial.CodeMember{
					DataStructID: currentStruct.Name,
					Type:         "function",
					Name:         funcName,
				}
				currentFile.Members = append(currentFile.Members, member)
			}
		}
		return true
	})

	return currentFile
}

func BuildStructType(currentStruct trial.CodeDataStruct, x *ast.StructType, currentFile *trial.CodeFile) {
	member := trial.CodeMember{
		DataStructID: currentStruct.Name,
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
		member.FileID = currentFile.FullName
		currentStruct.Properties = append(currentStruct.Properties, property)
	}
	currentFile.Members = append(currentFile.Members, member)
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
			typeName = typeX.X.(*ast.Ident).String() + "." + typeX.Sel.Name
		default:
			fmt.Println(reflect.TypeOf(x.Elt))
		}
	case *ast.FuncType:
		typeType = "Function"
		typeName = "func"
	default:
		fmt.Println(reflect.TypeOf(x))
	}
	return typeName, typeType
}
