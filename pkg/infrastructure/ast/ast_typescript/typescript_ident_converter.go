package ast_typescript

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/languages/ts"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/astutil"
)

func BuildConstructorMethod(ctx *parser.ConstructorDeclarationContext) *core_domain.CodeFunction {
	function := &core_domain.CodeFunction{
		Name: "constructor",
	}

	astutil.AddFunctionPosition(function, ctx.GetChild(0).GetParent().(*antlr.BaseParserRuleContext))

	if ctx.AccessibilityModifier() != nil {
		modifier := ctx.AccessibilityModifier().GetText()

		function.Modifiers = append(function.Modifiers, modifier)
	}

	return function
}

func BuildMemberMethod(ctx *parser.PropertyMemberDeclarationContext) *core_domain.CodeFunction {
	function := &core_domain.CodeFunction{
		Name: ctx.PropertyName().GetText(),
	}
	function.Position.StartLine = ctx.GetStart().GetLine()
	function.Position.StartLinePosition = ctx.GetStart().GetColumn()
	function.Position.StopLine = ctx.GetStop().GetLine()
	function.Position.StopLinePosition = ctx.GetStop().GetColumn()

	return function
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

func BuildMethodParameter(context *parser.ParameterListContext) ([]core_domain.CodeProperty) {
	childNode := context.GetChild(0)
	var parameters []core_domain.CodeProperty = nil

	switch x := childNode.(type) {
	case *parser.RequiredParameterListContext:
		listContext := x

		properties := buildRequireParameterList(listContext)
		parameters = append(parameters, properties...)

		if context.RestParameter() != nil {
			restParamCtx := context.RestParameter().(*parser.RestParameterContext)
			codeProperty := buildRestParameters(restParamCtx)

			parameters = append(parameters, codeProperty)
		}
	case *parser.PredefinedTypeContext:
		predefinedTypeContext := x
		parameter := core_domain.CodeProperty{
			TypeValue: "any",
			TypeType:  predefinedTypeContext.GetText(),
		}
		parameters = append(parameters, parameter)
	}

	return parameters
}

func buildRestParameters(ctx *parser.RestParameterContext) core_domain.CodeProperty {
	context := ctx.GetChild(1).(*parser.RequiredParameterContext)
	return buildRequiredParameter(context)
}

func buildRequireParameterList(listContext *parser.RequiredParameterListContext) []core_domain.CodeProperty {
	var requireCodeParams []core_domain.CodeProperty = nil

	for _, requiredParameter := range listContext.AllRequiredParameter() {
		paramCtx := requiredParameter.(*parser.RequiredParameterContext)
		property := buildRequiredParameter(paramCtx)

		requireCodeParams = append(requireCodeParams, property)
	}
	return requireCodeParams
}

func buildRequiredParameter(paramCtx *parser.RequiredParameterContext) core_domain.CodeProperty {
	paramType := ""
	if paramCtx.TypeAnnotation() != nil {
		annotationContext := paramCtx.TypeAnnotation().(*parser.TypeAnnotationContext)
		paramType = BuildTypeAnnotation(annotationContext)
	}
	parameter := core_domain.CodeProperty{
		TypeValue: paramCtx.IdentifierOrPattern().GetText(),
		TypeType:  paramType,
	}

	return parameter
}

func BuildTypeAnnotation(annotationContext *parser.TypeAnnotationContext) string {
	return annotationContext.Type_().GetText()
}
