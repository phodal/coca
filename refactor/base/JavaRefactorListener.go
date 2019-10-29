package base

import (
	. "../../language/java"
	. "./models"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

var node *JFullIdentifier;

type JavaRefactorListener struct {
	BaseJavaParserListener
}

func (s *JavaRefactorListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {
	node.Pkg = ctx.QualifiedName().GetText()

	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()

	pkgInfo := &JPkgInfo{node.Pkg, startLine, stopLine}
	node.SetPkgInfo(*pkgInfo)
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

	if ctx.IMPLEMENTS() != nil {
		context := ctx.TypeList()
		startLine := ctx.TypeList().GetStart().GetLine()
		stopLine := ctx.TypeList().GetStart().GetLine()

		split := strings.Split(context.GetText(), ",")
		for _, imp := range split {
			field := &JField{imp, node.Pkg, startLine, stopLine}
			node.AddField(*field)
		}
	}

	if ctx.EXTENDS() != nil {
		startLine := ctx.TypeType().GetStart().GetLine()
		stopLine := ctx.TypeType().GetStart().GetLine()
		field := &JField{ctx.TypeType().GetText(), node.Pkg, startLine, stopLine}
		node.AddField(*field)
	}
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

func (s *JavaRefactorListener) EnterAnnotation(ctx *AnnotationContext) {
	annotation := ctx.QualifiedName().GetText()

	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()

	field := &JField{annotation, node.Pkg, startLine, stopLine}
	node.AddField(*field)
}

func (s *JavaRefactorListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()
	field := &JField{"", node.Pkg, startLine, stopLine}

	if ctx.InterfaceDeclaration() != nil {
		field.Name = ctx.InterfaceDeclaration().GetText()
		node.AddField(*field)
	}

	if ctx.EnumDeclaration() != nil {
		field.Name = ctx.EnumDeclaration().GetText()
		node.AddField(*field)
	}

	if ctx.AnnotationTypeDeclaration() != nil {
		field.Name = ctx.AnnotationTypeDeclaration().GetText()
		node.AddField(*field)
	}
}

func (s *JavaRefactorListener) EnterLocalTypeDeclaration(ctx *LocalTypeDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()
	field := &JField{"", node.Pkg, startLine, stopLine}

	if ctx.ClassDeclaration() != nil {
		field.Name = ctx.ClassDeclaration().GetText()
	}

	if ctx.InterfaceDeclaration() != nil {
		field.Name = ctx.InterfaceDeclaration().GetText()
	}
	node.AddField(*field)
}

func (s *JavaRefactorListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {
	node.Type = "Interface"
	node.Name = ctx.IDENTIFIER().GetText()
}

func (s *JavaRefactorListener) InitNode(identifier *JFullIdentifier) {
	node = identifier
}
