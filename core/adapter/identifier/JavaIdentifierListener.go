package identifier

import (
	"coca/core/languages/java"
	"coca/core/models"
)

var node *JIdentifier

type JavaIdentifierListener struct {
	parser.BaseJavaParserListener
}

func (s *JavaIdentifierListener) EnterPackageDeclaration(ctx *parser.PackageDeclarationContext) {
	node.Package = ctx.QualifiedName().GetText()
}

func (s *JavaIdentifierListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	node.Type = "Class"
	if ctx.IDENTIFIER() != nil {
		node.Name = ctx.IDENTIFIER().GetText()
	}

	if ctx.EXTENDS() != nil {
		node.ExtendsName = ctx.TypeType().GetText()
	}
}

func (s *JavaIdentifierListener) EnterInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().GetColumn()
	name := ctx.IDENTIFIER().GetText()
	//XXX: find the start position of {, not public
	typeType := ctx.TypeTypeOrVoid().GetText()

	method := &models.JMethod{
		Name: name,
		Type: typeType,
		StartLine: startLine,
		StartLinePosition: startLinePosition,
		StopLine: stopLine,
		StopLinePosition: stopLinePosition,
		Override: isOverrideMethod,
		Annotation: nil,
	}
	node.AddMethod(*method)
}

var isOverrideMethod = false

func (s *JavaIdentifierListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().GetColumn()
	name := ctx.IDENTIFIER().GetText()
	//XXX: find the start position of {, not public

	typeType := ctx.TypeTypeOrVoid().GetText()

	method := &models.JMethod{
		Name: name,
		Type: typeType,
		StartLine: startLine,
		StartLinePosition: startLinePosition,
		StopLine: stopLine,
		StopLinePosition: stopLinePosition,
		Override: isOverrideMethod,
		Annotation: nil,
	}

	node.AddMethod(*method)

	isOverrideMethod = false
}

func (s *JavaIdentifierListener) EnterAnnotation(ctx *parser.AnnotationContext) {
	// Todo: support override method
	annotationName := ctx.QualifiedName().GetText()
	if annotationName == "Override" {
		isOverrideMethod = true
	}
}

func (s *JavaIdentifierListener) EnterInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	node.Type = "Interface"
	node.Name = ctx.IDENTIFIER().GetText()
}

func (s *JavaIdentifierListener) InitNode(identifier *JIdentifier) {
	node = identifier
}
