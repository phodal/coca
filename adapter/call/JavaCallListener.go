package call

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "github.com/phodal/coca/adapter/models"
	. "github.com/phodal/coca/language/java"
	"reflect"
	"strings"
)

var imports []string
var clzs []string
var currentPkg string
var currentClz string
var methods []JMethod
var methodCalls []JMethodCall
var currentType string

var fields = make(map[string]string)
var localVars = make(map[string]string)
var formalParameters = make(map[string]string)
var currentClzExtends = ""

var hasEnterClass = false
var isSpringRestController = false

func NewJavaCallListener() *JavaCallListener {
	currentClz = ""
	currentPkg = ""
	methods = nil
	methodCalls = nil
	return &JavaCallListener{}
}

type JavaCallListener struct {
	BaseJavaParserListener
}

func (s *JavaCallListener) getNodeInfo() *JClassNode {
	return &JClassNode{currentPkg, currentClz, currentType, "", methods, methodCalls}
}

func (s *JavaCallListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {
	currentPkg = ctx.QualifiedName().GetText()
}

func (s *JavaCallListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {
	importText := ctx.QualifiedName().GetText()
	imports = append(imports, importText)
}

func (s *JavaCallListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {
	hasEnterClass = true
	currentType = "Class"
	currentClz = ctx.IDENTIFIER().GetText()

	if ctx.EXTENDS() != nil {
		currentClzExtends = ctx.TypeType().GetText()
	}
}

func (s *JavaCallListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {
	currentType = "Interface"
	currentClz = ctx.IDENTIFIER().GetText()
}

func (s *JavaCallListener) EnterInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.IDENTIFIER().GetSymbol().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	name := ctx.IDENTIFIER().GetText()
	stopLinePosition := startLinePosition + len(name)

	typeType := ctx.TypeTypeOrVoid().GetText()

	method := &JMethod{name, typeType, startLine, startLinePosition, stopLine, stopLinePosition}
	methods = append(methods, *method)
}

func (s *JavaCallListener) EnterFormalParameter(ctx *FormalParameterContext) {
	formalParameters[ctx.VariableDeclaratorId().GetText()] = ctx.TypeType().GetText()
}

func (s *JavaCallListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {
	declarators := ctx.VariableDeclarators()
	variableName := declarators.GetParent().GetChild(0).(antlr.ParseTree).GetText()
	for _, declarator := range declarators.(*VariableDeclaratorsContext).AllVariableDeclarator() {
		value := declarator.(*VariableDeclaratorContext).VariableDeclaratorId().(*VariableDeclaratorIdContext).IDENTIFIER().GetText()
		fields[value] = variableName
	}
}

func (s *JavaCallListener) EnterLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {
	typ := ctx.GetChild(0).(antlr.ParseTree).GetText()
	variableName := ctx.GetChild(1).GetChild(0).GetChild(0).(antlr.ParseTree).GetText()
	localVars[variableName] = typ
}

func (s *JavaCallListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.IDENTIFIER().GetSymbol().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	name := ctx.IDENTIFIER().GetText()
	stopLinePosition := startLinePosition + len(name)

	typeType := ctx.TypeTypeOrVoid().GetText()

	method := &JMethod{name, typeType, startLine, startLinePosition, stopLine, stopLinePosition}
	methods = append(methods, *method)

	if ctx.FormalParameters() != nil {
		if ctx.FormalParameters().GetChild(0) == nil || ctx.FormalParameters().GetText() == "()" || ctx.FormalParameters().GetChild(1) == nil {
			return
		}

		parameterList := ctx.FormalParameters().GetChild(1).(*FormalParameterListContext)
		formalParameter := parameterList.AllFormalParameter()
		for _, param := range formalParameter {
			paramContext := param.(*FormalParameterContext)
			paramType := paramContext.TypeType().GetText()
			paramValue := paramContext.VariableDeclaratorId().(*VariableDeclaratorIdContext).IDENTIFIER().GetText()

			localVars[paramValue] = paramType
		}
	}
}

func (s *JavaCallListener) EnterCreator(ctx *CreatorContext) {
	variableName := ctx.GetParent().GetParent().GetChild(0).(antlr.ParseTree).GetText()
	localVars[variableName] = ctx.CreatedName().GetText()
}

func (s *JavaCallListener) EnterLocalTypeDeclaration(ctx *LocalTypeDeclarationContext) {

}

func (s *JavaCallListener) EnterMethodCall(ctx *MethodCallContext) {
	var targetCtx = ctx.GetParent().GetChild(0).(antlr.ParseTree).GetText()
	var targetType = parseTargetType(targetCtx)
	callee := ctx.GetChild(0).(antlr.ParseTree).GetText()

	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := startLinePosition + len(callee)

	// TODO: 处理链试调用
	if strings.Contains(targetType, "()") && strings.Contains(targetType, ".") {
		split := strings.Split(targetType, ".")
		sourceTarget := split[0]
		targetType = localVars[sourceTarget]
	}

	fullType := warpTargetFullType(targetType)
	if targetType == "super" {
		targetType = currentClzExtends
	}
	if fullType != "" {
		jMethodCall := &JMethodCall{removeTarget(fullType), "", targetType, callee, startLine, startLinePosition, stopLine, stopLinePosition}
		methodCalls = append(methodCalls, *jMethodCall)
	} else {
		if ctx.GetText() == targetType {
			methodName := ctx.IDENTIFIER().GetText()
			jMethodCall := &JMethodCall{currentPkg, "", currentClz, methodName, startLine, startLinePosition, stopLine, stopLinePosition}
			methodCalls = append(methodCalls, *jMethodCall)
		} else {
			methodName := ctx.IDENTIFIER().GetText()
			jMethodCall := &JMethodCall{currentPkg, "NEEDFIX", targetType, methodName, startLine, startLinePosition, stopLine, stopLinePosition}
			methodCalls = append(methodCalls, *jMethodCall)
		}
	}
}

var baseApiUrlName = ""

type RestApi struct {
	Uri            string
	Method         string
	ResponseStatus string
	Body           []string
}

func (s *JavaCallListener) EnterAnnotation(ctx *AnnotationContext) {
	annotationName := ctx.QualifiedName().GetText()
	if annotationName == "RestController" {
		isSpringRestController = true
	}

	if !hasEnterClass {
		if annotationName == "RequestMapping" {
			if ctx.ElementValuePairs() != nil {
				firstPair := ctx.ElementValuePairs().GetChild(0).(*ElementValuePairContext)
				if firstPair.IDENTIFIER().GetText() == "value" {
					baseApiUrlName = firstPair.ElementValue().GetText()
				}
			} else {
				baseApiUrlName = "/"
			}
		}
	}

	if !(annotationName == "GetMapping" || annotationName == "PutMapping" || annotationName == "PostMapping" || annotationName == "DeleteMapping") {
		return
	}

	uri := ""
	if ctx.ElementValue() != nil {
		uri = baseApiUrlName + ctx.ElementValue().GetText()
	} else {
		uri = baseApiUrlName
	}

	uriRemoveQuote := strings.ReplaceAll(uri, "\"", "")

	restApi := &RestApi{uriRemoveQuote, "", "", nil}
	if hasEnterClass {
		switch annotationName {
		case "GetMapping":
			restApi.Method = "GET"
		case "PutMapping":
			restApi.Method = "PUT"
		case "PostMapping":
			restApi.Method = "POST"
		case "DeleteMapping":
			restApi.Method = "DELETE"
		}
	}

	fmt.Println(restApi)
}

func (s *JavaCallListener) EnterExpression(ctx *ExpressionContext) {
	// lambda BlogPO::of
	if ctx.COLONCOLON() != nil {
		text := ctx.Expression(0).GetText()
		methodName := ctx.IDENTIFIER().GetText()
		targetType := parseTargetType(text)
		fullType := warpTargetFullType(targetType)

		startLine := ctx.GetStart().GetLine()
		startLinePosition := ctx.GetStart().GetColumn()
		stopLine := ctx.GetStop().GetLine()
		stopLinePosition := startLinePosition + len(text)

		jMethodCall := &JMethodCall{removeTarget(fullType), "", targetType, methodName, startLine, startLinePosition, stopLine, stopLinePosition}
		methodCalls = append(methodCalls, *jMethodCall)
	}
}

func (s *JavaCallListener) appendClasses(classes []string) {
	clzs = classes
}

func removeTarget(fullType string) string {
	split := strings.Split(fullType, ".")
	return strings.Join(split[:len(split)-1], ".")
}

func parseTargetType(targetCtx string) string {
	targetVar := targetCtx
	targetType := targetVar

	//TODO: update this reflect
	typeOf := reflect.TypeOf(targetCtx).String()
	if strings.HasSuffix(typeOf, "MethodCallContext") {
		targetType = currentClz
	} else {
		fieldType := fields[targetVar]
		formalType := formalParameters[targetVar]
		localVarType := localVars[targetVar]
		if fieldType != "" {
			targetType = fieldType
		} else if formalType != "" {
			targetType = formalType
		} else if localVarType != "" {
			targetType = localVarType
		}
	}

	return targetType
}

func warpTargetFullType(targetType string) string {
	if strings.EqualFold(currentClz, targetType) {
		return currentPkg + "." + targetType
	}

	// TODO: update for array
	split := strings.Split(targetType, ".")
	str := split[0]
	pureTargetType := strings.ReplaceAll(strings.ReplaceAll(str, "[", ""), "]", "")

	for index := range imports {
		imp := imports[index]
		if strings.HasSuffix(imp, pureTargetType) {
			return imp
		}
	}

	//maybe the same package
	for _, clz := range clzs {
		if strings.HasSuffix(clz, "."+pureTargetType) {
			return clz
		}
	}

	//1. current package, 2. import by *
	if pureTargetType == "super" {
		for index := range imports {
			imp := imports[index]
			if strings.HasSuffix(imp, currentClzExtends) {
				return imp
			}
		}
	}

	return ""
}
