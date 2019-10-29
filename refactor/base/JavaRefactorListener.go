package base

import (
	. "../../language/java"
	. "./models"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var node *JFullIdentifier;

type JavaRefactorListener struct {
	BaseJavaParserListener
}

func (s *JavaRefactorListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {
	node.Pkg = ctx.QualifiedName().GetText()
}

func (s *JavaRefactorListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {
	importText := ctx.QualifiedName().GetText()
	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()

	jImport := &JImport{importText, startLine, stopLine}

	node.AddImport(*jImport)
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

func (s *JavaRefactorListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {
	declarators := ctx.VariableDeclarators()
	variableName := declarators.GetParent().GetChild(0).(antlr.ParseTree).GetText()

	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()

	text := ctx.TypeType().GetText()
	if variableName != "" && text != "" {
		field := &JField{variableName, node.Pkg, startLine, stopLine}
		node.AddField(*field)
	}
}

func (s *JavaRefactorListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {
	node.Type = "Interface"
	node.Name = ctx.IDENTIFIER().GetText()
}

func (s *JavaRefactorListener) InitNode(identifier *JFullIdentifier) {
	node = identifier
}
