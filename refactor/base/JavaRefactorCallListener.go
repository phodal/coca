package base

import (
	. "../../language/java"
	"fmt"
)

type JavaRefactorCallListener struct {
	BaseJavaParserListener
}

func (s *JavaRefactorCallListener) EnterCompilationUnit(ctx *CompilationUnitContext) {
	declaration := ctx.AllImportDeclaration()
	fmt.Println(declaration)
}