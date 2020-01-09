package ast

import (
	parser "github.com/phodal/coca/languages/js"
	"github.com/phodal/coca/pkg/domain"
	"reflect"
)

var currentNode domain.JClassNode

type JavaScriptIdentListener struct {
	parser.BaseJavaScriptParserListener
}

func NewJavaScriptIdentListener() *JavaScriptIdentListener {
	currentNode = *domain.NewClassNode()
	return &JavaScriptIdentListener{}
}

func (s *JavaScriptIdentListener) EnterProgram(ctx *parser.ProgramContext) {

}

func (s *JavaScriptIdentListener) EnterArgumentsExpression(ctx *parser.ArgumentsExpressionContext) {
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

func (s *JavaScriptIdentListener) EnterMemberDotExpression(ctx *parser.MemberDotExpressionContext) {

}

func (s *JavaScriptIdentListener) GetNodeInfo() domain.JClassNode {
	return currentNode
}



