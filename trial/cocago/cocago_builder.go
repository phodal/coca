package cocago

import (
	"fmt"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"go/ast"
	"reflect"
)

func BuildPropertyField(name string, field *ast.Field) *core_domain.CodeProperty {
	var typeName string
	var typeType string
	var params []core_domain.CodeProperty
	var results []core_domain.CodeProperty
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

	property := &core_domain.CodeProperty{
		Modifiers:   nil,
		Name:        name,
		TypeType:    typeType,
		TypeName:    typeName,
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

func BuildFunction(x *ast.FuncDecl) *core_domain.CodeFunction {
	codeFunc := &core_domain.CodeFunction{
		Name: x.Name.String(),
	}

	if x.Type.Params != nil {
		codeFunc.Parameters = append(codeFunc.Parameters, BuildFieldToProperty(x.Type.Params.List)...)
	}

	if x.Type.Results != nil {
		codeFunc.ReturnTypes = append(codeFunc.Parameters, BuildFieldToProperty(x.Type.Results.List)...)
	}

	for _, item := range x.Body.List {
		BuildMethodCall(codeFunc, item)
	}
	return codeFunc
}

func BuildFieldToProperty(fieldList []*ast.Field) []core_domain.CodeProperty {
	var properties []core_domain.CodeProperty
	for _, field := range fieldList {
		property := BuildPropertyField(getFieldName(field), field)
		properties = append(properties, *property)
	}
	return properties
}
