package ast_java

import (
	"github.com/phodal/coca/languages/java"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"strings"
)

func BuildMethodParameters(parameters parser.IFormalParametersContext) []core_domain.CodeProperty {
	var methodParams []core_domain.CodeProperty = nil
	parameterList := parameters.GetChild(1).(*parser.FormalParameterListContext)
	formalParameter := parameterList.AllFormalParameter()
	for _, param := range formalParameter {
		paramContext := param.(*parser.FormalParameterContext)
		paramType := paramContext.TypeType().GetText()
		paramValue := paramContext.VariableDeclaratorId().(*parser.VariableDeclaratorIdContext).IDENTIFIER().GetText()

		localVars[paramValue] = paramType
		parameter := core_domain.NewCodeParameter(paramType, paramValue)
		methodParams = append(methodParams, parameter)
	}
	return methodParams
}

func BuildMethodCallMethod(jMethodCall *core_domain.CodeCall, callee string, targetType string, ctx *parser.MethodCallContext) {
	methodName := callee
	packageName := currentPkg

	fullType, callType := WarpTargetFullType(targetType)
	if targetType == "super" || callee == "super" {
		callType = "super"
		targetType = currentClzExtend
	}
	jMethodCall.Type = callType

	if fullType != "" {
		packageName = RemoveTarget(fullType)
		methodName = callee
	} else {
		targetType, packageName = HandleEmptyFullType(ctx, targetType, methodName, packageName)
	}

	// TODO: 处理链试调用
	// for normal builder chain call
	if isChainCall(targetType) {
		split := strings.Split(targetType, ".")
		targetType = split[0]
		targetType = ParseTargetType(targetType)
	}

	jMethodCall.Package = packageName
	jMethodCall.FunctionName = methodName
	jMethodCall.NodeName = targetType
}

func BuildMethodCallLocation(jMethodCall *core_domain.CodeCall, ctx *parser.MethodCallContext, callee string) {
	jMethodCall.Position.StartLine = ctx.GetStart().GetLine()
	jMethodCall.Position.StartLinePosition = ctx.GetStart().GetColumn()
	jMethodCall.Position.StopLine = ctx.GetStop().GetLine()
	jMethodCall.Position.StopLinePosition = jMethodCall.Position.StartLinePosition + len(callee)
}

func BuildMethodCallParameters(jMethodCall *core_domain.CodeCall, ctx *parser.MethodCallContext) {
	if ctx.ExpressionList() != nil {
		var parameters []core_domain.CodeProperty
		for _, expression := range ctx.ExpressionList().(*parser.ExpressionListContext).AllExpression() {
			expressionCtx := expression.(*parser.ExpressionContext)

			parameter := core_domain.NewCodeParameter("", expressionCtx.GetText())
			parameters = append(parameters, parameter)
		}
		jMethodCall.Parameters = parameters
	}
}
