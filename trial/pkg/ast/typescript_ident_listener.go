package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/ts"
	"github.com/phodal/coca/pkg/domain"
	"reflect"
)

var currentNode *domain.JClassNode
var classNodeQueue []domain.JClassNode
var classNodes []domain.JClassNode

var default_class = "default"

type TypeScriptIdentListener struct {
	parser.BaseTypeScriptParserListener
}

func NewTypeScriptIdentListener() *TypeScriptIdentListener {
	classNodes = nil
	currentNode = domain.NewClassNode()
	return &TypeScriptIdentListener{}
}

func (s *TypeScriptIdentListener) GetNodeInfo() []domain.JClassNode {
	if currentNode.IsNotEmpty() {
		currentNode.Class = default_class
		classNodes = append(classNodes, *currentNode)
		currentNode = domain.NewClassNode()
	}
	return classNodes
}

func (s *TypeScriptIdentListener) EnterProgram(ctx *parser.ProgramContext) {

}

func (s *TypeScriptIdentListener) EnterInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	currentNode = domain.NewClassNode()
	currentNode.Type = "Interface"

	currentNode.Class = ctx.Identifier().GetText()

	if ctx.InterfaceExtendsClause() != nil {
		extendsContext := ctx.InterfaceExtendsClause().(*parser.InterfaceExtendsClauseContext)
		implements := BuildImplements(extendsContext.ClassOrInterfaceTypeList())
		currentNode.Extend = implements[0]
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
		currentType := reflect.TypeOf(memberChild).String()
		switch currentType {
		case "*parser.PropertySignatureContext":
			{
				signatureCtx := memberChild.(*parser.PropertySignatureContext)
				BuildInterfacePropertySignature(signatureCtx, classNode)
			}
		case "*parser.MethodSignatureContext":
			{
				methodSignature := memberChild.(*parser.MethodSignatureContext)
				method := domain.NewJMethod()
				method.Name = methodSignature.PropertyName().GetText()
				FillMethodFromCallSignature(methodSignature.CallSignature().(*parser.CallSignatureContext), &method)

				classNode.Methods = append(classNode.Methods, method)
			}
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

	heritageContext := ctx.ClassHeritage().(*parser.ClassHeritageContext)
	if heritageContext.ImplementsClause() != nil {
		typeList := heritageContext.ImplementsClause().(*parser.ImplementsClauseContext).ClassOrInterfaceTypeList()
		currentNode.Implements = append(currentNode.Implements, BuildImplements(typeList)...)
	}

	if heritageContext.ClassExtendsClause() != nil {
		referenceContext := heritageContext.ClassExtendsClause().(*parser.ClassExtendsClauseContext).TypeReference().(*parser.TypeReferenceContext)
		currentNode.Extend = referenceContext.TypeName().GetText()
	}

	classTailContext := ctx.ClassTail().(*parser.ClassTailContext)
	handleClassBodyElements(classTailContext)
	classNodeQueue = append(classNodeQueue, *currentNode)
}

func handleClassBodyElements(classTailContext *parser.ClassTailContext) {
	for _, classElement := range classTailContext.AllClassElement() {
		elementChild := classElement.GetChild(0)
		elementTypeStr := reflect.TypeOf(elementChild).String()
		switch elementTypeStr {
		case "*parser.ConstructorDeclarationContext":
			constructorDeclCtx := elementChild.(*parser.ConstructorDeclarationContext)
			currentNode.Methods = append(currentNode.Methods, BuildConstructorMethod(constructorDeclCtx))
		case "*parser.PropertyMemberDeclarationContext":
			handlePropertyMember(elementChild)
		}
	}
}

func handlePropertyMember(elementChild antlr.Tree) {
	propertyMemberCtx := elementChild.(*parser.PropertyMemberDeclarationContext)
	callSignaturePos := 3
	if propertyMemberCtx.PropertyName() != nil {
		field := domain.JField{
			Type:  "",
			Value: "",
		}
		field.Value = propertyMemberCtx.PropertyName().GetText()
		field.Modifier = propertyMemberCtx.PropertyMemberBase().GetText()
		if propertyMemberCtx.TypeAnnotation() != nil {
			field.Type = BuildTypeAnnotation(propertyMemberCtx.TypeAnnotation().(*parser.TypeAnnotationContext))
		}
		currentNode.Fields = append(currentNode.Fields, field)
	}

	if propertyMemberCtx.GetChildCount() >= callSignaturePos {
		if reflect.TypeOf(propertyMemberCtx.GetChild(2)).String() == "*parser.CallSignatureContext" {
			method := BuildMemberMethod(propertyMemberCtx)
			currentNode.Methods = append(currentNode.Methods, method)
		}
	}

}

func (s *TypeScriptIdentListener) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {
	exitClass()
}

func exitClass() {
	classNodes = append(classNodes, *currentNode)
	if len(classNodeQueue) >= 1 {
		if len(classNodeQueue) == 1 {
			currentNode = &classNodeQueue[0]
		} else {
			classNodeQueue = classNodeQueue[0 : len(classNodeQueue)-1]
			currentNode = &classNodeQueue[len(classNodeQueue)-1]
		}
	} else {
		currentNode = domain.NewClassNode()
	}
}

func (s *TypeScriptIdentListener) EnterArgumentsExpression(ctx *parser.ArgumentsExpressionContext) {
	if reflect.TypeOf(ctx.GetChild(0)).String() == "*parser.MemberDotExpressionContext" {
		call := BuildArgExpressCall(ctx.GetChild(0).(*parser.MemberDotExpressionContext))
		currentNode.MethodCalls = append(currentNode.MethodCalls, call)
	}
}

func (s *TypeScriptIdentListener) EnterMemberDotExpression(ctx *parser.MemberDotExpressionContext) {

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
