package ts

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/languages/ts"
	"github.com/phodal/coca/pkg/domain"
	"github.com/phodal/coca/pkg/domain/trial"
	"github.com/phodal/coca/trial/pkg/ast/ast_util"
)

func BuildArgExpressCall(memberDotExprCtx *parser.MemberDotExpressionContext) domain.JMethodCall {
	call := domain.NewJMethodCall()
	memberChild := memberDotExprCtx.GetChild(0)
	switch x := memberChild.(type) {
	case *parser.IdentifierExpressionContext:
		call.Class = x.GetText()
		call.MethodName = memberDotExprCtx.IdentifierName().GetText()
	}

	return call
}

func BuildConstructorMethod(ctx *parser.ConstructorDeclarationContext) (domain.JMethod, *trial.CodeFunction) {
	method := domain.NewJMethod()
	method.Name = "constructor"

	function := &trial.CodeFunction{
		Name: "constructor",
	}

	ast_util.AddPosition(&method, ctx.GetChild(0).GetParent().(*antlr.BaseParserRuleContext))

	if ctx.AccessibilityModifier() != nil {
		modifier := ctx.AccessibilityModifier().GetText()

		method.Modifiers = append(method.Modifiers, modifier)
		function.Modifiers = append(function.Modifiers, modifier)
	}

	return method, function
}

func BuildMemberMethod(ctx *parser.PropertyMemberDeclarationContext) (domain.JMethod, *trial.CodeFunction) {
	method := domain.NewJMethod()
	method.Name = ctx.PropertyName().GetText()

	method.StartLine = ctx.GetStart().GetLine()
	method.StartLinePosition = ctx.GetStart().GetColumn()
	method.StopLine = ctx.GetStop().GetLine()
	method.StopLinePosition = ctx.GetStop().GetColumn()

	function := &trial.CodeFunction{
		Name: ctx.PropertyName().GetText(),
	}
	function.CodePosition.StartLine = ctx.GetStart().GetLine()
	function.CodePosition.StartLinePosition = ctx.GetStart().GetColumn()
	function.CodePosition.StopLine = ctx.GetStop().GetLine()
	function.CodePosition.StopLinePosition = ctx.GetStop().GetColumn()

	return method, function
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

func BuildMethodParameter(context *parser.ParameterListContext) ([]domain.JParameter, []trial.CodeProperty) {
	childNode := context.GetChild(0)
	var parameters []domain.JParameter = nil
	var codeParameters []trial.CodeProperty = nil

	switch x := childNode.(type) {
	case *parser.RequiredParameterListContext:
		listContext := x

		list, properties := buildRequireParameterList(listContext)
		parameters = append(parameters, list...)
		codeParameters = append(codeParameters, properties...)

		if context.RestParameter() != nil {
			restParamCtx := context.RestParameter().(*parser.RestParameterContext)
			restParameters, codeProperty := buildRestParameters(restParamCtx)

			parameters = append(parameters, restParameters)
			codeParameters = append(codeParameters, codeProperty)
		}
	case *parser.PredefinedTypeContext:
		predefinedTypeContext := x
		parameters = append(parameters, domain.JParameter{
			Name: "any",
			Type: predefinedTypeContext.GetText(),
		})

		parameter := trial.CodeProperty{
			TypeName: "any",
			TypeType: predefinedTypeContext.GetText(),
		}
		codeParameters = append(codeParameters, parameter)
	}

	return parameters, codeParameters
}

func buildRestParameters(ctx *parser.RestParameterContext) (domain.JParameter, trial.CodeProperty) {
	context := ctx.GetChild(1).(*parser.RequiredParameterContext)
	return buildRequiredParameter(context)
}

func buildRequireParameterList(listContext *parser.RequiredParameterListContext) ([]domain.JParameter, []trial.CodeProperty) {
	var requireParamsList []domain.JParameter = nil
	var requireCodeParams []trial.CodeProperty = nil

	for _, requiredParameter := range listContext.AllRequiredParameter() {
		paramCtx := requiredParameter.(*parser.RequiredParameterContext)
		parameter, property := buildRequiredParameter(paramCtx)
		requireParamsList = append(requireParamsList, parameter)

		requireCodeParams = append(requireCodeParams, property)
	}
	return requireParamsList, requireCodeParams
}

func buildRequiredParameter(paramCtx *parser.RequiredParameterContext) (domain.JParameter, trial.CodeProperty) {
	paramType := ""
	if paramCtx.TypeAnnotation() != nil {
		annotationContext := paramCtx.TypeAnnotation().(*parser.TypeAnnotationContext)
		paramType = BuildTypeAnnotation(annotationContext)
	}
	parameter := domain.JParameter{
		Name: paramCtx.IdentifierOrPattern().GetText(),
		Type: paramType,
	}
	codeParamter := trial.CodeProperty{
		TypeName: paramCtx.IdentifierOrPattern().GetText(),
		TypeType: paramType,
	}

	return parameter, codeParamter
}

func BuildTypeAnnotation(annotationContext *parser.TypeAnnotationContext) string {
	return annotationContext.Type_().GetText()
}
