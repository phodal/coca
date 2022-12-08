package base

import (
	. "github.com/modernizing/coca/languages/java"
	model "github.com/modernizing/coca/pkg/application/refactor/base/models"
	"strings"
	"unicode"
)

var node model.JFullIdentifier

type JavaRefactorListener struct {
	BaseJavaParserListener
}

func (s *JavaRefactorListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {
	node.Pkg = ctx.QualifiedName().GetText()

	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()

	pkgInfo := model.JPkgInfo{Name: node.Pkg, StartLine: startLine, StopLine: stopLine}
	node.SetPkgInfo(pkgInfo)
}

func (s *JavaRefactorListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {
	importText := ctx.QualifiedName().GetText()
	if ctx.MUL() != nil {
		importText = importText + ".*"
	}
	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()

	jImport := model.JImport{Name: importText, StartLine: startLine, StopLine: stopLine}

	node.AddImport(jImport)
}

func (s *JavaRefactorListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {
	node.Type = "Class"
	node.Name = ctx.Identifier().GetText()
}

func (s *JavaRefactorListener) EnterQualifiedNameList(ctx *QualifiedNameListContext) {
	for _, qualified := range ctx.AllQualifiedName() {
		startLine := ctx.GetStart().GetLine()
		stopLine := ctx.GetStop().GetLine()
		field := model.JField{Name: qualified.GetText(), Source: node.Pkg, StartLine: startLine, StopLine: stopLine}
		node.AddField(field)
	}
}

func (s *JavaRefactorListener) EnterCatchType(ctx *CatchTypeContext) {
	for _, qualified := range ctx.AllQualifiedName() {
		startLine := ctx.GetStart().GetLine()
		stopLine := ctx.GetStop().GetLine()
		field := model.JField{Name: qualified.GetText(), Source: node.Pkg, StartLine: startLine, StopLine: stopLine}
		node.AddField(field)
	}
}

func (s *JavaRefactorListener) EnterInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
	//fmt.Println(ctx.TypeTypeOrVoid())
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().GetColumn()
	name := ctx.Identifier().GetText()
	method := model.JFullMethod{Name: name, StartLine: startLine, StartLinePosition: startLinePosition, StopLine: stopLine, StopLinePosition: stopLinePosition}
	node.AddMethod(method)
}

func (s *JavaRefactorListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {
	node.Type = "Interface"
	node.Name = ctx.Identifier().GetText()
}

func (s *JavaRefactorListener) EnterTypeType(ctx *TypeTypeContext) {
	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()
	field := model.JField{Name: ctx.GetText(), Source: node.Pkg, StartLine: startLine, StopLine: stopLine}
	node.AddField(field)
}

func (s *JavaRefactorListener) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {
	identifiers := ctx.AllIdentifier()
	for index := range identifiers {
		context := ctx.Identifier(index)
		name := context.GetText()
		startLine := ctx.GetStart().GetLine()
		stopLine := ctx.GetStop().GetLine()

		field := model.JField{Name: name, Source: node.Pkg, StartLine: startLine, StopLine: stopLine}
		node.AddField(field)
	}
}

func (s *JavaRefactorListener) EnterAnnotation(ctx *AnnotationContext) {
	annotation := ctx.QualifiedName().GetText()

	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()

	field := model.JField{Name: annotation, Source: node.Pkg, StartLine: startLine, StopLine: stopLine}
	node.AddField(field)
}

func (s *JavaRefactorListener) EnterLambdaParameters(ctx *LambdaParametersContext) {
	identifiers := ctx.AllIdentifier()
	for index := range identifiers {
		context := ctx.Identifier(index)
		name := context.GetText()
		startLine := ctx.GetStart().GetLine()
		stopLine := ctx.GetStop().GetLine()

		field := model.JField{Name: name, Source: node.Pkg, StartLine: startLine, StopLine: stopLine}
		node.AddField(field)
	}
}

func (s *JavaRefactorListener) EnterMethodCall(ctx *MethodCallContext) {
	text := ctx.Identifier().GetText()
	startLine := ctx.GetStart().GetLine()
	stopLine := ctx.GetStop().GetLine()
	field := model.JField{Name: text, Source: node.Pkg, StartLine: startLine, StopLine: stopLine}
	node.AddField(field)
}

func (s *JavaRefactorListener) EnterExpressionList(ctx *ExpressionListContext) {
	for _, expression := range ctx.AllExpression() {
		expText := expression.GetText()
		if isUppercaseText(expText) {
			startLine := ctx.GetStart().GetLine()
			stopLine := ctx.GetStop().GetLine()
			field := model.JField{Name: expText, Source: node.Pkg, StartLine: startLine, StopLine: stopLine}
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
			field := model.JField{Name: expText, Source: node.Pkg, StartLine: startLine, StopLine: stopLine}
			node.AddField(field)
		}
	}
}

func (s *JavaRefactorListener) EnterCreatedName(ctx *CreatedNameContext) {
	identifiers := ctx.AllIdentifier()
	for index := range identifiers {
		context := ctx.Identifier(index)
		name := context.GetText()
		startLine := ctx.GetStart().GetLine()
		stopLine := ctx.GetStop().GetLine()

		field := model.JField{Name: name, Source: node.Pkg, StartLine: startLine, StopLine: stopLine}
		node.AddField(field)
	}
}

func (s *JavaRefactorListener) EnterExpression(ctx *ExpressionContext) {
	if ctx.Expression(0) != nil {
		expText := ctx.Expression(0).GetText()

		if isUppercaseText(expText) {
			startLine := ctx.GetStart().GetLine()
			stopLine := ctx.GetStop().GetLine()
			field := model.JField{Name: expText, Source: node.Pkg, StartLine: startLine, StopLine: stopLine}
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
				field := model.JField{Name: expText, Source: node.Pkg, StartLine: startLine, StopLine: stopLine}
				node.AddField(field)
			}
		}
	}
}

func isUppercaseText(text string) bool {
	return !strings.Contains(text, ".") && unicode.IsUpper([]rune(text)[0])
}

func (s *JavaRefactorListener) InitNode(identifier model.JFullIdentifier) {
	node = identifier
}

func (s *JavaRefactorListener) GetNodeInfo() model.JFullIdentifier {
	return node
}
