package ast

import (
	parser "github.com/phodal/coca/languages/ts"
	"github.com/phodal/coca/pkg/domain"
	"reflect"
)

var currentNode domain.JClassNode

type TypeScriptIdentListener struct {
	parser.BaseTypeScriptParserListener
}

func NewTypeScriptIdentListener() *TypeScriptIdentListener {
	currentNode = *domain.NewClassNode()
	return &TypeScriptIdentListener{}
}

func (s *TypeScriptIdentListener) EnterProgram(ctx *parser.ProgramContext) {

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

func (s *TypeScriptIdentListener) GetNodeInfo() domain.JClassNode {
	return currentNode
}



