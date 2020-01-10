package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/languages/ts"
	"github.com/phodal/coca/pkg/domain"
)

func BuildArgExpressCall(memberDotExprCtx *parser.MemberDotExpressionContext) domain.JMethodCall {
	call := domain.NewJMethodCall()
	call.Class = memberDotExprCtx.GetChild(0).(*parser.IdentifierExpressionContext).GetText()
	call.MethodName = memberDotExprCtx.IdentifierName().GetText()

	return call
}

func BuildConstructorMethod(ctx *parser.ConstructorDeclarationContext) domain.JMethod {
	method := domain.NewJMethod()
	method.Name = "constructor"

	method.AddPosition(ctx.GetChild(0).GetParent().(*antlr.BaseParserRuleContext))

	if ctx.AccessibilityModifier() != nil {
		method.Modifiers = append(method.Modifiers, ctx.AccessibilityModifier().GetText())
	}

	return method
}

func BuildMemberMethod(ctx *parser.PropertyMemberDeclarationContext) domain.JMethod {
	method := domain.NewJMethod()
	method.Name = ctx.PropertyName().GetText()

	method.StartLine = ctx.GetStart().GetLine()
	method.StartLinePosition = ctx.GetStart().GetColumn()
	method.StopLine = ctx.GetStop().GetLine()
	method.StopLinePosition = ctx.GetStop().GetColumn()

	return method
}

func BuildImplements(typeList parser.IClassOrInterfaceTypeListContext) []string {
	typeListContext := typeList.(*parser.ClassOrInterfaceTypeListContext)

	var implements []string = nil
	for _, typeType := range typeListContext.AllTypeReference() {
		typeRefs := typeType.(*parser.TypeReferenceContext).TypeName().GetText()
		implements = append(implements, typeRefs)
	}

	return implements
}

