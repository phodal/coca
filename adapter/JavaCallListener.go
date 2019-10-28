package adapter

import (
	. "../language/java"
	. "./models"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)


var imports []string
var currentPkg = ""
var currentClz = ""
var methodCalls []JMethodCall
var currentMethodCall *JMethodCall;

type JavaCallListener struct {
	BaseJavaParserListener
}

func (s *JavaCallListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {
	currentPkg = ctx.QualifiedName().GetText()
}

func (s *JavaCallListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {
	importText := ctx.QualifiedName().GetText()
	imports = append(imports, importText)
}


func (s *JavaCallListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {
	currentClz = ctx.IDENTIFIER().GetText()
}

func (s *JavaCallListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {
	currentClz = ctx.IDENTIFIER().GetText()
}

func (s *JavaCallListener) EnterInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
	methodName := ctx.IDENTIFIER().GetText()
	currentMethodCall = &JMethodCall{currentPkg, currentClz, methodName}
	methodCalls = append(methodCalls, *currentMethodCall)
}

func (s *JavaCallListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {
	methodName := ctx.IDENTIFIER().GetText()
	currentMethodCall = &JMethodCall{currentPkg, currentClz, methodName}
	methodCalls = append(methodCalls, *currentMethodCall)
}

func (s *JavaCallListener) EnterMethodCall(ctx *MethodCallContext) {
	if currentMethodCall != nil {
		var targetType = parseTargetType(ctx);
		callee := ctx.GetChild(0).(antlr.ParseTree).GetText()

		fmt.Println(targetType, callee)
	}
}


func parseTargetType(ctx *MethodCallContext) string {
	var targetCtx antlr.ParseTree = ctx.GetParent().GetChild(0).(antlr.ParseTree)
	targetVar := targetCtx.GetText();
	targetType := targetVar;

	return targetType
}
