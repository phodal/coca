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

	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()

	pkgInfo := &JPkgInfo{node.Pkg, startLine, stopLine}
	node.SetPkgInfo(*pkgInfo)
}

func (s *JavaRefactorListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {
	importText := ctx.QualifiedName().GetText()
	if ctx.MUL() != nil {
		importText = importText + ".*"
	}
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
	//fmt.Println(ctx.TypeTypeOrVoid())
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

func (s *JavaRefactorListener) EnterTypeType(ctx *TypeTypeContext) {
	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()
	field := &JField{ctx.GetText(), node.Pkg, startLine, stopLine}
	node.AddField(*field)
}

func (s *JavaRefactorListener) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {
	identifiers := ctx.AllIDENTIFIER()
	for index, _ := range identifiers {
		context := ctx.IDENTIFIER(index)
		name := context.GetText()
		startLine := ctx.GetStart().GetLine()
		stopLine := ctx.GetStop().GetLine()

		field := &JField{name, node.Pkg, startLine, stopLine}
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

func (s *JavaRefactorListener) EnterLambdaParameters(ctx *LambdaParametersContext) {
	identifiers := ctx.AllIDENTIFIER()
	for index, _ := range identifiers {
		context := ctx.IDENTIFIER(index)
		name := context.GetText()
		startLine := ctx.GetStart().GetLine()
		stopLine := ctx.GetStop().GetLine()

		field := &JField{name, node.Pkg, startLine, stopLine}
		node.AddField(*field)
	}
}

func (s *JavaRefactorListener) EnterMethodCall(ctx *MethodCallContext) {
	text := ctx.IDENTIFIER().GetText()
	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()
	field := &JField{text, node.Pkg, startLine, stopLine}
	node.AddField(*field)
}

func (s *JavaRefactorListener) InitNode(identifier *JFullIdentifier) {
	node = identifier
}
