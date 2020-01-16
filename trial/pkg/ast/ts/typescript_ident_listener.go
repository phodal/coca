package ts

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/ts"
	"github.com/phodal/coca/pkg/domain"
	"github.com/phodal/coca/pkg/domain/trial"
	"github.com/phodal/coca/trial/pkg/ast/ast_util"
	"strings"
)

var defaultClass = "default"

type TypeScriptIdentListener struct {
	currentNode    *domain.JClassNode
	classNodeQueue []domain.JClassNode
	classNodes     []domain.JClassNode

	currentDataStruct *trial.CodeDataStruct
	dataStructures    []trial.CodeDataStruct
	dataStructQueue   []trial.CodeDataStruct
	filePath          string
	codeFile          trial.CodeFile

	parser.BaseTypeScriptParserListener
}

func NewTypeScriptIdentListener(fileName string) *TypeScriptIdentListener {
	listener := &TypeScriptIdentListener{}
	listener.filePath = fileName
	return listener
}

func (s *TypeScriptIdentListener) GetNodeInfo() trial.CodeFile {
	if s.currentNode != nil && s.currentNode.IsNotEmpty() {
		s.currentNode.Class = defaultClass
		s.currentNode.Type = "Default"
		s.classNodes = append(s.classNodes, *s.currentNode)
		s.currentNode = domain.NewClassNode()
	}

	isScriptCalls := s.currentDataStruct != nil && s.currentDataStruct.IsNotEmpty()
	if isScriptCalls {
		if len(s.currentDataStruct.Functions) < 1 {
			function := &trial.CodeFunction{}
			function.Name = "default"
			function.MethodCalls = append(function.MethodCalls, s.currentDataStruct.FunctionCalls...)

			s.currentDataStruct.Functions = append(s.currentDataStruct.Functions, *function)
		}

		s.dataStructures = append(s.dataStructures, *s.currentDataStruct)
	}

	s.codeFile.ClassNodes = s.classNodes
	s.codeFile.DataStructures = s.dataStructures
	return s.codeFile
}

func (s *TypeScriptIdentListener) EnterImportFromBlock(ctx *parser.ImportFromBlockContext) {
	replaceSingleQuote := UpdateImportStr(ctx.StringLiteral().GetText())
	imp := &trial.CodeImport{Source: replaceSingleQuote}
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
	imp := &trial.CodeImport{Source: replaceSingleQuote}
	s.codeFile.Imports = append(s.codeFile.Imports, *imp)
}

func (s *TypeScriptIdentListener) EnterImportAll(ctx *parser.ImportAllContext) {
	replaceSingleQuote := UpdateImportStr(ctx.StringLiteral().GetText())
	imp := &trial.CodeImport{Source: replaceSingleQuote}
	s.codeFile.Imports = append(s.codeFile.Imports, *imp)
}

func (s *TypeScriptIdentListener) EnterInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	s.currentNode = domain.NewClassNode()
	s.currentNode.Type = "Interface"
	s.currentNode.Class = ctx.Identifier().GetText()

	s.currentDataStruct = &trial.CodeDataStruct{
		Type: "Interface",
		Name: ctx.Identifier().GetText(),
	}

	if ctx.InterfaceExtendsClause() != nil {
		extendsContext := ctx.InterfaceExtendsClause().(*parser.InterfaceExtendsClauseContext)
		implements := BuildImplements(extendsContext.ClassOrInterfaceTypeList())
		s.currentNode.Extend = implements[0]

		s.currentDataStruct.Extend = implements[0]
	}

	objectTypeCtx := ctx.ObjectType().(*parser.ObjectTypeContext)
	if objectTypeCtx.TypeBody() != nil {
		typeMemberListCtx := objectTypeCtx.TypeBody().(*parser.TypeBodyContext).TypeMemberList().(*parser.TypeMemberListContext)
		BuildInterfaceTypeBody(typeMemberListCtx, s.currentNode, s.currentDataStruct)
	}
}

func BuildInterfaceTypeBody(ctx *parser.TypeMemberListContext, classNode *domain.JClassNode, dataStruct *trial.CodeDataStruct) {
	for _, typeMember := range ctx.AllTypeMember() {
		typeMemberCtx := typeMember.(*parser.TypeMemberContext)
		memberChild := typeMemberCtx.GetChild(0)
		switch x := memberChild.(type) {
		case *parser.PropertySignatureContext:
			BuildInterfacePropertySignature(x, classNode, dataStruct)
		case *parser.MethodSignatureContext:
			method := domain.NewJMethod()
			method.Name = x.PropertyName().GetText()

			function := trial.CodeFunction{
				Name: x.PropertyName().GetText(),
			}

			FillMethodFromCallSignature(x.CallSignature().(*parser.CallSignatureContext), &method, &function)

			dataStruct.Functions = append(dataStruct.Functions, function)
			classNode.Methods = append(classNode.Methods, method)
		}
	}
}

func BuildInterfacePropertySignature(signatureCtx *parser.PropertySignatureContext, classNode *domain.JClassNode, dataStruct *trial.CodeDataStruct) {
	typeType := BuildTypeAnnotation(signatureCtx.TypeAnnotation().(*parser.TypeAnnotationContext))
	typeValue := signatureCtx.PropertyName().(*parser.PropertyNameContext).GetText()

	isArrowFunc := signatureCtx.Type_() != nil
	if isArrowFunc {
		method := domain.NewJMethod()
		method.Name = typeValue
		parameter := domain.JParameter{
			Name: "any",
			Type: typeType,
		}
		method.Parameters = append(method.Parameters, parameter)
		method.Type = signatureCtx.Type_().GetText()

		classNode.Methods = append(classNode.Methods, method)

		function := &trial.CodeFunction{
			Name: typeValue,
		}
		param := trial.CodeProperty{
			Name:     "any",
			TypeType: typeType,
		}

		returnType := trial.CodeProperty{
			TypeType: signatureCtx.Type_().GetText(),
		}
		function.Parameters = append(function.Parameters, param)
		function.ReturnTypes = append(function.ReturnTypes, returnType)

		dataStruct.Functions = append(dataStruct.Functions, *function)
	} else {
		field := &domain.JField{
			Type:  typeType,
			Value: typeValue,
		}

		codeField := &trial.CodeField{}
		codeField.TypeType = typeType
		codeField.TypeValue = typeValue

		classNode.Fields = append(classNode.Fields, *field)
		dataStruct.Fields = append(dataStruct.Fields, *codeField)
	}
}

func (s *TypeScriptIdentListener) ExitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	s.exitClass()
}

func (s *TypeScriptIdentListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	s.currentNode = domain.NewClassNode()
	s.currentNode.Type = "Class"
	s.currentNode.Class = ctx.Identifier().GetText()

	s.currentDataStruct = &trial.CodeDataStruct{
		Type: "Class",
		Name: ctx.Identifier().GetText(),
	}

	heritageContext := ctx.ClassHeritage().(*parser.ClassHeritageContext)
	if heritageContext.ImplementsClause() != nil {
		typeList := heritageContext.ImplementsClause().(*parser.ImplementsClauseContext).ClassOrInterfaceTypeList()

		implements := BuildImplements(typeList)
		s.currentNode.Implements = implements
		s.currentDataStruct.Implements = implements
	}

	if heritageContext.ClassExtendsClause() != nil {
		referenceContext := heritageContext.ClassExtendsClause().(*parser.ClassExtendsClauseContext).TypeReference().(*parser.TypeReferenceContext)

		s.currentNode.Extend = referenceContext.TypeName().GetText()
		s.currentDataStruct.Extend = referenceContext.TypeName().GetText()
	}

	classTailContext := ctx.ClassTail().(*parser.ClassTailContext)
	s.handleClassBodyElements(classTailContext)

	s.classNodeQueue = append(s.classNodeQueue, *s.currentNode)
	s.dataStructQueue = append(s.dataStructQueue, *s.currentDataStruct)
}

func (s *TypeScriptIdentListener) handleClassBodyElements(classTailContext *parser.ClassTailContext) {
	for _, classElement := range classTailContext.AllClassElement() {
		elementChild := classElement.GetChild(0)
		switch x := elementChild.(type) {
		case *parser.ConstructorDeclarationContext:
			constructorMethod, codeFunction := BuildConstructorMethod(x)

			s.currentNode.Methods = append(s.currentNode.Methods, constructorMethod)
			s.currentDataStruct.Functions = append(s.currentDataStruct.Functions, *codeFunction)
		case *parser.PropertyMemberDeclarationContext:
			s.HandlePropertyMember(x, s.currentNode, s.currentDataStruct)
		}
	}
}

func (s *TypeScriptIdentListener) HandlePropertyMember(propertyMemberCtx *parser.PropertyMemberDeclarationContext, node *domain.JClassNode, dataStruct *trial.CodeDataStruct) {
	callSignatureSizePos := 3
	if propertyMemberCtx.PropertyName() != nil {
		field := domain.JField{}
		field.Value = propertyMemberCtx.PropertyName().GetText()
		field.Modifier = propertyMemberCtx.PropertyMemberBase().GetText()
		if propertyMemberCtx.TypeAnnotation() != nil {
			field.Type = BuildTypeAnnotation(propertyMemberCtx.TypeAnnotation().(*parser.TypeAnnotationContext))
		}
		node.Fields = append(s.currentNode.Fields, field)
	}

	if propertyMemberCtx.GetChildCount() >= callSignatureSizePos {
		callSignCtxPos := 2
		switch propertyMemberCtx.GetChild(callSignCtxPos).(type) {
		case *parser.CallSignatureContext:
			memberMethod, memberFunction := BuildMemberMethod(propertyMemberCtx)
			node.Methods = append(s.currentNode.Methods, memberMethod)

			dataStruct.Functions = append(dataStruct.Functions, *memberFunction)
		}
	}

}

func (s *TypeScriptIdentListener) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {
	s.exitClass()
}

func (s *TypeScriptIdentListener) exitClass() {
	s.classNodes = append(s.classNodes, *s.currentNode)
	if len(s.classNodeQueue) > 1 {
		s.classNodeQueue = s.classNodeQueue[0 : len(s.classNodeQueue)-1]
		s.currentNode = &s.classNodeQueue[len(s.classNodeQueue)-1]
	} else {
		s.currentNode = domain.NewClassNode()
	}

	s.dataStructures = append(s.dataStructures, *s.currentDataStruct)
	if len(s.dataStructQueue) > 1 {
		s.dataStructQueue = s.dataStructQueue[0 : len(s.dataStructQueue)-1]
		s.currentDataStruct = &s.dataStructQueue[len(s.dataStructQueue)-1]
	} else {
		s.currentDataStruct = trial.NewDataStruct()
	}
}

func (s *TypeScriptIdentListener) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	method := domain.NewJMethod()

	method.Name = ctx.Identifier().GetText()
	ast_util.AddPosition(&method, ctx.GetChild(0).GetParent().(*antlr.BaseParserRuleContext))
	function := &trial.CodeFunction{
		Name: ctx.Identifier().GetText(),
	}
	
	callSignatureContext := ctx.CallSignature().(*parser.CallSignatureContext)
	FillMethodFromCallSignature(callSignatureContext, &method, function)

	ast_util.AddFunctionPosition(function, ctx.GetChild(0).GetParent().(*antlr.BaseParserRuleContext))

	if s.currentNode == nil {
		s.currentNode = domain.NewClassNode()
	}
	if s.currentDataStruct == nil {
		s.currentDataStruct = &trial.CodeDataStruct{}
	}
	s.currentNode.Methods = append(s.currentNode.Methods, method)
	s.currentDataStruct.Functions = append(s.currentDataStruct.Functions, *function)
}

func FillMethodFromCallSignature(callSignatureContext *parser.CallSignatureContext, method *domain.JMethod, function *trial.CodeFunction) {
	if callSignatureContext.ParameterList() != nil {
		parameterListContext := callSignatureContext.ParameterList().(*parser.ParameterListContext)
		methodParameters, _ := BuildMethodParameter(parameterListContext)

		method.Parameters = append(method.Parameters, methodParameters...)
	}

	if callSignatureContext.TypeAnnotation() != nil {
		annotationContext := callSignatureContext.TypeAnnotation().(*parser.TypeAnnotationContext)
		typeAnnotation := BuildTypeAnnotation(annotationContext)
		method.Type = typeAnnotation

		returnType := function.BuildSingleReturnType(typeAnnotation)
		function.ReturnTypes = append(function.ReturnTypes, *returnType)
	}
}
