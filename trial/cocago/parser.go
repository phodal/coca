package cocago

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

type WalkFunc func(ast.Node) (ast.Node, bool)

func ProcessFile(fileName string) {
	content, _  := ioutil.ReadFile(fileName)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", string(content), 0)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.Ident:
			s = x.Name
		}
		if s != "" {
			fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}
		return true
	})
}
