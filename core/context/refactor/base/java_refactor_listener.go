package base

import (
	models2 "github.com/phodal/coca/core/context/refactor/base/models"
	. "github.com/phodal/coca/languages/java"
	"strings"
	"unicode"
)

var node models2.JFullIdentifier;

type JavaRefactorListener struct {
	BaseJavaParserListener
}

func (s *JavaRefactorListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {
	node.Pkg = ctx.QualifiedName().GetText()

	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()

	pkgInfo := *&models2.JPkgInfo{node.Pkg, startLine, stopLine}
	node.SetPkgInfo(pkgInfo)
}

func (s *JavaRefactorListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {
	importText := ctx.QualifiedName().GetText()
	if ctx.MUL() != nil {
		importText = importText + ".*"
	}
	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()

	jImport := *&models2.JImport{importText, startLine, stopLine}

	node.AddImport(jImport)
}

func (s *JavaRefactorListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {
	node.Type = "Class"
	node.Name = ctx.IDENTIFIER().GetText()
}

// throws
func (s *JavaRefactorListener) EnterQualifiedNameList(ctx *QualifiedNameListContext) {
	for _, qualified := range ctx.AllQualifiedName() {
		startLine := ctx.GetStart().GetLine()
		stopLine := ctx.GetStop().GetLine()
		field := *&models2.JField{qualified.GetText(), node.Pkg, startLine, stopLine}
		node.AddField(field)
	}
}

func (s *JavaRefactorListener) EnterCatchType(ctx *CatchTypeContext) {
	for _, qualified := range ctx.AllQualifiedName() {
		startLine := ctx.GetStart().GetLine()
		stopLine := ctx.GetStop().GetLine()
		field := *&models2.JField{qualified.GetText(), node.Pkg, startLine, stopLine}
		node.AddField(field)
	}
}

func (s *JavaRefactorListener) EnterInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
	//fmt.Println(ctx.TypeTypeOrVoid())
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().GetColumn()
	name := ctx.IDENTIFIER().GetText()
	//XXX: find the start position of {, not public
	method := *&models2.JFullMethod{name, startLine, startLinePosition, stopLine, stopLinePosition}
	node.AddMethod(method)
}

func (s *JavaRefactorListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {
	node.Type = "Interface"
	node.Name = ctx.IDENTIFIER().GetText()
}

func (s *JavaRefactorListener) EnterTypeType(ctx *TypeTypeContext) {
	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()
	field := *&models2.JField{ctx.GetText(), node.Pkg, startLine, stopLine}
	node.AddField(field)
}

func (s *JavaRefactorListener) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {
	identifiers := ctx.AllIDENTIFIER()
	for index, _ := range identifiers {
		context := ctx.IDENTIFIER(index)
		name := context.GetText()
		startLine := ctx.GetStart().GetLine()
		stopLine := ctx.GetStop().GetLine()

		field := *&models2.JField{name, node.Pkg, startLine, stopLine}
		node.AddField(field)
	}
}

func (s *JavaRefactorListener) EnterAnnotation(ctx *AnnotationContext) {
	annotation := ctx.QualifiedName().GetText()

	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()

	field := *&models2.JField{annotation, node.Pkg, startLine, stopLine}
	node.AddField(field)
}

func (s *JavaRefactorListener) EnterLambdaParameters(ctx *LambdaParametersContext) {
	identifiers := ctx.AllIDENTIFIER()
	for index, _ := range identifiers {
		context := ctx.IDENTIFIER(index)
		name := context.GetText()
		startLine := ctx.GetStart().GetLine()
		stopLine := ctx.GetStop().GetLine()

		field := *&models2.JField{name, node.Pkg, startLine, stopLine}
		node.AddField(field)
	}
}

func (s *JavaRefactorListener) EnterMethodCall(ctx *MethodCallContext) {
	text := ctx.IDENTIFIER().GetText()
	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()
	field := *&models2.JField{text, node.Pkg, startLine, stopLine}
	node.AddField(field)
}

func (s *JavaRefactorListener) EnterExpressionList(ctx *ExpressionListContext) {
	for _, expression := range ctx.AllExpression() {
		expText := expression.GetText()
		if isUppercaseText(expText) {
			startLine := ctx.GetStart().GetLine()
			stopLine := ctx.GetStop().GetLine()
			field := *&models2.JField{expText, node.Pkg, startLine, stopLine}
			node.AddField(field)
		}
	}
}

func (s *JavaRefactorListener) EnterStatement(ctx *StatementContext) {
	for _, expression := range ctx.AllExpression() {
		expText := expression.GetText()
		if isUppercaseText(expText) {
			startLine := ctx.GetStart().GetLine()
			stopLine := ctx.GetStop().GetLine()
			field := *&models2.JField{expText, node.Pkg, startLine, stopLine}
			node.AddField(field)
		}
	}
}

func (s *JavaRefactorListener) EnterCreatedName(ctx *CreatedNameContext) {
	identifiers := ctx.AllIDENTIFIER()
	for index, _ := range identifiers {
		context := ctx.IDENTIFIER(index)
		name := context.GetText()
		startLine := ctx.GetStart().GetLine()
		stopLine := ctx.GetStop().GetLine()

		field := &models2.JField{name, node.Pkg, startLine, stopLine}
		node.AddField(*field)
	}
}

func (s *JavaRefactorListener) EnterExpression(ctx *ExpressionContext) {
	if ctx.Expression(0) != nil {
		expText := ctx.Expression(0).GetText()

		if isUppercaseText(expText) {
			startLine := ctx.GetStart().GetLine()
			stopLine := ctx.GetStop().GetLine()
			field := *&models2.JField{expText, node.Pkg, startLine, stopLine}
			node.AddField(field)
		}
	}

	if ctx.GetBop() == nil {
		return
	}

	if ctx.GetBop().GetText() != "." {
		return
	}

	if ctx.Expression(0) != nil {
		expText := ctx.Expression(0).GetText()
		// UUID.toString 形式的直接调用
		if ctx.MethodCall() != nil {
			if isUppercaseText(expText) {
				startLine := ctx.GetStart().GetLine()
				stopLine := ctx.GetStop().GetLine()
				field := &models2.JField{expText, node.Pkg, startLine, stopLine}
				node.AddField(*field)
			}
		}
	}
}

func isUppercaseText(text string) bool {
	return !strings.Contains(text, ".") && unicode.IsUpper([]rune(text)[0])
}

func (s *JavaRefactorListener) InitNode(identifier models2.JFullIdentifier) {
	node = identifier
}

func (s *JavaRefactorListener) GetNodeInfo() models2.JFullIdentifier {
	return node
}
