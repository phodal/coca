package call

import (
	. "../../language/java"
	. "./models"
	"github.com/antlr/antlr4/runtime/Go/antlr"
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
	startLinePosition := ctx.GetStart().(antlr.TokenSource).GetCharPositionInLine()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().(antlr.TokenSource).GetCharPositionInLine()
	name := ctx.IDENTIFIER().GetText()
	//XXX: find the start position of {, not public
	method := &JMethod{name, startLine, startLinePosition, stopLine, stopLinePosition}
	node.AddMethod(*method)
}

func (s *JavaIdentifierListener) EnterMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().(antlr.TokenSource).GetCharPositionInLine()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().(antlr.TokenSource).GetCharPositionInLine()
	name := ctx.IDENTIFIER().GetText()
	//XXX: find the start position of {, not public
	method := &JMethod{name, startLine, startLinePosition, stopLine, stopLinePosition}
	node.AddMethod(*method)
}

func (s *JavaIdentifierListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {
	node.Type = "Interface"
	node.Name = ctx.IDENTIFIER().GetText()
}

func (s *JavaIdentifierListener) InitNode(identifier *JIdentifier) {
	node = identifier
}

