package ast

import (
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

	classNodeQueue = append(classNodeQueue, *currentNode)
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
