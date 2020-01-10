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

func (s *TypeScriptIdentListener) EnterProgram(ctx *parser.ProgramContext) {

}

func (s *TypeScriptIdentListener) EnterInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	currentNode = domain.NewClassNode()
	currentNode.Type = "Interface"

	currentNode.Class = ctx.Identifier().GetText()

	if ctx.InterfaceExtendsClause() != nil {
		extendsContext := ctx.InterfaceExtendsClause().(*parser.InterfaceExtendsClauseContext)
		buildImplements(extendsContext.ClassOrInterfaceTypeList())
	}
}

func (s *TypeScriptIdentListener) ExitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	exitClass()
}

func buildImplements(typeList parser.IClassOrInterfaceTypeListContext) {
	typeListContext := typeList.(*parser.ClassOrInterfaceTypeListContext)
	for _, typeType := range typeListContext.AllTypeReference() {
		typeRefs := typeType.(*parser.TypeReferenceContext).TypeName().GetText()
		currentNode.Implements = append(currentNode.Implements, typeRefs)
	}
}

func (s *TypeScriptIdentListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	currentNode = domain.NewClassNode()
	currentNode.Type = "Class"
	currentNode.Class = ctx.Identifier().GetText()

	heritageContext := ctx.ClassHeritage().(*parser.ClassHeritageContext)
	if heritageContext.ImplementsClause() != nil {
		typeList := heritageContext.ImplementsClause().(*parser.ImplementsClauseContext).ClassOrInterfaceTypeList()
		buildImplements(typeList)
	}

	for _, classElement := range ctx.ClassTail().(*parser.ClassTailContext).AllClassElement() {
		elementChild := classElement.GetChild(0)
		if reflect.TypeOf(elementChild).String() == "*parser.ConstructorDeclarationContext" {
			constructorDeclCtx := elementChild.(*parser.ConstructorDeclarationContext)
			appendConstructorMethod(constructorDeclCtx)
		} else if reflect.TypeOf(elementChild).String() == "*parser.PropertyMemberDeclarationContext"{
			propertyMemberCtx := elementChild.(*parser.PropertyMemberDeclarationContext)
			if propertyMemberCtx.GetChildCount() >= 3 {
				if reflect.TypeOf(propertyMemberCtx.GetChild(2)).String() == "*parser.CallSignatureContext" {
					appendNormalMethod(propertyMemberCtx)
				}
			}
		}
	}
	classNodeQueue = append(classNodeQueue, *currentNode)
}

func appendNormalMethod(ctx *parser.PropertyMemberDeclarationContext) {
	method := domain.NewJMethod()
	method.Name = ctx.PropertyName().GetText()

	method.StartLine = ctx.GetStart().GetLine()
	method.StartLinePosition = ctx.GetStart().GetColumn()
	method.StopLine = ctx.GetStop().GetLine()
	method.StopLinePosition = ctx.GetStop().GetColumn()

	currentNode.Methods = append(currentNode.Methods, method)
}

func appendConstructorMethod(ctx *parser.ConstructorDeclarationContext) {
	method := domain.NewJMethod()
	method.Name = "constructor"

	method.AddPosition(ctx.GetChild(0).GetParent().(*antlr.BaseParserRuleContext))

	if ctx.AccessibilityModifier() != nil  {
		method.Modifiers = append(method.Modifiers, ctx.AccessibilityModifier().GetText())
	}

	currentNode.Methods = append(currentNode.Methods, method)
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
	if currentNode.Class == "" {
		currentNode.Class = default_class
	}

	if reflect.TypeOf(ctx.GetChild(0)).String() == "*parser.MemberDotExpressionContext" {
		memberDotExprCtx := ctx.GetChild(0).(*parser.MemberDotExpressionContext)
		buildMemberDotExpr(memberDotExprCtx)
	}
}

func buildMemberDotExpr(memberDotExprCtx *parser.MemberDotExpressionContext) {
	call := domain.NewJMethodCall()
	call.Class = memberDotExprCtx.GetChild(0).(*parser.IdentifierExpressionContext).GetText()
	call.MethodName = memberDotExprCtx.IdentifierName().GetText()

	currentNode.MethodCalls = append(currentNode.MethodCalls, call)
}

func (s *TypeScriptIdentListener) EnterMemberDotExpression(ctx *parser.MemberDotExpressionContext) {

}

func (s *TypeScriptIdentListener) GetNodeInfo() []domain.JClassNode {
	if currentNode.Class == default_class {
		classNodes = append(classNodes, *currentNode)
		currentNode = domain.NewClassNode()
	}
	return classNodes
}
