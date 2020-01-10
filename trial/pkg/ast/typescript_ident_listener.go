package ast

import (
	parser "github.com/phodal/coca/languages/ts"
	"github.com/phodal/coca/pkg/domain"
	"reflect"
)

var currentNode *domain.JClassNode
var classNodeQueue []domain.JClassNode
var classNodes []domain.JClassNode

type TypeScriptIdentListener struct {
	parser.BaseTypeScriptParserListener
}

func NewTypeScriptIdentListener() *TypeScriptIdentListener {
	currentNode = domain.NewClassNode()
	return &TypeScriptIdentListener{}
}

func (s *TypeScriptIdentListener) EnterProgram(ctx *parser.ProgramContext) {

}

func (s *TypeScriptIdentListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	currentNode = domain.NewClassNode()
	currentNode.Class = ctx.Identifier().GetText()

	heritageContext := ctx.ClassHeritage().(*parser.ClassHeritageContext)
	if heritageContext.ImplementsClause() != nil {
		typeList := heritageContext.ImplementsClause().(*parser.ImplementsClauseContext).ClassOrInterfaceTypeList()
		typeListContext := typeList.(*parser.ClassOrInterfaceTypeListContext)
		for _, typeType := range typeListContext.AllTypeReference() {
			typeRefs := typeType.(*parser.TypeReferenceContext).TypeName().GetText()
			currentNode.Implements = append(currentNode.Implements, typeRefs)
		}
	}

	classNodeQueue = append(classNodeQueue, *currentNode)
}

func (s *TypeScriptIdentListener) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {
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
	return classNodes
}
