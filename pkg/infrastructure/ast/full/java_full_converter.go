package full

import (
	"github.com/phodal/coca/languages/java"
	"github.com/phodal/coca/pkg/domain"
	"strings"
)

func BuildMethodParameters(parameters parser.IFormalParametersContext) []domain.JParameter {
	var methodParams []domain.JParameter = nil
	parameterList := parameters.GetChild(1).(*parser.FormalParameterListContext)
	formalParameter := parameterList.AllFormalParameter()
	for _, param := range formalParameter {
		paramContext := param.(*parser.FormalParameterContext)
		paramType := paramContext.TypeType().GetText()
		paramValue := paramContext.VariableDeclaratorId().(*parser.VariableDeclaratorIdContext).IDENTIFIER().GetText()

		localVars[paramValue] = paramType
		methodParams = append(methodParams, domain.JParameter{Name: paramValue, Type: paramType})
	}
	return methodParams
}

func BuildMethodCallMethods(jMethodCall *domain.JMethodCall, callee string, targetType string, ctx *parser.MethodCallContext) {
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
	}

	jMethodCall.Package = packageName
	jMethodCall.MethodName = methodName
	jMethodCall.Class = targetType
}

func BuildMethodCallLocation(jMethodCall *domain.JMethodCall, ctx *parser.MethodCallContext, callee string) {
	jMethodCall.StartLine = ctx.GetStart().GetLine()
	jMethodCall.StartLinePosition = ctx.GetStart().GetColumn()
	jMethodCall.StopLine = ctx.GetStop().GetLine()
	jMethodCall.StopLinePosition = jMethodCall.StartLinePosition + len(callee)
}

func BuildMethodCallParameters(jMethodCall *domain.JMethodCall, ctx *parser.MethodCallContext) {
	if ctx.ExpressionList() != nil {
		var parameters []string
		for _, expression := range ctx.ExpressionList().(*parser.ExpressionListContext).AllExpression() {
			expressionCtx := expression.(*parser.ExpressionContext)
			parameters = append(parameters, expressionCtx.GetText())
		}
		jMethodCall.Parameters = parameters
	}
}

