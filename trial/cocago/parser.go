package cocago

import (
	"fmt"
	"github.com/phodal/coca/pkg/domain"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

type WalkFunc func(ast.Node) (ast.Node, bool)

func ProcessFile(fileName string) domain.CodeFile {
	content, _  := ioutil.ReadFile(fileName)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", string(content), 0)
	if err != nil {
		panic(err)
	}

	return Visitor(f, fset)
}

func Visitor(f *ast.File, fset *token.FileSet) domain.CodeFile {
	var currentStruct domain.CodeDataStruct
	var currentFile domain.CodeFile
	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.File:
			currentFile.PackageName = x.Name.String()
		case *ast.Ident:
			s = x.Name
		case *ast.TypeSpec:
			currentStruct = domain.CodeDataStruct{}
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

func BuildStructType(currentStruct domain.CodeDataStruct, x *ast.StructType, currentFile *domain.CodeFile) {
	member := domain.CodeMember{
		DataStructID: currentStruct.ID,
		Type:         "struct",
	}
	for _, field := range x.Fields.List {
		fmt.Println(field.Names, field.Type)
		property := domain.CodeProperty{
			Modifiers: nil,
			Name:      field.Names[0].String(),
			Type:      "",
		}
		member.Properties = append(member.Properties, property)
	}
	currentFile.Members = append(currentFile.Members, member)
}
