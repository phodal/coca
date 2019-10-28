package visitor;

import (
	. "../../language/java"
	"fmt"
)

var text string
var currentClz string
var currentPkg string
var imports []string

func NewJavaCallVisitor() *JavaCallVisitor {
	return new(JavaCallVisitor)
}

type JavaCallVisitor struct {
	BaseJavaParserVisitor
}

func (v *JavaCallVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *JavaCallVisitor) VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{} {
	text = ctx.QualifiedName().GetText()
	fmt.Println(text)
	return v.VisitChildren(ctx)
}

func (v *JavaCallVisitor) VisitImportDeclaration(ctx *ImportDeclarationContext) interface{} {
	_ = append(imports, ctx.QualifiedName().GetText())
	return v.VisitChildren(ctx)
}

func (v *JavaCallVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	currentClz = ctx.IDENTIFIER().GetText()
	return v.VisitChildren(ctx)
}

