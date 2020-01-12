package ts

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/languages/ts"
	"github.com/phodal/coca/pkg/domain"
	"reflect"
)

func BuildArgExpressCall(memberDotExprCtx *parser.MemberDotExpressionContext) domain.JMethodCall {
	call := domain.NewJMethodCall()
	memberChild := memberDotExprCtx.GetChild(0)
	if reflect.TypeOf(memberChild).String() == "*parser.IdentifierExpressionContext" {
		call.Class = memberChild.(*parser.IdentifierExpressionContext).GetText()
		call.MethodName = memberDotExprCtx.IdentifierName().GetText()
	}

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

func BuildMethodParameter(context *parser.ParameterListContext) []domain.JParameter {
	childNode := context.GetChild(0)
	var parameters []domain.JParameter = nil
	switch x := childNode.(type) {
	case *parser.RequiredParameterListContext:
		listContext := x

		parameters = append(parameters, buildRequireParameterList(listContext)...)

		if context.RestParameter() != nil  {
			restParamCtx := context.RestParameter().(*parser.RestParameterContext)
			parameters = append(parameters, buildRestParameters(restParamCtx))
		}
	case *parser.PredefinedTypeContext:
		predefinedTypeContext := x
		parameters = append(parameters, domain.JParameter{
			Name: "any",
			Type: predefinedTypeContext.GetText(),
		})
	}

	return parameters
}

func buildRestParameters(ctx *parser.RestParameterContext) domain.JParameter {
	context := ctx.GetChild(1).(*parser.RequiredParameterContext)
	return buildRequiredParameter(context)
}

func buildRequireParameterList(listContext *parser.RequiredParameterListContext) []domain.JParameter {
	var requireParamsList []domain.JParameter = nil
	for _, requiredParameter := range listContext.AllRequiredParameter() {
		paramCtx := requiredParameter.(*parser.RequiredParameterContext)
		parameter := buildRequiredParameter(paramCtx)
		requireParamsList = append(requireParamsList, parameter)
	}
	return requireParamsList
}

func buildRequiredParameter(paramCtx *parser.RequiredParameterContext) domain.JParameter {
	name := paramCtx.IdentifierOrPattern().GetText()
	paramType := ""
	if paramCtx.TypeAnnotation() != nil {
		annotationContext := paramCtx.TypeAnnotation().(*parser.TypeAnnotationContext)
		paramType = BuildTypeAnnotation(annotationContext)
	}
	parameter := domain.JParameter{
		Name: name,
		Type: paramType,
	}
	return parameter
}

func BuildTypeAnnotation(annotationContext *parser.TypeAnnotationContext) string {
	return annotationContext.Type_().GetText()
}
