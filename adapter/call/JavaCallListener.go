package call

import (
	. "../../language/java"
	. "./models"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"strings"
)

var imports []string
var currentPkg string
var currentClz string
var methodCalls []JMethodCall
var currentMethodCall *JMethodCall

var fields = make(map[string]string)
var localVars = make(map[string]string)
var formalParameters = make(map[string]string)

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

func (s *JavaCallListener) EnterFormalParameter(ctx *FormalParameterContext) {
	formalParameters[ctx.VariableDeclaratorId().GetText()] = ctx.TypeType().GetText()
}

func (s *JavaCallListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {
	declarators := ctx.VariableDeclarators()
	variableName := declarators.GetParent().GetChild(0).(antlr.ParseTree).GetText()
	fields[variableName] = ctx.TypeType().GetText()
}

func (s *JavaCallListener) EnterLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {
	typ := ctx.GetChild(0).(antlr.ParseTree).GetText()
	variableName := ctx.GetChild(1).GetChild(0).GetChild(0).(antlr.ParseTree).GetText()
	localVars[variableName] = typ
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

		fullType := warpTargetFullType(targetType);

		if fullType != "" {
			currentMethodCall.AddMethodCall(fullType, callee);
		} else {

		}

	}
}

func parseTargetType(ctx *MethodCallContext) string {
	var targetCtx = ctx.GetParent().GetChild(0).(antlr.ParseTree)
	targetVar := targetCtx.GetText();
	targetType := targetVar;

	//TODO: update this reflect
	typeOf := reflect.TypeOf(targetCtx).String()
	if strings.HasSuffix(typeOf, "MethodCallContext") {
		targetType = currentClz;
	} else {
		fieldType := fields[targetVar]
		formalType := formalParameters[targetVar]
		localVarType := localVars[targetVar]
		if fieldType != "" {
			targetType = fieldType
		} else if formalType != "" {
			targetType = formalType;
		} else if localVarType != "" {
			targetType = localVarType;
		}
	}

	return targetType
}

func warpTargetFullType(targetType string) string {
	if strings.EqualFold(currentClz, targetType) {
		return currentPkg + "." + targetType;
	}
	for index := range imports {
		imp := imports[index]
		if strings.HasSuffix(imp, targetType) {
			return imp
		}
	}

	return ""
}
