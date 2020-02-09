package ast_java

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/languages/java"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/ast_java/common_listener"
	"reflect"
	"strconv"
	"strings"
)

var imports []string
var clzs []string
var currentPkg string
var currentClz string
var fields []core_domain.CodeField
var methodCalls []core_domain.CodeCall
var currentType string

var mapFields = make(map[string]string)
var localVars = make(map[string]string)
var formalParameters = make(map[string]string)

var currentClzExtend = ""
var currentMethod core_domain.CodeFunction
var methodMap = make(map[string]core_domain.CodeFunction)
var creatorMethodMap = make(map[string]core_domain.CodeFunction)

var methodQueue []core_domain.CodeFunction
var classStringQueue []string

var identMap map[string]core_domain.CodeDataStruct
var isOverrideMethod = false

var classNodeQueue []core_domain.CodeDataStruct

var currentNode *core_domain.CodeDataStruct
var classNodes []core_domain.CodeDataStruct
var currentCreatorNode core_domain.CodeDataStruct
var fileName = ""
var hasEnterClass = false

func NewJavaFullListener(nodes map[string]core_domain.CodeDataStruct, file string) *JavaFullListener {
	identMap = nodes
	imports = nil
	fileName = file
	currentPkg = ""
	classNodes = nil
	currentNode = core_domain.NewDataStruct()
	classStringQueue = nil
	classNodeQueue = nil
	methodQueue = nil

	initClass()
	return &JavaFullListener{}
}

func initClass() {
	currentClz = ""
	currentClzExtend = ""
	currentMethod = core_domain.NewJMethod()
	currentNode.FunctionCalls = nil

	methodMap = make(map[string]core_domain.CodeFunction)
	methodCalls = nil
	fields = nil
	isOverrideMethod = false
}

type JavaFullListener struct {
	parser.BaseJavaParserListener
}

func (s *JavaFullListener) GetNodeInfo() []core_domain.CodeDataStruct {
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
	if currentNode.NodeName != "" {
		currentNode.Fields = fields
		currentNode.FilePath = fileName
		currentNode.SetMethodFromMap(methodMap)
	}

	if currentType == "CreatorClass" {
		currentNode.SetMethodFromMap(creatorMethodMap)
		return
	}

	if currentNode.NodeName == "" {
		currentNode = core_domain.NewDataStruct()
		initClass()
		return
	}

	if currentNode.Type == "InnerStructures" && len(classNodeQueue) >= 1 {
		classNodeQueue[0].InnerStructures = append(currentNode.InnerStructures, *currentNode)
	} else {
		classNodes = append(classNodes, *currentNode)
	}

	if len(classNodeQueue) >= 1 {
		if len(classNodeQueue) == 1 {
			currentNode = &classNodeQueue[0]
		} else {
			classNodeQueue = classNodeQueue[0 : len(classNodeQueue)-1]
			currentNode = &classNodeQueue[len(classNodeQueue)-1]
		}
	} else {
		currentNode = core_domain.NewDataStruct()
	}

	initClass()
}

func (s *JavaFullListener) EnterPackageDeclaration(ctx *parser.PackageDeclarationContext) {
	currentNode.Package = ctx.QualifiedName().GetText()
	currentPkg = ctx.QualifiedName().GetText()
}

func (s *JavaFullListener) EnterImportDeclaration(ctx *parser.ImportDeclarationContext) {
	importText := ctx.QualifiedName().GetText()
	imports = append(imports, importText)
	currentNode.Imports = append(currentNode.Imports, core_domain.NewJImport(importText))
}

func (s *JavaFullListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	if currentNode.NodeName != "" {
		classNodeQueue = append(classNodeQueue, *currentNode)
		currentType = "InnerStructures"
	} else {
		currentType = "Class"
	}

	hasEnterClass = true
	currentClzExtend = ""
	if ctx.IDENTIFIER() != nil {
		currentClz = ctx.IDENTIFIER().GetText()
		currentNode.NodeName = currentClz
	}

	if ctx.EXTENDS() != nil {
		currentClzExtend = ctx.TypeType().GetText()
		currentNode.Extend = buildExtend(currentClzExtend)
	}

	if ctx.IMPLEMENTS() != nil {
		types := ctx.TypeList().(*parser.TypeListContext).AllTypeType()
		for _, typ := range types {
			currentNode.Implements = append(currentNode.Implements, buildImplement(typ.GetText()))
		}
	}

	currentNode.Type = currentType
	// TODO: 支持依赖注入
}

func (s *JavaFullListener) EnterInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	hasEnterClass = true
	currentType = "Interface"
	currentNode.NodeName = ctx.IDENTIFIER().GetText()

	if ctx.EXTENDS() != nil {
		types := ctx.TypeList().(*parser.TypeListContext).AllTypeType()
		for _, typ := range types {
			currentNode.Extend = buildExtend(typ.GetText())
		}
	}

	currentNode.Type = currentType
}

func (s *JavaFullListener) EnterInterfaceBodyDeclaration(ctx *parser.InterfaceBodyDeclarationContext) {
	hasEnterClass = true
	for _, modifier := range ctx.AllModifier() {
		modifier := modifier.(*parser.ModifierContext).GetChild(0)
		if reflect.TypeOf(modifier.GetChild(0)).String() == "*parser.AnnotationContext" {
			annotationContext := modifier.GetChild(0).(*parser.AnnotationContext)
			common_listener.BuildAnnotation(annotationContext)
		}
	}
}

func (s *JavaFullListener) EnterInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) {
	name := ctx.IDENTIFIER().GetText()
	typeType := ctx.TypeTypeOrVoid().GetText()

	if reflect.TypeOf(ctx.GetParent().GetParent().GetChild(0)).String() == "*parser.ModifierContext" {
		common_listener.BuildAnnotationForMethod(ctx.GetParent().GetParent().GetChild(0).(*parser.ModifierContext), &currentMethod)
	}

	position := BuildPosition(ctx.BaseParserRuleContext, name)

	method := &core_domain.CodeFunction{Name: name, ReturnType: typeType, Position: position}
	updateMethod(method)
}

func (s *JavaFullListener) EnterFormalParameter(ctx *parser.FormalParameterContext) {
	formalParameters[ctx.VariableDeclaratorId().GetText()] = ctx.TypeType().GetText()
}

func (s *JavaFullListener) EnterFieldDeclaration(ctx *parser.FieldDeclarationContext) {
	declarators := ctx.VariableDeclarators()
	typeType := declarators.GetParent().GetChild(0).(*parser.TypeTypeContext)
	for _, declarator := range declarators.(*parser.VariableDeclaratorsContext).AllVariableDeclarator() {
		var typeCtx = BuildTypeCtxByIndex(typeType, nil, 0)
		if typeType.GetChildCount() > 1 {
			typeCtx = BuildTypeCtxByIndex(typeType, typeCtx, 1)
		}

		if typeCtx == nil {
			continue
		}

		typeTypeText := typeCtx.IDENTIFIER(0).GetText()
		value := declarator.(*parser.VariableDeclaratorContext).VariableDeclaratorId().(*parser.VariableDeclaratorIdContext).IDENTIFIER().GetText()
		mapFields[value] = typeTypeText
		field := core_domain.NewJField(typeTypeText, value, "")
		fields = append(fields, field)

		buildFieldCall(typeTypeText, ctx)
	}
}

func BuildTypeCtxByIndex(typeType *parser.TypeTypeContext, typeCtx *parser.ClassOrInterfaceTypeContext, index int) *parser.ClassOrInterfaceTypeContext {
	switch x := typeType.GetChild(index).(type) {
	case *parser.ClassOrInterfaceTypeContext:
		typeCtx = x
	}
	return typeCtx
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
		annotation := common_listener.BuildAnnotation(ctx)
		if currentType == "CreatorClass" {
			currentCreatorNode.Annotations = append(currentCreatorNode.Annotations, annotation)
		} else {
			currentNode.Annotations = append(currentNode.Annotations, annotation)
		}
	}
}

func (s *JavaFullListener) EnterConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) {
	name := ctx.IDENTIFIER().GetText()
	position := BuildPosition(ctx.BaseParserRuleContext, name)

	method := &core_domain.CodeFunction{
		Name:          name,
		ReturnType:    "",
		Override:      isOverrideMethod,
		Parameters:    nil,
		Annotations:   currentMethod.Annotations,
		IsConstructor: true,
		Position:      position,
	}

	parameters := ctx.FormalParameters()
	if buildMethodParameters(parameters, method) {
		return
	}

	updateMethod(method)
}

func (s *JavaFullListener) ExitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) {
	currentMethod = core_domain.NewJMethod()
	isOverrideMethod = false
}

func (s *JavaFullListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	name := ctx.IDENTIFIER().GetText()
	typeType := ctx.TypeTypeOrVoid().GetText()

	if reflect.TypeOf(ctx.GetParent().GetParent().GetChild(0)).String() == "*parser.ModifierContext" {
		common_listener.BuildAnnotationForMethod(ctx.GetParent().GetParent().GetChild(0).(*parser.ModifierContext), &currentMethod)
	}

	// check, before your refactor
	position := core_domain.CodePosition{
		StartLine:         ctx.GetStart().GetLine(),
		StartLinePosition: ctx.IDENTIFIER().GetSymbol().GetColumn(), // different
		StopLine:          ctx.GetStop().GetLine(),
		StopLinePosition:  ctx.IDENTIFIER().GetSymbol().GetColumn() + len(name),
	}

	method := &core_domain.CodeFunction{
		Name:            name,
		ReturnType:      typeType,
		Annotations:     currentMethod.Annotations,
		Override:        isOverrideMethod,
		Parameters:      nil,
		InnerStructures: nil,
		Position:        position,
	}

	parameters := ctx.FormalParameters()
	if buildMethodParameters(parameters, method) {
		return
	}

	updateMethod(method)
}

func buildMethodParameters(parameters parser.IFormalParametersContext, method *core_domain.CodeFunction) bool {
	if parameters != nil {
		if parameters.GetChild(0) == nil || parameters.GetText() == "()" || parameters.GetChild(1) == nil {
			updateMethod(method)
			return true
		}

		method.Parameters = BuildMethodParameters(parameters)
		updateMethod(method)
	}
	return false
}

func updateMethod(method *core_domain.CodeFunction) {
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

	currentMethod = core_domain.NewJMethod()
}

// TODO: add inner creator examples
func (s *JavaFullListener) EnterInnerCreator(ctx *parser.InnerCreatorContext) {
	if ctx.IDENTIFIER() != nil {
		currentClz = ctx.IDENTIFIER().GetText()
		classStringQueue = append(classStringQueue, currentClz)
	}
}

// TODO: add inner creator examples
func (s *JavaFullListener) ExitInnerCreator(ctx *parser.InnerCreatorContext) {
	if classStringQueue == nil || len(classStringQueue) <= 1 {
		return
	}

	classStringQueue = classStringQueue[0 : len(classStringQueue)-1]
	currentClz = classStringQueue[len(classStringQueue)-1]
}

func getMethodMapName(method core_domain.CodeFunction) string {
	name := method.Name
	if name == "" && len(methodQueue) > 1 {
		name = methodQueue[len(methodQueue)-1].Name
	}
	return currentPkg + "." + currentClz + "." + name + ":" + strconv.Itoa(method.Position.StartLine)
}

func (s *JavaFullListener) EnterCreator(ctx *parser.CreatorContext) {
	variableName := ctx.GetParent().GetParent().GetChild(0).(antlr.ParseTree).GetText()
	allIdentifiers := ctx.CreatedName().(*parser.CreatedNameContext).AllIDENTIFIER()

	for _, identifier := range allIdentifiers {
		createdName := identifier.GetText()
		localVars[variableName] = createdName

		buildCreatorCall(createdName, ctx)

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
		creatorNode := &core_domain.CodeDataStruct{
			Package:       currentPkg,
			NodeName:      text,
			Type:          "CreatorClass",
			FilePath:      "",
			Fields:        nil,
			Functions:     nil,
			FunctionCalls: nil,
			Extend:        "",
			Implements:    nil,
			Annotations:   nil,
		}

		currentCreatorNode = *creatorNode
	}
}

func (s *JavaFullListener) ExitCreator(ctx *parser.CreatorContext) {
	if currentCreatorNode.NodeName != "" {
		method := methodMap[getMethodMapName(currentMethod)]
		method.InnerStructures = append(method.InnerStructures, currentCreatorNode)
		methodMap[getMethodMapName(currentMethod)] = method
	}

	if currentType == "CreatorClass" {
		currentType = ""
	}
	currentCreatorNode = *core_domain.NewDataStruct()

	if classNodeQueue == nil || len(classNodeQueue) < 1 {
		return
	}
}

func buildCreatorCall(createdName string, ctx *parser.CreatorContext) {
	method := methodMap[getMethodMapName(currentMethod)]
	fullType, _ := WarpTargetFullType(createdName)

	position := BuildPosition(ctx.BaseParserRuleContext, createdName)

	jMethodCall := &core_domain.CodeCall{
		Package:  RemoveTarget(fullType),
		Type:     "CreatorClass",
		NodeName: createdName,
		Position: position,
	}

	method.FunctionCalls = append(method.FunctionCalls, *jMethodCall)
	methodMap[getMethodMapName(currentMethod)] = method
}

func BuildPosition(ctx *antlr.BaseParserRuleContext, nodeName string) core_domain.CodePosition {
	position := core_domain.CodePosition{
		StartLine:         ctx.GetStart().GetLine(),
		StartLinePosition: ctx.GetStart().GetColumn(),
		StopLine:          ctx.GetStop().GetLine(),
		StopLinePosition:  ctx.GetStop().GetColumn() + len(nodeName),
	}

	return position
}

func (s *JavaFullListener) EnterMethodCall(ctx *parser.MethodCallContext) {
	var jMethodCall = core_domain.NewCodeMethodCall()

	targetCtx := ctx.GetParent().GetChild(0).(antlr.ParseTree)
	var targetType = ParseTargetType(targetCtx.GetText())

	if targetCtx.GetChild(0) != nil {
		switch x := targetCtx.GetChild(0).(type) {
		case *parser.MethodCallContext:
			targetType = x.IDENTIFIER().GetText()
		}
	}

	callee := ctx.GetChild(0).(antlr.ParseTree).GetText()

	BuildMethodCallLocation(&jMethodCall, ctx, callee)
	BuildMethodCallMethod(&jMethodCall, callee, targetType, ctx)
	BuildMethodCallParameters(&jMethodCall, ctx)

	sendResultToMethodCallMap(jMethodCall)
}

func sendResultToMethodCallMap(jMethodCall core_domain.CodeCall) {
	methodCalls = append(methodCalls, jMethodCall)

	method := methodMap[getMethodMapName(currentMethod)]
	method.FunctionCalls = append(method.FunctionCalls, jMethodCall)
	methodMap[getMethodMapName(currentMethod)] = method
}

func isChainCall(targetType string) bool {
	return strings.Contains(targetType, "(") && strings.Contains(targetType, ")") && strings.Contains(targetType, ".")
}

func buildSelfThisTarget(targetType string) string {
	targetType = strings.ReplaceAll(targetType, "this.", "")
	for _, field := range fields {
		if field.TypeValue == targetType {
			targetType = field.TypeType
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
		targetType := ParseTargetType(text)

		fullType, _ := WarpTargetFullType(targetType)

		position := BuildPosition(ctx.BaseParserRuleContext, text)

		jMethodCall := &core_domain.CodeCall{
			Package:      RemoveTarget(fullType),
			Type:         "lambda",
			NodeName:     targetType,
			FunctionName: methodName,
			Position:     position,
		}
		sendResultToMethodCallMap(*jMethodCall)
	}
}

func (s *JavaFullListener) AppendClasses(classes []string) {
	clzs = classes
}

func buildExtend(extendName string) string {
	var extend = extendName
	target, _ := WarpTargetFullType(extendName)
	if target != "" {
		extend = target
	}

	return extend
}

func buildFieldCall(typeType string, ctx *parser.FieldDeclarationContext) {
	target, _ := WarpTargetFullType(typeType)
	if target != "" {
		position := BuildPosition(ctx.BaseParserRuleContext, target)

		jMethodCall := &core_domain.CodeCall{
			Package:  RemoveTarget(target),
			Type:     "field",
			NodeName: typeType,
			Position: position,
		}

		currentNode.FunctionCalls = append(currentNode.FunctionCalls, *jMethodCall)
	}
}

func buildImplement(text string) string {
	target, _ := WarpTargetFullType(text)
	return target
}
