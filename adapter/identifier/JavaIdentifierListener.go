package identifier

import (
	. "github.com/phodal/coca/adapter/models"
	. "github.com/phodal/coca/language/java"
)

var node *JIdentifier;

type JavaIdentifierListener struct {
	BaseJavaParserListener
}

func (s *JavaIdentifierListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {
	node.Pkg = ctx.QualifiedName().GetText()
}

func (s *JavaIdentifierListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {
	node.Type = "Class"
	node.Name = ctx.IDENTIFIER().GetText()
}

func (s *JavaIdentifierListener) EnterInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().GetColumn()
	name := ctx.IDENTIFIER().GetText()
	//XXX: find the start position of {, not public
	typeType := ctx.TypeTypeOrVoid().GetText()

	method := &JMethod{name, typeType, startLine, startLinePosition, stopLine, stopLinePosition}
	node.AddMethod(*method)
}

func (s *JavaIdentifierListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().GetColumn()
	name := ctx.IDENTIFIER().GetText()
	//XXX: find the start position of {, not public

	typeType := ctx.TypeTypeOrVoid().GetText()

	method := &JMethod{name, typeType, startLine, startLinePosition, stopLine, stopLinePosition}
	node.AddMethod(*method)
}

func (s *JavaIdentifierListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {
	node.Type = "Interface"
	node.Name = ctx.IDENTIFIER().GetText()
}

func (s *JavaIdentifierListener) InitNode(identifier *JIdentifier) {
	node = identifier
}

