package full

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/domain"
	common_listener2 "github.com/phodal/coca/core/infrastructure/ast/common_listener"
	"github.com/phodal/coca/languages/java"
	"reflect"
	"strconv"
	"strings"
)

var imports []string
var clzs []string
var currentPkg string
var currentClz string
var fields []domain.JAppField
var methodCalls []domain.JMethodCall
var currentType string

var mapFields = make(map[string]string)
var localVars = make(map[string]string)
var formalParameters = make(map[string]string)
var currentClzExtend = ""
var currentMethod domain.JMethod
var methodMap = make(map[string]domain.JMethod)
var creatorMethodMap = make(map[string]domain.JMethod)

var methodQueue []domain.JMethod
var classQueue []string

var identMap map[string]domain.JIdentifier
var isOverrideMethod = false

var classNodeQueue []domain.JClassNode

var currentNode *domain.JClassNode
var classNodes []domain.JClassNode
var creatorNodes []domain.JClassNode
var currentCreatorNode domain.JClassNode
var fileName = ""
var hasEnterClass = false

func NewJavaFullListener(nodes map[string]domain.JIdentifier, file string) *JavaFullListener {
	identMap = nodes
	imports = nil
	fileName = file
	currentPkg = ""
	classNodes = nil
	currentNode = domain.NewClassNode()
	classQueue = nil
	methodQueue = nil

	initClass()
	return &JavaFullListener{}
}

func initClass() {
	currentClz = ""
	currentClzExtend = ""
	currentMethod = domain.NewJMethod()
	currentNode.MethodCalls = nil

	methodMap = make(map[string]domain.JMethod)
	methodCalls = nil
	fields = nil
	isOverrideMethod = false
}

type JavaFullListener struct {
	parser.BaseJavaParserListener
}

func (s *JavaFullListener) GetNodeInfo() []domain.JClassNode {
	return classNodes
}

func (s *JavaFullListener) ExitClassBody(ctx *parser.ClassBodyContext) {
	hasEnterClass = false
	s.exitBody()
}

func (s *JavaFullListener) ExitInterfaceBody(ctx *parser.InterfaceBodyContext) {
	hasEnterClass = false
	s.exitBody()
}

func (s *JavaFullListener) exitBody() {
	if currentNode.Class != "" {
		currentNode.Fields = fields
		currentNode.Type = currentType
		currentNode.Path = fileName
		currentNode.SetMethodFromMap(methodMap)
	}

	if currentType == "CreatorClass" {
		currentNode.SetMethodFromMap(creatorMethodMap)
		return
	}

	if currentNode.Class != "" {
		classNodes = append(classNodes, *currentNode)
	}

	currentNode = domain.NewClassNode()
	initClass()
}

func (s *JavaFullListener) EnterPackageDeclaration(ctx *parser.PackageDeclarationContext) {
	currentNode.Package = ctx.QualifiedName().GetText()
	currentPkg = ctx.QualifiedName().GetText()
}

func (s *JavaFullListener) EnterImportDeclaration(ctx *parser.ImportDeclarationContext) {
	importText := ctx.QualifiedName().GetText()
	imports = append(imports, importText)
}

func (s *JavaFullListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	// TODO: support inner class
	if currentNode.Class != "" {
		return
	}

	hasEnterClass = true
	currentClzExtend = ""
	currentType = "Class"
	if ctx.IDENTIFIER() != nil {
		currentClz = ctx.IDENTIFIER().GetText()
		currentNode.Class = currentClz
	}

	if ctx.EXTENDS() != nil {
		currentClzExtend = ctx.TypeType().GetText()
		buildExtend(currentClzExtend)
	}

	if ctx.IMPLEMENTS() != nil {
		types := ctx.TypeList().(*parser.TypeListContext).AllTypeType()
		for _, typ := range types {
			typeText := typ.GetText()
			buildImplement(typeText)
		}
	}

	// TODO: 支持依赖注入
}

func (s *JavaFullListener) EnterInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	hasEnterClass = true
	currentType = "Interface"
	currentNode.Class = ctx.IDENTIFIER().GetText()

	if ctx.EXTENDS() != nil {
		types := ctx.TypeList().(*parser.TypeListContext).AllTypeType()
		for _, typ := range types {
			buildExtend(typ.GetText())
		}
	}
}

func (s *JavaFullListener) EnterInterfaceBodyDeclaration(ctx *parser.InterfaceBodyDeclarationContext) {
	hasEnterClass = true
	for _, modifier := range ctx.AllModifier() {
		modifier := modifier.(*parser.ModifierContext).GetChild(0)
		if reflect.TypeOf(modifier.GetChild(0)).String() == "*parser.AnnotationContext" {
			annotationContext := modifier.GetChild(0).(*parser.AnnotationContext)
			common_listener2.BuildAnnotation(annotationContext)
		}
	}
}

func (s *JavaFullListener) EnterInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.IDENTIFIER().GetSymbol().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	name := ctx.IDENTIFIER().GetText()
	stopLinePosition := startLinePosition + len(name)

	typeType := ctx.TypeTypeOrVoid().GetText()

	if reflect.TypeOf(ctx.GetParent().GetParent().GetChild(0)).String() == "*parser.ModifierContext" {
		common_listener2.BuildAnnotationForMethod(ctx.GetParent().GetParent().GetChild(0).(*parser.ModifierContext), &currentMethod)
	}

	method := &domain.JMethod{Name: name, Type: typeType, StartLine: startLine, StartLinePosition: startLinePosition, StopLine: stopLine, StopLinePosition: stopLinePosition}
	updateMethod(method)
}

func (s *JavaFullListener) ExitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {

}

func (s *JavaFullListener) EnterFormalParameter(ctx *parser.FormalParameterContext) {
	formalParameters[ctx.VariableDeclaratorId().GetText()] = ctx.TypeType().GetText()
}

func (s *JavaFullListener) EnterFieldDeclaration(ctx *parser.FieldDeclarationContext) {
	decelerators := ctx.VariableDeclarators()
	typeType := decelerators.GetParent().GetChild(0).(*parser.TypeTypeContext)
	for _, declarator := range decelerators.(*parser.VariableDeclaratorsContext).AllVariableDeclarator() {
		var typeCtx *parser.ClassOrInterfaceTypeContext = nil
		if reflect.TypeOf(typeType.GetChild(0)).String() == "*parser.ClassOrInterfaceTypeContext" {
			typeCtx = typeType.GetChild(0).(*parser.ClassOrInterfaceTypeContext)
		}

		if typeType.GetChildCount() > 1 {
			if reflect.TypeOf(typeType.GetChild(1)).String() == "*parser.ClassOrInterfaceTypeContext" {
				typeCtx = typeType.GetChild(1).(*parser.ClassOrInterfaceTypeContext)
			}
		}

		if typeCtx == nil {
			continue
		}

		typeTypeText := typeCtx.IDENTIFIER(0).GetText()
		value := declarator.(*parser.VariableDeclaratorContext).VariableDeclaratorId().(*parser.VariableDeclaratorIdContext).IDENTIFIER().GetText()
		mapFields[value] = typeTypeText
		fields = append(fields, *&domain.JAppField{Type: typeTypeText, Value: value})

		buildFieldCall(typeTypeText, ctx)
	}
}

func (s *JavaFullListener) EnterLocalVariableDeclaration(ctx *parser.LocalVariableDeclarationContext) {
	typ := ctx.GetChild(0).(antlr.ParseTree).GetText()
	if ctx.GetChild(1) != nil {
		if ctx.GetChild(1).GetChild(0) != nil && ctx.GetChild(1).GetChild(0).GetChild(0) != nil {
			variableName := ctx.GetChild(1).GetChild(0).GetChild(0).(antlr.ParseTree).GetText()
			localVars[variableName] = typ
		}
	}
}

func (s *JavaFullListener) EnterAnnotation(ctx *parser.AnnotationContext) {
	// Todo: support override method
	annotationName := ctx.QualifiedName().GetText()
	if annotationName == "Override" {
		isOverrideMethod = true
	} else {
		isOverrideMethod = false
	}

	if !hasEnterClass {
		annotation := common_listener2.BuildAnnotation(ctx)
		if currentType == "CreatorClass" {
			currentCreatorNode.Annotations = append(currentCreatorNode.Annotations, annotation)
		} else {
			currentNode.Annotations = append(currentNode.Annotations, annotation)
		}
	}
}

func (s *JavaFullListener) EnterConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) {
	method := &domain.JMethod{
		Name:              ctx.IDENTIFIER().GetText(),
		Type:              "",
		StartLine:         ctx.GetStart().GetLine(),
		StartLinePosition: ctx.GetStart().GetColumn(),
		StopLine:          ctx.GetStop().GetLine(),
		StopLinePosition:  ctx.GetStop().GetColumn(),
		Override:          isOverrideMethod,
		Parameters:        nil,
		Annotations:       currentMethod.Annotations,
		IsConstructor:     true,
	}

	parameters := ctx.FormalParameters()
	if buildMethodParameters(parameters, method) {
		return
	}

	updateMethod(method)
}

func (s *JavaFullListener) ExitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) {
	currentMethod = domain.NewJMethod()
	isOverrideMethod = false
}

func (s *JavaFullListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.IDENTIFIER().GetSymbol().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	name := ctx.IDENTIFIER().GetText()
	stopLinePosition := startLinePosition + len(name)

	typeType := ctx.TypeTypeOrVoid().GetText()

	if reflect.TypeOf(ctx.GetParent().GetParent().GetChild(0)).String() == "*parser.ModifierContext" {
		common_listener2.BuildAnnotationForMethod(ctx.GetParent().GetParent().GetChild(0).(*parser.ModifierContext), &currentMethod)
	}

	method := &domain.JMethod{
		Name:              name,
		Type:              typeType,
		StartLine:         startLine,
		StartLinePosition: startLinePosition,
		StopLine:          stopLine,
		StopLinePosition:  stopLinePosition,
		Annotations:       currentMethod.Annotations,
		Override:          isOverrideMethod,
		Parameters:        nil,
		Creators:          nil,
	}

	parameters := ctx.FormalParameters()
	if buildMethodParameters(parameters, method) {
		return
	}

	updateMethod(method)
}

func buildMethodParameters(parameters parser.IFormalParametersContext, method *domain.JMethod) bool {
	if parameters != nil {
		if parameters.GetChild(0) == nil || parameters.GetText() == "()" || parameters.GetChild(1) == nil {
			updateMethod(method)
			return true
		}

		methodParams := getMethodParameters(parameters)

		method.Parameters = methodParams
		updateMethod(method)
	}
	return false
}

func getMethodParameters(parameters parser.IFormalParametersContext) []domain.JParameter {
	var methodParams []domain.JParameter = nil
	parameterList := parameters.GetChild(1).(*parser.FormalParameterListContext)
	formalParameter := parameterList.AllFormalParameter()
	for _, param := range formalParameter {
		paramContext := param.(*parser.FormalParameterContext)
		paramType := paramContext.TypeType().GetText()
		paramValue := paramContext.VariableDeclaratorId().(*parser.VariableDeclaratorIdContext).IDENTIFIER().GetText()

		localVars[paramValue] = paramType
		methodParams = append(methodParams, *&domain.JParameter{Name: paramValue, Type: paramType})
	}
	return methodParams
}

func updateMethod(method *domain.JMethod) {
	if currentType == "CreatorClass" {
		creatorMethodMap[getMethodMapName(*method)] = *method
	} else {
		currentMethod = *method
		methodQueue = append(methodQueue, *method)
		methodMap[getMethodMapName(*method)] = *method
	}
}

func (s *JavaFullListener) ExitMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	exitMethod()
}

func exitMethod() {
	if currentType == "CreatorClass" {
		return
	}

	currentMethod = domain.NewJMethod()
}

// TODO: add inner creator examples
func (s *JavaFullListener) EnterInnerCreator(ctx *parser.InnerCreatorContext) {
	if ctx.IDENTIFIER() != nil {
		currentClz = ctx.IDENTIFIER().GetText()
		classQueue = append(classQueue, currentClz)
	}
}

// TODO: add inner creator examples
func (s *JavaFullListener) ExitInnerCreator(ctx *parser.InnerCreatorContext) {
	if classQueue == nil || len(classQueue) <= 1 {
		return
	}

	classQueue = classQueue[0 : len(classQueue)-1]
	currentClz = classQueue[len(classQueue)-1]
}

func getMethodMapName(method domain.JMethod) string {
	name := method.Name
	if name == "" && len(methodQueue) > 1 {
		name = methodQueue[len(methodQueue)-1].Name
	}
	return currentPkg + "." + currentClz + "." + name + ":" + strconv.Itoa(method.StartLine)
}

func (s *JavaFullListener) EnterCreator(ctx *parser.CreatorContext) {
	variableName := ctx.GetParent().GetParent().GetChild(0).(antlr.ParseTree).GetText()
	allIdentifiers := ctx.CreatedName().(*parser.CreatedNameContext).AllIDENTIFIER()

	for _, identifier := range allIdentifiers {
		createdName := identifier.GetText()
		localVars[variableName] = createdName

		classNodeQueue = append(classNodeQueue, *currentNode)
		buildCreatedCall(createdName, ctx)

		if currentMethod.Name == "" {
			return
		}

		if ctx.ClassCreatorRest() == nil {
			return
		}

		if ctx.ClassCreatorRest().(*parser.ClassCreatorRestContext).ClassBody() == nil {
			return
		}

		currentType = "CreatorClass"
		text := ctx.CreatedName().GetText()
		creatorNode := &domain.JClassNode{
			Package:     currentPkg,
			Class:       text,
			Type:        "CreatorClass",
			Path:        "",
			Fields:      nil,
			Methods:     nil,
			MethodCalls: nil,
			Extend:      "",
			Implements:  nil,
			Annotations: nil,
		}

		currentCreatorNode = *creatorNode
		creatorNodes = append(creatorNodes, *creatorNode)
	}
}

func (s *JavaFullListener) ExitCreator(ctx *parser.CreatorContext) {
	if currentCreatorNode.Class != "" {
		method := methodMap[getMethodMapName(currentMethod)]
		method.Creators = append(method.Creators, currentCreatorNode)
		methodMap[getMethodMapName(currentMethod)] = method
	}

	currentType = ""
	currentCreatorNode = *domain.NewClassNode()

	if classNodeQueue == nil || len(classNodeQueue) < 1 {
		return
	}
}

func buildCreatedCall(createdName string, ctx *parser.CreatorContext) {
	method := methodMap[getMethodMapName(currentMethod)]
	fullType, _ := warpTargetFullType(createdName)

	jMethodCall := &domain.JMethodCall{
		Package:           removeTarget(fullType),
		Type:              "CreatorClass",
		Class:             createdName,
		MethodName:        "",
		StartLine:         ctx.GetStart().GetLine(),
		StartLinePosition: ctx.GetStart().GetColumn(),
		StopLine:          ctx.GetStop().GetLine(),
		StopLinePosition:  ctx.GetStop().GetColumn(),
	}

	method.MethodCalls = append(method.MethodCalls, *jMethodCall)
	methodMap[getMethodMapName(currentMethod)] = method
}

func (s *JavaFullListener) EnterMethodCall(ctx *parser.MethodCallContext) {
	var jMethodCall = domain.NewJMethodCall()

	targetCtx := ctx.GetParent().GetChild(0).(antlr.ParseTree)
	var targetType = parseTargetType(targetCtx.GetText())

	if targetCtx.GetChild(0) != nil {
		if reflect.TypeOf(targetCtx.GetChild(0)).String() == "*parser.MethodCallContext" {
			methodCallContext := targetCtx.GetChild(0).(*parser.MethodCallContext)
			targetType = methodCallContext.IDENTIFIER().GetText()
		}
	}

	callee := ctx.GetChild(0).(antlr.ParseTree).GetText()

	jMethodCall.StartLine = ctx.GetStart().GetLine()
	jMethodCall.StartLinePosition = ctx.GetStart().GetColumn()
	jMethodCall.StopLine = ctx.GetStop().GetLine()
	jMethodCall.StopLinePosition = jMethodCall.StartLinePosition + len(callee)

	fullType, callType := warpTargetFullType(targetType)
	if targetType == "super" || callee == "super" {
		callType = "super"
		targetType = currentClzExtend
	}
	jMethodCall.Type = callType

	methodName := callee
	packageName := currentPkg

	if fullType != "" {
		packageName = removeTarget(fullType)
		methodName = callee
	} else {
		if ctx.GetText() == targetType {
			clz := currentClz
			// 处理 static 方法，如 now()
			for _, imp := range imports {
				if strings.HasSuffix(imp, "."+methodName) {
					packageName = imp
					clz = ""
				}
			}

			targetType = clz
		} else {
			targetType = buildSelfThisTarget(targetType)
			targetType = buildMethodNameForBuilder(ctx, targetType)
		}
	}

	jMethodCall.Package = packageName
	jMethodCall.MethodName = methodName
	jMethodCall.Class = targetType

	if ctx.ExpressionList() != nil {
		var parameters []string
		for _, expression := range ctx.ExpressionList().(*parser.ExpressionListContext).AllExpression() {
			expressionCtx := expression.(*parser.ExpressionContext)
			parameters = append(parameters, expressionCtx.GetText())
		}
		jMethodCall.Parameters = parameters
	}

	addMethodCall(jMethodCall)
}

func addMethodCall(jMethodCall domain.JMethodCall) {
	methodCalls = append(methodCalls, jMethodCall)

	method := methodMap[getMethodMapName(currentMethod)]
	method.MethodCalls = append(method.MethodCalls, jMethodCall)
	methodMap[getMethodMapName(currentMethod)] = method
}

func isChainCall(targetType string) bool {
	return strings.Contains(targetType, "(") && strings.Contains(targetType, ")") && strings.Contains(targetType, ".")
}

func buildMethodNameForBuilder(ctx *parser.MethodCallContext, targetType string) string {
	// TODO: refactor
	if reflect.TypeOf(ctx.GetParent()).String() == "*parser.ExpressionContext" {
		parentCtx := ctx.GetParent().(*parser.ExpressionContext)
		if reflect.TypeOf(parentCtx.GetParent()).String() == "*parser.VariableInitializerContext" {
			varParent := parentCtx.GetParent().(*parser.VariableInitializerContext).GetParent()
			if reflect.TypeOf(varParent).String() == "*parser.VariableDeclaratorContext" {
				varDeclParent := varParent.(*parser.VariableDeclaratorContext).GetParent()
				if reflect.TypeOf(varDeclParent).String() == "*parser.VariableDeclaratorsContext" {
					parent := varDeclParent.(*parser.VariableDeclaratorsContext).GetParent()
					if reflect.TypeOf(parent).String() == "*parser.LocalVariableDeclarationContext" {
						context := parent.(*parser.LocalVariableDeclarationContext)
						targetType = context.TypeType().GetText()
					}
				}
			}
		}
	}

	return targetType
}

func buildSelfThisTarget(targetType string) string {
	isSelfFieldCall := strings.Contains(targetType, "this.")
	if isSelfFieldCall {
		targetType = strings.ReplaceAll(targetType, "this.", "")
		for _, field := range fields {
			if field.Value == targetType {
				targetType = field.Type
			}
		}
	}
	return targetType
}

func (s *JavaFullListener) EnterExpression(ctx *parser.ExpressionContext) {
	// lambda BlogPO::of
	if ctx.COLONCOLON() != nil {
		if ctx.Expression(0) == nil {
			return
		}

		text := ctx.Expression(0).GetText()
		methodName := ctx.IDENTIFIER().GetText()
		targetType := parseTargetType(text)

		fullType, _ := warpTargetFullType(targetType)

		startLine := ctx.GetStart().GetLine()
		startLinePosition := ctx.GetStart().GetColumn()
		stopLine := ctx.GetStop().GetLine()
		stopLinePosition := startLinePosition + len(text)

		jMethodCall := &domain.JMethodCall{
			Package:           removeTarget(fullType),
			Type:              "lambda",
			Class:             targetType,
			MethodName:        methodName,
			Parameters:        nil,
			StartLine:         startLine,
			StartLinePosition: startLinePosition,
			StopLine:          stopLine,
			StopLinePosition:  stopLinePosition,
		}
		addMethodCall(*jMethodCall)
	}
}

func (s *JavaFullListener) AppendClasses(classes []string) {
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
		//if isChainCall(targetVar) {
		//	split := strings.Split(targetType, ".")
		//	targetVar = split[0]
		//}

		fieldType := mapFields[targetVar]
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

func warpTargetFullType(targetType string) (string, string) {
	callType := ""
	if strings.EqualFold(currentClz, targetType) {
		callType = "self"
		return currentPkg + "." + targetType, ""
	}

	// TODO: update for array
	split := strings.Split(targetType, ".")
	str := split[0]
	pureTargetType := strings.ReplaceAll(strings.ReplaceAll(str, "[", ""), "]", "")

	if pureTargetType != "" {
		for _, imp := range imports {
			if strings.HasSuffix(imp, pureTargetType) {
				callType = "chain"
				return imp, callType
			}
		}
	}

	//maybe the same package
	for _, clz := range clzs {
		if strings.HasSuffix(clz, "."+pureTargetType) {
			callType = "same package"
			return clz, callType
		}
	}

	//1. current package, 2. import by *
	if pureTargetType == "super" || pureTargetType == "this" {
		for _, imp := range imports {
			if strings.HasSuffix(imp, currentClzExtend) {
				callType = pureTargetType
				return imp, callType
			}
		}
	}

	if _, ok := identMap[currentPkg+"."+targetType]; ok {
		callType = "same package 2"
		return currentPkg + "." + targetType, callType
	}

	return "", callType
}

func buildExtend(extendName string) {
	target, _ := warpTargetFullType(extendName)
	if target != "" {
		currentNode.Extend = target
	}
}

func buildFieldCall(typeType string, ctx *parser.FieldDeclarationContext) {
	target, _ := warpTargetFullType(typeType)
	if target != "" {
		jMethodCall := &domain.JMethodCall{
			Package:           removeTarget(target),
			Type:              "field",
			Class:             typeType,
			MethodName:        "",
			StartLine:         ctx.GetStart().GetLine(),
			StartLinePosition: ctx.GetStart().GetColumn(),
			StopLine:          ctx.GetStop().GetLine(),
			StopLinePosition:  ctx.GetStop().GetColumn(),
		}

		currentNode.MethodCalls = append(currentNode.MethodCalls, *jMethodCall)
	}
}

func buildImplement(text string) {
	target, _ := warpTargetFullType(text)
	currentNode.Implements = append(currentNode.Implements, target)
}
