package ts

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/ts"
	"github.com/phodal/coca/pkg/domain"
	"github.com/phodal/coca/pkg/domain/trial"
	"strings"
)

var currentNode *domain.JClassNode
var classNodeQueue []domain.JClassNode
var classNodes []domain.JClassNode

var currentDataStruct *trial.CodeDataStruct
var dataStructures []trial.CodeDataStruct
var dataStructQueue []trial.CodeDataStruct

var defaultClass = "default"
var filePath string
var codeFile trial.CodeFile

type TypeScriptIdentListener struct {
	parser.BaseTypeScriptParserListener
}

func NewTypeScriptIdentListener(fileName string) *TypeScriptIdentListener {
	classNodes = nil
	filePath = fileName
	currentNode = domain.NewClassNode()
	currentDataStruct = trial.NewDataStruct()
	codeFile = trial.CodeFile{
		FullName:       filePath,
		Imports:        nil,
		ClassNodes:     nil,
		DataStructures: nil,
	}
	return &TypeScriptIdentListener{}
}

func (s *TypeScriptIdentListener) GetNodeInfo() trial.CodeFile {
	if currentNode.IsNotEmpty() {
		currentNode.Class = defaultClass
		currentNode.Type = "Default"
		classNodes = append(classNodes, *currentNode)
		currentNode = domain.NewClassNode()
	}

	isScriptCalls := currentDataStruct.IsNotEmpty()
	if isScriptCalls {
		//function := &trial.CodeFunction{}
	}

	codeFile.ClassNodes = classNodes
	codeFile.DataStructures = dataStructures
	return codeFile
}

func (s *TypeScriptIdentListener) EnterImportFromBlock(ctx *parser.ImportFromBlockContext) {
	replaceSingleQuote := UpdateImportStr(ctx.StringLiteral().GetText())
	imp := &trial.CodeImport{Source: replaceSingleQuote}
	importName := ctx.GetChild(0).(antlr.ParseTree).GetText()
	imp.ImportName = importName
	codeFile.Imports = append(codeFile.Imports, *imp)
}

func UpdateImportStr(importText string) string {
	replaceDoubleQuote := strings.ReplaceAll(importText, "\"", "")
	replaceSingleQuote := strings.ReplaceAll(replaceDoubleQuote, "'", "")
	return replaceSingleQuote
}

func (s *TypeScriptIdentListener) EnterImportAliasDeclaration(ctx *parser.ImportAliasDeclarationContext) {
	replaceSingleQuote := UpdateImportStr(ctx.StringLiteral().GetText())
	imp := &trial.CodeImport{Source: replaceSingleQuote}
	codeFile.Imports = append(codeFile.Imports, *imp)
}

func (s *TypeScriptIdentListener) EnterImportAll(ctx *parser.ImportAllContext) {
	replaceSingleQuote := UpdateImportStr(ctx.StringLiteral().GetText())
	imp := &trial.CodeImport{Source: replaceSingleQuote}
	codeFile.Imports = append(codeFile.Imports, *imp)
}

func (s *TypeScriptIdentListener) EnterInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	currentNode = domain.NewClassNode()
	currentNode.Type = "Interface"
	currentNode.Class = ctx.Identifier().GetText()

	currentDataStruct = &trial.CodeDataStruct{
		Type: "Interface",
		Name: ctx.Identifier().GetText(),
	}

	if ctx.InterfaceExtendsClause() != nil {
		extendsContext := ctx.InterfaceExtendsClause().(*parser.InterfaceExtendsClauseContext)
		implements := BuildImplements(extendsContext.ClassOrInterfaceTypeList())
		currentNode.Extend = implements[0]

		currentDataStruct.Extend = implements[0]
	}

	objectTypeCtx := ctx.ObjectType().(*parser.ObjectTypeContext)
	if objectTypeCtx.TypeBody() != nil {
		typeMemberListCtx := objectTypeCtx.TypeBody().(*parser.TypeBodyContext).TypeMemberList().(*parser.TypeMemberListContext)
		BuildInterfaceTypeBody(typeMemberListCtx, currentNode)
	}
}

func BuildInterfaceTypeBody(ctx *parser.TypeMemberListContext, classNode *domain.JClassNode) {
	for _, typeMember := range ctx.AllTypeMember() {
		typeMemberCtx := typeMember.(*parser.TypeMemberContext)
		memberChild := typeMemberCtx.GetChild(0)
		switch x := memberChild.(type) {
		case *parser.PropertySignatureContext:
			BuildInterfacePropertySignature(x, classNode)
		case *parser.MethodSignatureContext:
			method := domain.NewJMethod()
			method.Name = x.PropertyName().GetText()
			FillMethodFromCallSignature(x.CallSignature().(*parser.CallSignatureContext), &method)

			classNode.Methods = append(classNode.Methods, method)
		}
	}
}

func BuildInterfacePropertySignature(signatureCtx *parser.PropertySignatureContext, classNode *domain.JClassNode) {
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
	} else {
		field := &domain.JField{
			Type:  typeType,
			Value: typeValue,
		}

		classNode.Fields = append(classNode.Fields, *field)
	}
}

func (s *TypeScriptIdentListener) ExitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	exitClass()
}

func (s *TypeScriptIdentListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	currentNode = domain.NewClassNode()
	currentNode.Type = "Class"
	currentNode.Class = ctx.Identifier().GetText()

	currentDataStruct = &trial.CodeDataStruct{
		Type: "Class",
		Name: ctx.Identifier().GetText(),
	}

	heritageContext := ctx.ClassHeritage().(*parser.ClassHeritageContext)
	if heritageContext.ImplementsClause() != nil {
		typeList := heritageContext.ImplementsClause().(*parser.ImplementsClauseContext).ClassOrInterfaceTypeList()

		currentNode.Implements = append(currentNode.Implements, BuildImplements(typeList)...)
		currentDataStruct.Implements  = append(currentNode.Implements, BuildImplements(typeList)...)
	}

	if heritageContext.ClassExtendsClause() != nil {
		referenceContext := heritageContext.ClassExtendsClause().(*parser.ClassExtendsClauseContext).TypeReference().(*parser.TypeReferenceContext)

		currentNode.Extend = referenceContext.TypeName().GetText()
		currentDataStruct.Extend = referenceContext.TypeName().GetText()
	}

	classTailContext := ctx.ClassTail().(*parser.ClassTailContext)
	handleClassBodyElements(classTailContext)

	classNodeQueue = append(classNodeQueue, *currentNode)
	dataStructQueue = append(dataStructQueue, *currentDataStruct)
}

func handleClassBodyElements(classTailContext *parser.ClassTailContext) {
	for _, classElement := range classTailContext.AllClassElement() {
		elementChild := classElement.GetChild(0)
		switch x := elementChild.(type) {
		case *parser.ConstructorDeclarationContext:
			constructorMethod, codeFunction := BuildConstructorMethod(x)

			currentNode.Methods = append(currentNode.Methods, constructorMethod)
			currentDataStruct.Functions = append(currentDataStruct.Functions, *codeFunction)
		case *parser.PropertyMemberDeclarationContext:
			HandlePropertyMember(x, currentNode, currentDataStruct)
		}
	}
}

func HandlePropertyMember(propertyMemberCtx *parser.PropertyMemberDeclarationContext, node *domain.JClassNode, dataStruct *trial.CodeDataStruct) {
	callSignatureSizePos := 3
	if propertyMemberCtx.PropertyName() != nil {
		field := domain.JField{}
		field.Value = propertyMemberCtx.PropertyName().GetText()
		field.Modifier = propertyMemberCtx.PropertyMemberBase().GetText()
		if propertyMemberCtx.TypeAnnotation() != nil {
			field.Type = BuildTypeAnnotation(propertyMemberCtx.TypeAnnotation().(*parser.TypeAnnotationContext))
		}
		node.Fields = append(currentNode.Fields, field)
	}

	if propertyMemberCtx.GetChildCount() >= callSignatureSizePos {
		callSignCtxPos := 2
		switch propertyMemberCtx.GetChild(callSignCtxPos).(type) {
		case *parser.CallSignatureContext:
			memberMethod, memberFunction := BuildMemberMethod(propertyMemberCtx)
			node.Methods = append(currentNode.Methods, memberMethod)

			dataStruct.Functions = append(dataStruct.Functions, *memberFunction)
		}
	}

}

func (s *TypeScriptIdentListener) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {
	exitClass()
}

func exitClass() {
	classNodes = append(classNodes, *currentNode)
	if len(classNodeQueue) > 1 {
		classNodeQueue = classNodeQueue[0 : len(classNodeQueue)-1]
		currentNode = &classNodeQueue[len(classNodeQueue)-1]
	} else {
		currentNode = domain.NewClassNode()
	}

	dataStructures = append(dataStructures, *currentDataStruct)
	if len(dataStructQueue) > 1 {
		dataStructQueue = dataStructQueue[0 : len(dataStructQueue)-1]
		currentDataStruct = &dataStructQueue[len(dataStructQueue)-1]
	} else {
		currentDataStruct = trial.NewDataStruct()
	}
}

func (s *TypeScriptIdentListener) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	method := domain.NewJMethod()

	method.Name = ctx.Identifier().GetText()
	method.AddPosition(ctx.GetChild(0).GetParent().(*antlr.BaseParserRuleContext))

	callSignatureContext := ctx.CallSignature().(*parser.CallSignatureContext)
	FillMethodFromCallSignature(callSignatureContext, &method)

	currentNode.Methods = append(currentNode.Methods, method)
}

func FillMethodFromCallSignature(callSignatureContext *parser.CallSignatureContext, method *domain.JMethod) {
	if callSignatureContext.ParameterList() != nil {
		parameterListContext := callSignatureContext.ParameterList().(*parser.ParameterListContext)
		methodParameters := BuildMethodParameter(parameterListContext)
		method.Parameters = append(method.Parameters, methodParameters...)
	}

	if callSignatureContext.TypeAnnotation() != nil {
		annotationContext := callSignatureContext.TypeAnnotation().(*parser.TypeAnnotationContext)
		method.Type = BuildTypeAnnotation(annotationContext)
	}
}
