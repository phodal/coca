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
		currentNode.Implements = append(currentNode.Implements, implements...)
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

	for _, classElement := range ctx.ClassTail().(*parser.ClassTailContext).AllClassElement() {
		elementChild := classElement.GetChild(0)
		if reflect.TypeOf(elementChild).String() == "*parser.ConstructorDeclarationContext" {
			constructorDeclCtx := elementChild.(*parser.ConstructorDeclarationContext)
			currentNode.Methods = append(currentNode.Methods, BuildConstructorMethod(constructorDeclCtx))
		} else if reflect.TypeOf(elementChild).String() == "*parser.PropertyMemberDeclarationContext" {
			s.handlePropertyMember(elementChild)
		}
	}
	classNodeQueue = append(classNodeQueue, *currentNode)
}

func (s *TypeScriptIdentListener) handlePropertyMember(elementChild antlr.Tree) {
	propertyMemberCtx := elementChild.(*parser.PropertyMemberDeclarationContext)
	if propertyMemberCtx.GetChildCount() >= 3 {
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
	if callSignatureContext.ParameterList() != nil {
		parameterListContext := callSignatureContext.ParameterList().(*parser.ParameterListContext)
		methodParameters := BuildMethodParameter(parameterListContext)
		method.Parameters = append(method.Parameters, methodParameters...)
	}

	if callSignatureContext.TypeAnnotation() != nil {
		annotationContext := callSignatureContext.TypeAnnotation().(*parser.TypeAnnotationContext)
		method.Type = BuildAnnotationType(annotationContext)
	}

	currentNode.Methods = append(currentNode.Methods, method)
}
