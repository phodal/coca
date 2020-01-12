package cocago

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"reflect"
)

type WalkFunc func(ast.Node) (ast.Node, bool)

func ProcessFile(fileName string) {
	content, _  := ioutil.ReadFile(fileName)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", string(content), 0)
	if err != nil {
		panic(err)
	}

	Visitor(f, fset)
}

func Visitor(f *ast.File, fset *token.FileSet) {
	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.File:
			//packageName := x.Name
		case *ast.Ident:
			s = x.Name
		case *ast.TypeSpec:
			fmt.Println(x.Name)
		case *ast.StructType:
			fmt.Println(reflect.TypeOf(x))
			for _, field := range x.Fields.List {
				fmt.Println(field.Names, field.Type)
			}
		}
		if s != "" {
			//fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}
		return true
	})
}
