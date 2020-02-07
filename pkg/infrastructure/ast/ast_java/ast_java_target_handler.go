package ast_java

import (
	"github.com/phodal/coca/languages/java"
	"reflect"
	"strings"
)

func ParseTargetType(targetCtx string) string {
	targetType := targetCtx

	//TODO: update this reflect
	typeOf := reflect.TypeOf(targetCtx).String()
	if strings.HasSuffix(typeOf, "MethodCallContext") {
		targetType = currentClz
	} else {
		fieldType := mapFields[targetCtx]
		formalType := formalParameters[targetCtx]
		localVarType := localVars[targetCtx]
		if fieldType != "" {
			targetType = fieldType
		} else if formalType != "" {
			targetType = formalType
		} else if localVarType != "" {
			targetType = localVarType
		}
	}

	return targetType
}

func WarpTargetFullType(targetType string) (string, string) {
	callType := ""
	if strings.EqualFold(currentClz, targetType) {
		callType = "self"
		return currentPkg + "." + targetType, callType
	}

	// TODO: update for array
	split := strings.Split(targetType, ".")
	str := split[0]
	pureTargetType := strings.ReplaceAll(strings.ReplaceAll(str, "[", ""), "]", "")

	if pureTargetType != "" {
		for _, imp := range imports {
			if strings.HasSuffix(imp, pureTargetType) {
				callType = "chain"
				return imp, callType
			}
		}
	}

	for _, clz := range clzs {
		if strings.HasSuffix(clz, "."+pureTargetType) {
			callType = "same package"
			return clz, callType
		}
	}

	if pureTargetType == "super" || pureTargetType == "this" {
		for _, imp := range imports {
			if strings.HasSuffix(imp, currentClzExtend) {
				callType = "super"
				return imp, callType
			}
		}
	}

	if _, ok := identMap[currentPkg+"."+targetType]; ok {
		callType = "same package 2"
		return currentPkg + "." + targetType, callType
	}

	return "", callType
}

func RemoveTarget(fullType string) string {
	split := strings.Split(fullType, ".")
	return strings.Join(split[:len(split)-1], ".")
}

func HandleEmptyFullType(ctx *parser.MethodCallContext, targetType string, methodName string, packageName string) (string, string) {
	if ctx.GetText() == targetType {
		clz := currentClz
		// 处理 static 方法，如 now()
		for _, imp := range imports {
			if strings.HasSuffix(imp, "."+methodName) {
				packageName = imp
				clz = ""
			}
		}

		targetType = clz
	} else {
		if strings.Contains(targetType, "this.") {
			targetType = buildSelfThisTarget(targetType)
		}
		//targetType = buildMethodNameForBuilder(ctx, targetType)
	}
	return targetType, packageName
}
// todo: check usecases
//func buildMethodNameForBuilder(ctx *parser.MethodCallContext, targetType string) string {
//	switch parentCtx := ctx.GetParent().(type) {
//	case *parser.ExpressionContext:
//		switch parentParentCtx := parentCtx.GetParent().(type) {
//		case *parser.VariableInitializerContext:
//			switch varDeclCtx := parentParentCtx.GetParent().(type) {
//			case *parser.VariableDeclaratorContext:
//				targetType = getTargetFromVarDecl(varDeclCtx, targetType)
//			}
//		}
//	}
//
//	return targetType
//}
//
//func getTargetFromVarDecl(ctx *parser.VariableDeclaratorContext, targetType string) string {
//	switch x := ctx.GetParent().(type) {
//	case *parser.VariableDeclaratorsContext:
//		switch parentType := x.GetParent().(type) {
//		case *parser.LocalVariableDeclarationContext:
//			{
//				targetType = parentType.TypeType().GetText()
//			}
//		}
//	}
//	return targetType
//}

