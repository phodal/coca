package adapter

import (
	. "../language/java"
	. "./models"
	"fmt"
)


var imports []string
var currentPkg = ""
var currentClz = ""
var methodCalls []JMethodCall

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
	currentMethodCall := &JMethodCall{currentPkg, currentClz, methodName}
	methodCalls = append(methodCalls, *currentMethodCall)
}

func (s *JavaCallListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {
	methodName := ctx.IDENTIFIER().GetText()
	currentMethodCall := &JMethodCall{currentPkg, currentClz, methodName}
	methodCalls = append(methodCalls, *currentMethodCall)

	fmt.Println(methodCalls)
}
