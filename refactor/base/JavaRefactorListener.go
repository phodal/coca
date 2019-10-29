package base

import (
	. "../../language/java"
	. "./models"
)

var node *JFullIdentifier;

type JavaRefactorListener struct {
	BaseJavaParserListener
}

func (s *JavaRefactorListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {
	node.Pkg = ctx.QualifiedName().GetText()
}

func (s *JavaRefactorListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {
	node.Type = "Class"
	node.Name = ctx.IDENTIFIER().GetText()
}

func (s *JavaRefactorListener) EnterInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetTokenSource().GetCharPositionInLine()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().GetTokenSource().GetCharPositionInLine()
	name := ctx.IDENTIFIER().GetText()
	//XXX: find the start position of {, not public
	method := &JFullMethod{name, startLine, startLinePosition, stopLine, stopLinePosition}
	node.AddMethod(*method)
}

func (s *JavaRefactorListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetTokenSource().GetCharPositionInLine()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().GetTokenSource().GetCharPositionInLine()
	name := ctx.IDENTIFIER().GetText()
	//XXX: find the start position of {, not public
	method := &JFullMethod{name, startLine, startLinePosition, stopLine, stopLinePosition}
	node.AddMethod(*method)
}

func (s *JavaRefactorListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {
	node.Type = "Interface"
	node.Name = ctx.IDENTIFIER().GetText()
}

func (s *JavaRefactorListener) InitNode(identifier *JFullIdentifier) {
	node = identifier
}

