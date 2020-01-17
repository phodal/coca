package cocago

import (
	"fmt"
	. "github.com/phodal/coca/pkg/domain/core_domain"
	"go/ast"
	"reflect"
)

func BuildPropertyField(name string, field *ast.Field) *CodeProperty {
	var typeName string
	var typeType string
	var params []CodeProperty
	var results []CodeProperty
	switch x := field.Type.(type) {
	case *ast.Ident:
		typeType = "Identify"
		typeName = x.String()
	case *ast.ArrayType:
		typeType = "ArrayType"
		switch typeX := x.Elt.(type) {
		case *ast.Ident:
			typeName = typeX.String()
		case *ast.SelectorExpr:
			typeName = getSelectorName(*typeX)
		default:
			fmt.Fprintf(output, "BuildPropertyField ArrayType %s\n", reflect.TypeOf(x.Elt))
		}
	case *ast.FuncType:
		typeType = "Function"
		typeName = "func"
		if x.Params != nil {
			params = BuildFieldToProperty(x.Params.List)
		}
		if x.Results != nil {
			results = BuildFieldToProperty(x.Results.List)
		}
	case *ast.StarExpr:
		typeName = getStarExprName(*x)
		typeType = "Star"
	case *ast.SelectorExpr:
		typeName = getSelectorName(*x)
	default:
		fmt.Fprintf(output, "BuildPropertyField %s\n", reflect.TypeOf(x))
	}

	property := &CodeProperty{
		Modifiers:   nil,
		Name:        name,
		TypeType:    typeType,
		TypeValue:   typeName,
		ReturnTypes: results,
		Parameters:  params,
	}
	return property
}

func getSelectorName(typeX ast.SelectorExpr) string {
	return typeX.X.(*ast.Ident).String() + "." + typeX.Sel.Name
}

func getStarExprName(starExpr ast.StarExpr) string {
	switch x := starExpr.X.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.SelectorExpr:
		return getSelectorName(*x)
	default:
		fmt.Println("getStarExprName", reflect.TypeOf(x))
		return ""
	}
}

func BuildFunction(x *ast.FuncDecl, file *CodeFile) *CodeFunction {
	codeFunc := &CodeFunction{
		Name: x.Name.String(),
	}

	if x.Type.Params != nil {
		codeFunc.Parameters = append(codeFunc.Parameters, BuildFieldToProperty(x.Type.Params.List)...)
	}

	if x.Type.Results != nil {
		codeFunc.MultipleReturns = append(codeFunc.Parameters, BuildFieldToProperty(x.Type.Results.List)...)
	}

	fields := file.Fields
	var localVars []CodeProperty
	for _, item := range x.Body.List {
		BuildMethodCall(codeFunc, item, fields, localVars)
	}
	return codeFunc
}

func BuildFieldToProperty(fieldList []*ast.Field) []CodeProperty {
	var properties []CodeProperty
	for _, field := range fieldList {
		property := BuildPropertyField(getFieldName(field), field)
		properties = append(properties, *property)
	}
	return properties
}

func BuildMethodCall(codeFunc *CodeFunction, item ast.Stmt, fields []CodeField, localVars []CodeProperty) {
	switch it := item.(type) {
	case *ast.ExprStmt:
		BuildMethodCallExprStmt(it, codeFunc, fields)
	default:
		fmt.Fprintf(output, "methodCall %s\n", reflect.TypeOf(it))
	}
}

func BuildMethodCallExprStmt(it *ast.ExprStmt, codeFunc *CodeFunction, fields []CodeField) {
	switch expr := it.X.(type) {
	case *ast.CallExpr:
		selector, selName := BuildExpr(expr.Fun.(ast.Expr))
		target := ParseTarget(selector, fields)
		fmt.Println(target)
		call := CodeCall{
			Package:    target,
			Type:       "",
			NodeName:   selector,
			MethodName: selName,
		}

		for _, arg := range expr.Args {
			value, kind := BuildExpr(arg.(ast.Expr))
			property := &CodeProperty{
				TypeValue: value,
				TypeType:  kind,
			}

			call.Parameters = append(call.Parameters, *property)
		}

		codeFunc.MethodCalls = append(codeFunc.MethodCalls, call)
	}
}

func ParseTarget(selector string, fields []CodeField) string {
	for _, field := range fields {
		if field.TypeValue == selector {
			return field.TypeType
		}
	}
	return ""
}
