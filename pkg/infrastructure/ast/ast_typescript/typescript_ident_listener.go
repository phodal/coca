package ast_typescript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/ts"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/astutil"
	"strings"
)

var defaultClass = "default"

type TypeScriptIdentListener struct {
	currentDataStruct *core_domain.CodeDataStruct
	dataStructures    []core_domain.CodeDataStruct
	dataStructQueue   []core_domain.CodeDataStruct
	filePath          string
	codeFile          core_domain.CodeContainer

	parser.BaseTypeScriptParserListener
}

func NewTypeScriptIdentListener(fileName string) *TypeScriptIdentListener {
	listener := &TypeScriptIdentListener{}
	listener.filePath = fileName
	return listener
}

func (s *TypeScriptIdentListener) GetNodeInfo() core_domain.CodeContainer {
	isScriptCalls := s.currentDataStruct != nil && s.currentDataStruct.IsNotEmpty()
	if isScriptCalls {
		if len(s.currentDataStruct.Functions) < 1 {
			function := &core_domain.CodeFunction{}
			function.Name = "default"
			function.FunctionCalls = append(function.FunctionCalls, s.currentDataStruct.FunctionCalls...)

			s.currentDataStruct.Functions = append(s.currentDataStruct.Functions, *function)
		}

		s.dataStructures = append(s.dataStructures, *s.currentDataStruct)
	}

	s.codeFile.DataStructures = s.dataStructures
	return s.codeFile
}

func (s *TypeScriptIdentListener) EnterImportFromBlock(ctx *parser.ImportFromBlockContext) {
	replaceSingleQuote := UpdateImportStr(ctx.StringLiteral().GetText())
	imp := &core_domain.CodeImport{Source: replaceSingleQuote}
	importName := ctx.GetChild(0).(antlr.ParseTree).GetText()
	imp.ImportName = importName
	s.codeFile.Imports = append(s.codeFile.Imports, *imp)
}

func UpdateImportStr(importText string) string {
	replaceDoubleQuote := strings.ReplaceAll(importText, "\"", "")
	replaceSingleQuote := strings.ReplaceAll(replaceDoubleQuote, "'", "")
	return replaceSingleQuote
}

func (s *TypeScriptIdentListener) EnterImportAliasDeclaration(ctx *parser.ImportAliasDeclarationContext) {
	replaceSingleQuote := UpdateImportStr(ctx.StringLiteral().GetText())
	imp := &core_domain.CodeImport{Source: replaceSingleQuote}
	s.codeFile.Imports = append(s.codeFile.Imports, *imp)
}

func (s *TypeScriptIdentListener) EnterImportAll(ctx *parser.ImportAllContext) {
	replaceSingleQuote := UpdateImportStr(ctx.StringLiteral().GetText())
	imp := &core_domain.CodeImport{Source: replaceSingleQuote}
	s.codeFile.Imports = append(s.codeFile.Imports, *imp)
}

func (s *TypeScriptIdentListener) EnterInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	s.currentDataStruct = &core_domain.CodeDataStruct{
		Type:     "Interface",
		NodeName: ctx.Identifier().GetText(),
	}

	if ctx.InterfaceExtendsClause() != nil {
		extendsContext := ctx.InterfaceExtendsClause().(*parser.InterfaceExtendsClauseContext)
		elements := BuildImplements(extendsContext.ClassOrInterfaceTypeList())
		s.currentDataStruct.Extend = elements[0]
	}

	objectTypeCtx := ctx.ObjectType().(*parser.ObjectTypeContext)
	if objectTypeCtx.TypeBody() != nil {
		bodyCtx := objectTypeCtx.TypeBody().(*parser.TypeBodyContext)
		typeMemberListCtx := bodyCtx.TypeMemberList().(*parser.TypeMemberListContext)

		BuildInterfaceTypeBody(typeMemberListCtx, s.currentDataStruct)
	}
}

func BuildInterfaceTypeBody(ctx *parser.TypeMemberListContext, dataStruct *core_domain.CodeDataStruct) {
	for _, typeMember := range ctx.AllTypeMember() {
		typeMemberCtx := typeMember.(*parser.TypeMemberContext)
		memberChild := typeMemberCtx.GetChild(0)
		switch x := memberChild.(type) {
		case *parser.PropertySignatureContext:
			BuildInterfacePropertySignature(x, dataStruct)
		case *parser.MethodSignatureContext:
			function := core_domain.CodeFunction{
				Name: x.PropertyName().GetText(),
			}

			FillMethodFromCallSignature(x.CallSignature().(*parser.CallSignatureContext), &function)

			dataStruct.Functions = append(dataStruct.Functions, function)
		}
	}
}

func BuildInterfacePropertySignature(signatureCtx *parser.PropertySignatureContext, dataStruct *core_domain.CodeDataStruct) {
	typeType := BuildTypeAnnotation(signatureCtx.TypeAnnotation().(*parser.TypeAnnotationContext))
	typeValue := signatureCtx.PropertyName().(*parser.PropertyNameContext).GetText()

	isArrowFunc := signatureCtx.Type_() != nil
	if isArrowFunc {
		function := &core_domain.CodeFunction{
			Name: typeValue,
		}
		param := core_domain.CodeProperty{
			TypeValue: "any",
			TypeType:  typeType,
		}

		returnType := core_domain.CodeProperty{
			TypeType: signatureCtx.Type_().GetText(),
		}
		function.Parameters = append(function.Parameters, param)
		function.MultipleReturns = append(function.MultipleReturns, returnType)

		dataStruct.Functions = append(dataStruct.Functions, *function)
	} else {
		codeField := &core_domain.CodeField{}
		codeField.TypeType = typeType
		codeField.TypeValue = typeValue

		dataStruct.Fields = append(dataStruct.Fields, *codeField)
	}
}

func (s *TypeScriptIdentListener) ExitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	s.exitClass()
}

func (s *TypeScriptIdentListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	s.currentDataStruct = &core_domain.CodeDataStruct{
		Type:     "Class",
		NodeName: ctx.Identifier().GetText(),
	}

	heritageContext := ctx.ClassHeritage().(*parser.ClassHeritageContext)
	if heritageContext.ImplementsClause() != nil {
		typeList := heritageContext.ImplementsClause().(*parser.ImplementsClauseContext).ClassOrInterfaceTypeList()

		implements := BuildImplements(typeList)
		s.currentDataStruct.Implements = implements
	}

	if heritageContext.ClassExtendsClause() != nil {
		referenceContext := heritageContext.ClassExtendsClause().(*parser.ClassExtendsClauseContext).TypeReference().(*parser.TypeReferenceContext)

		s.currentDataStruct.Extend = referenceContext.TypeName().GetText()
	}

	classTailContext := ctx.ClassTail().(*parser.ClassTailContext)
	s.handleClassBodyElements(classTailContext)

	s.dataStructQueue = append(s.dataStructQueue, *s.currentDataStruct)
}

func (s *TypeScriptIdentListener) handleClassBodyElements(classTailContext *parser.ClassTailContext) {
	for _, classElement := range classTailContext.AllClassElement() {
		elementChild := classElement.GetChild(0)
		switch x := elementChild.(type) {
		case *parser.ConstructorDeclarationContext:
			codeFunction := BuildConstructorMethod(x)

			s.currentDataStruct.Functions = append(s.currentDataStruct.Functions, *codeFunction)
		case *parser.PropertyMemberDeclarationContext:
			s.HandlePropertyMember(x, s.currentDataStruct)
		}
	}
}

func (s *TypeScriptIdentListener) HandlePropertyMember(propertyMemberCtx *parser.PropertyMemberDeclarationContext, dataStruct *core_domain.CodeDataStruct) {
	callSignatureSizePos := 3
	if propertyMemberCtx.PropertyName() != nil {
		modifier := propertyMemberCtx.PropertyMemberBase().GetText()
		codeField := core_domain.CodeField{
			TypeValue: propertyMemberCtx.PropertyName().GetText(),
		}
		codeField.Modifiers = append(codeField.Modifiers, modifier)
		if propertyMemberCtx.TypeAnnotation() != nil {
			codeField.TypeType = BuildTypeAnnotation(propertyMemberCtx.TypeAnnotation().(*parser.TypeAnnotationContext))
		}

		dataStruct.Fields = append(dataStruct.Fields, codeField)
	}

	if propertyMemberCtx.GetChildCount() >= callSignatureSizePos {
		callSignCtxPos := 2
		switch propertyMemberCtx.GetChild(callSignCtxPos).(type) {
		case *parser.CallSignatureContext:
			memberFunction := BuildMemberMethod(propertyMemberCtx)
			dataStruct.Functions = append(dataStruct.Functions, *memberFunction)
		}
	}
}

func (s *TypeScriptIdentListener) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {
	s.exitClass()
}

func (s *TypeScriptIdentListener) exitClass() {
	s.dataStructures = append(s.dataStructures, *s.currentDataStruct)
	if len(s.dataStructQueue) > 1 {
		s.dataStructQueue = s.dataStructQueue[0 : len(s.dataStructQueue)-1]
		s.currentDataStruct = &s.dataStructQueue[len(s.dataStructQueue)-1]
	} else {
		s.currentDataStruct = core_domain.NewDataStruct()
	}
}

func (s *TypeScriptIdentListener) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	function := &core_domain.CodeFunction{
		Name: ctx.Identifier().GetText(),
	}

	callSignatureContext := ctx.CallSignature().(*parser.CallSignatureContext)
	FillMethodFromCallSignature(callSignatureContext, function)

	astutil.AddFunctionPosition(function, ctx.GetChild(0).GetParent().(*antlr.BaseParserRuleContext))

	if s.currentDataStruct == nil {
		s.currentDataStruct = &core_domain.CodeDataStruct{}
	}
	s.currentDataStruct.Functions = append(s.currentDataStruct.Functions, *function)
}

func FillMethodFromCallSignature(callSignatureContext *parser.CallSignatureContext, function *core_domain.CodeFunction) {
	if callSignatureContext.ParameterList() != nil {
		parameterListContext := callSignatureContext.ParameterList().(*parser.ParameterListContext)
		functionParameters := BuildMethodParameter(parameterListContext)

		function.Parameters = append(function.Parameters, functionParameters...)
	}

	if callSignatureContext.TypeAnnotation() != nil {
		annotationContext := callSignatureContext.TypeAnnotation().(*parser.TypeAnnotationContext)
		typeAnnotation := BuildTypeAnnotation(annotationContext)

		returnType := function.BuildSingleReturnType(typeAnnotation)
		function.MultipleReturns = append(function.MultipleReturns, *returnType)
	}
}
