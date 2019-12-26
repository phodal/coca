package bs

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	models2 "github.com/phodal/coca/core/adapter/bs/models"
	. "github.com/phodal/coca/languages/java"
	"reflect"
	"strings"
)

var imports []string
var clzs []string
var currentPkg string
var currentClz string
var currentClzType string

var currentClzExtends string
var currentClzImplements []string

var methods []models2.BsJMethod
var methodCalls []models2.BsJMethodCall

var fields = make(map[string]string)
var localVars = make(map[string]string)
var formalParameters = make(map[string]string)
var currentClassBs models2.ClassBadSmellInfo

func NewBadSmellListener() *BadSmellListener {
	currentClz = ""
	currentPkg = ""
	methods = nil
	methodCalls = nil
	currentClzImplements = nil
	currentClzExtends = ""
	return &BadSmellListener{}
}

type BadSmellListener struct {
	BaseJavaParserListener
}

func (s *BadSmellListener) getNodeInfo() models2.BsJClass {
	return *&models2.BsJClass{
		currentPkg,
		currentClz,
		currentClzType,
		"",
		currentClzExtends,
		currentClzImplements,
		methods,
		methodCalls,
		currentClassBs,
	}
}

func (s *BadSmellListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {
	currentPkg = ctx.QualifiedName().GetText()
}

func (s *BadSmellListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {
	importText := ctx.QualifiedName().GetText()
	imports = append(imports, importText)
}

func (s *BadSmellListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {
	currentClzType = "Class"
	currentClz = ctx.IDENTIFIER().GetText()

	if ctx.EXTENDS() != nil {
		currentClzExtends = ctx.TypeType().GetText()
	}

	if ctx.IMPLEMENTS() != nil {
		typeList := ctx.TypeList().(*TypeListContext)
		for _, typ := range typeList.AllTypeType() {
			typeData := getTypeDATA(typ.(*TypeTypeContext))
			currentClzImplements = append(currentClzImplements, typeData)
		}
	}
}

func getTypeDATA(typ *TypeTypeContext) string {
	var typeData string
	classOrInterface := typ.ClassOrInterfaceType().(*ClassOrInterfaceTypeContext)
	if classOrInterface != nil {
		identifiers := classOrInterface.AllIDENTIFIER()
		typeData = identifiers[len(identifiers)-1].GetText()
	}

	return typeData
}

func (s *BadSmellListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {
	currentClzType = "Interface"
	currentClz = ctx.IDENTIFIER().GetText()
}

func (s *BadSmellListener) EnterInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.IDENTIFIER().GetSymbol().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	name := ctx.IDENTIFIER().GetText()
	stopLinePosition := startLinePosition + len(name)
	methodBody := ctx.MethodBody().GetText()

	var modifiers = ""
	allModifier := ctx.AllInterfaceMethodModifier()
	methodModifierLen := len(allModifier)
	for index, modifier := range allModifier {
		modifiers = modifiers + modifier.GetText()
		if index < methodModifierLen-1 {
			modifiers = modifiers + ","
		}
	}

	typeType := ctx.TypeTypeOrVoid().GetText()

	var methodParams []models2.JFullParameter = nil
	parameters := ctx.FormalParameters()
	if parameters != nil {
		if reflect.TypeOf(parameters.GetChild(1)).String() == "*parser.FormalParameterListContext" {
			allFormal := parameters.GetChild(1).(*FormalParameterListContext)
			formalParameter := allFormal.AllFormalParameter()
			for _, param := range formalParameter {
				paramContext := param.(*FormalParameterContext)
				paramType := paramContext.TypeType().GetText()
				paramValue := paramContext.VariableDeclaratorId().(*VariableDeclaratorIdContext).IDENTIFIER().GetText()
				methodParams = append(methodParams, *&models2.JFullParameter{paramType, paramValue})
			}
		}
	}

	methodBSInfo := models2.NewMethodBadSmellInfo()

	method := &models2.BsJMethod{
		Name:              name,
		Type:              typeType,
		StartLine:         startLine,
		StartLinePosition: startLinePosition,
		StopLine:          stopLine,
		StopLinePosition:  stopLinePosition,
		MethodBody:        methodBody,
		Modifier:          modifiers,
		Parameters:        methodParams,
		MethodBs:          methodBSInfo,
	}

	methods = append(methods, *method)
}

func (s *BadSmellListener) EnterFormalParameter(ctx *FormalParameterContext) {
	formalParameters[ctx.VariableDeclaratorId().GetText()] = ctx.TypeType().GetText()
}

func (s *BadSmellListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {
	declarators := ctx.VariableDeclarators()
	variableName := declarators.GetParent().GetChild(0).(antlr.ParseTree).GetText()

	for _, declarator := range declarators.(*VariableDeclaratorsContext).AllVariableDeclarator() {
		value := declarator.(*VariableDeclaratorContext).VariableDeclaratorId().(*VariableDeclaratorIdContext).IDENTIFIER().GetText()
		fields[value] = variableName
	}
}

func (s *BadSmellListener) EnterLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {
	typ := ctx.GetChild(0).(antlr.ParseTree).GetText()
	variableName := ctx.GetChild(1).GetChild(0).GetChild(0).(antlr.ParseTree).GetText()
	localVars[variableName] = typ
}

func (s *BadSmellListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.IDENTIFIER().GetSymbol().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	name := ctx.IDENTIFIER().GetText()
	stopLinePosition := startLinePosition + len(name)

	modifier := getModifier(ctx)

	typeType := ctx.TypeTypeOrVoid().GetText()
	methodBody := ctx.MethodBody().GetText()

	var methodParams []models2.JFullParameter = nil
	parameters := ctx.FormalParameters()
	if parameters != nil {
		if reflect.TypeOf(parameters.GetChild(1)).String() == "*parser.FormalParameterListContext" {
			allFormal := parameters.GetChild(1).(*FormalParameterListContext)
			formalParameter := allFormal.AllFormalParameter()
			for _, param := range formalParameter {
				paramContext := param.(*FormalParameterContext)
				paramType := paramContext.TypeType().GetText()
				paramValue := paramContext.VariableDeclaratorId().(*VariableDeclaratorIdContext).IDENTIFIER().GetText()
				methodParams = append(methodParams, models2.JFullParameter{paramType, paramValue})

				localVars[paramValue] = paramType
			}
		}
	}

	methodBSInfo := models2.NewMethodBadSmellInfo()
	methodBadSmellInfo := buildMethodBSInfo(ctx, methodBSInfo)

	method := &models2.BsJMethod{
		Name:              name,
		Type:              typeType,
		StartLine:         startLine,
		StartLinePosition: startLinePosition,
		StopLine:          stopLine,
		StopLinePosition:  stopLinePosition,
		MethodBody:        methodBody,
		Modifier:          modifier,
		Parameters:        methodParams,
		MethodBs:          methodBadSmellInfo,
	}
	methods = append(methods, *method)
}

func getModifier(ctx *MethodDeclarationContext) string {
	var modifier = ""
	if reflect.TypeOf(ctx.GetParent()).String() == "*parser.MemberDeclarationContext" {
		firstChild := ctx.GetParent().(*MemberDeclarationContext).GetParent().GetChild(0)
		if reflect.TypeOf(firstChild).String() == "*parser.ModifierContext" {
			modifierCtx := firstChild.(*ModifierContext)
			if reflect.TypeOf(modifierCtx.GetChild(0)).String() == "*parser.ClassOrInterfaceModifierContext" {
				context := modifierCtx.GetChild(0).(*ClassOrInterfaceModifierContext)
				modifier = context.GetText()
			}
		}
	}
	return modifier
}

func buildMethodBSInfo(context *MethodDeclarationContext, bsInfo models2.MethodBadSmellInfo) models2.MethodBadSmellInfo {
	methodBody := context.MethodBody()
	blockContext := methodBody.GetChild(0)
	if reflect.TypeOf(blockContext).String() == "*parser.BlockContext" {
		blcStatement := blockContext.(*BlockContext).AllBlockStatement()
		for _, statement := range blcStatement {
			if reflect.TypeOf(statement.GetChild(0)).String() == "*parser.StatementContext" {
				if len(statement.GetChild(0).(*StatementContext).GetChildren()) < 3 {
					continue
				}

				statementCtx := statement.GetChild(0).(*StatementContext)
				if (reflect.TypeOf(statementCtx.GetChild(1)).String()) == "*parser.ParExpressionContext" {
					if statementCtx.GetChild(0).(antlr.ParseTree).GetText() == "if" {
						if reflect.TypeOf(statementCtx.GetChild(1)).String() == "*parser.ParExpressionContext" {
							parCtx := statementCtx.GetChild(1).(*ParExpressionContext)
							startLine := parCtx.GetStart().GetLine()
							endLine := parCtx.GetStop().GetLine()

							info := models2.NewIfPairInfo()
							info.StartLine = startLine
							info.EndLine = endLine
							bsInfo.IfInfo = append(bsInfo.IfInfo, info)
						}

						bsInfo.IfSize = bsInfo.IfSize + 1
					}

					if statementCtx.GetChild(0).(antlr.ParseTree).GetText() == "switch" {
						bsInfo.SwitchSize = bsInfo.SwitchSize + 1
					}

				}
			}
		}
	}

	return bsInfo
}

func (s *BadSmellListener) EnterFormalParameterList(ctx *FormalParameterListContext) {
	//fmt.Println(ctx.GetParent().GetParent().(antlr.RuleNode).get)
	//fmt.Println(ctx.AllFormalParameter()
}

func (s *BadSmellListener) EnterAnnotation(ctx *AnnotationContext) {
	if currentClzType == "Class" && ctx.QualifiedName().GetText() == "Override" {
		currentClassBs.OverrideSize++
	}
}

func (s *BadSmellListener) EnterCreator(ctx *CreatorContext) {
	variableName := ctx.GetParent().GetParent().GetChild(0).(antlr.ParseTree).GetText()
	localVars[variableName] = ctx.CreatedName().GetText()
}

func (s *BadSmellListener) EnterLocalTypeDeclaration(ctx *LocalTypeDeclarationContext) {

}

func (s *BadSmellListener) EnterMethodCall(ctx *MethodCallContext) {
	var targetCtx = ctx.GetParent().GetChild(0).(antlr.ParseTree).GetText()
	var targetType = parseTargetType(targetCtx)
	callee := ctx.GetChild(0).(antlr.ParseTree).GetText()

	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := startLinePosition + len(callee)

	//typeType := ctx.GetChild(0).(antlr.ParseTree).TypeTypeOrVoid().GetText()

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
		jMethodCall := *&models2.BsJMethodCall{removeTarget(fullType), "", targetType, callee, startLine, startLinePosition, stopLine, stopLinePosition}
		methodCalls = append(methodCalls, jMethodCall)
	} else {
		if ctx.GetText() == targetType {
			jMethodCall := *&models2.BsJMethodCall{currentPkg, "", currentClz, callee, startLine, startLinePosition, stopLine, stopLinePosition}
			methodCalls = append(methodCalls, jMethodCall)
		} else {
			jMethodCall := *&models2.BsJMethodCall{currentPkg, "NEEDFIX", targetType, callee, startLine, startLinePosition, stopLine, stopLinePosition}
			methodCalls = append(methodCalls, jMethodCall)
		}
	}
}

func (s *BadSmellListener) EnterExpression(ctx *ExpressionContext) {
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

		jMethodCall := &models2.BsJMethodCall{removeTarget(fullType), "", targetType, methodName, startLine, startLinePosition, stopLine, stopLinePosition}
		methodCalls = append(methodCalls, *jMethodCall)
	}
}

func (s *BadSmellListener) appendClasses(classes []string) {
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
