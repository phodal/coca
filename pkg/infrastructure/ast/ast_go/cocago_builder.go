package ast_go

import (
	"fmt"
	. "github.com/phodal/coca/pkg/domain/core_domain"
	"go/ast"
	"reflect"
	"strings"
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
	case *ast.InterfaceType:
		typeType = "interface{}" // todo: need check for all interface
		typeName = "interface{}"
	default:
		fmt.Fprintf(output, "BuildPropertyField %s\n", reflect.TypeOf(x))
	}

	property := &CodeProperty{
		Modifiers:   nil,
		ParamName:   name,
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

func BuildFunction(x *ast.FuncDecl, file *CodeContainer) *CodeFunction {
	codeFunc := &CodeFunction{
		Name: x.Name.String(),
	}

	if x.Type.Params != nil {
		parameters := BuildFieldToProperty(x.Type.Params.List)
		codeFunc.Parameters = append(codeFunc.Parameters, parameters...)
	}

	if x.Type.Results != nil {
		codeFunc.MultipleReturns = append(codeFunc.Parameters, BuildFieldToProperty(x.Type.Results.List)...)
	}

	fields := file.Fields
	var localVars []CodeProperty
	for _, param := range codeFunc.Parameters {
		localVars = append(localVars, CodeProperty{
			TypeType:  param.ParamName,
			TypeValue: param.TypeValue,
		})
	}
	for _, item := range x.Body.List {
		localVars, _ = BuildMethodCall(codeFunc, item, fields, localVars, file.Imports, file.PackageName)
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

func BuildMethodCall(codeFunc *CodeFunction, item ast.Stmt, fields []CodeField, localVars []CodeProperty, imports []CodeImport, packageName string) ([]CodeProperty, CodeCall) {
	var call CodeCall
	switch it := item.(type) {
	case *ast.ExprStmt:
		BuildMethodCallExprStmt(it, codeFunc, fields, imports, packageName, localVars)
	case *ast.DeferStmt:
		call = BuildCallFromExpr(it.Call, codeFunc, fields, imports, packageName, localVars)
		codeFunc.FunctionCalls = append(codeFunc.FunctionCalls, call)
	case *ast.AssignStmt:
		vars := BuildLocalVars(it, codeFunc, imports)
		localVars = vars
	case *ast.IfStmt:
	case *ast.ReturnStmt:
		for _, resultExpr := range it.Results {
			typ, caller, callee := BuildExpr(resultExpr)
			if typ == "call" {
				var calls []CodeCall
				for _, param := range codeFunc.Parameters {
					if param.ParamName == caller {
						target := ParseTarget(caller, fields, localVars, codeFunc)
						packageName := getPackageName(target, callee, imports)

						call.Package = packageName
						call.FunctionName = callee
						call.NodeName = target
						calls = append(calls, call)
					}
				}

				codeFunc.FunctionCalls = append(codeFunc.FunctionCalls, calls...)
			}
		}
	default:
		fmt.Fprintf(output, "methodCall %s\n", reflect.TypeOf(it))
	}

	return localVars, call
}

func BuildLocalVars(it *ast.AssignStmt, codeFunc *CodeFunction, imports []CodeImport) []CodeProperty {
	var vars []CodeProperty
	for _, lh := range it.Lhs {
		var left string
		switch lhx := lh.(type) {
		case *ast.Ident:
			left = lhx.Name
		}

		for _, expr := range it.Rhs {
			typ, exprName, kind := BuildExpr(expr)
			property := CodeProperty{
				TypeValue: left,
				TypeType:  kind,
			}

			vars = append(vars, property)

			if typ == "call" {
				packageName := getPackageName(exprName, "", imports)
				if packageName != "" {
					call := CodeCall{
						Package:  packageName,
						Type:     "",
						NodeName: exprName,
					}

					codeFunc.FunctionCalls = append(codeFunc.FunctionCalls, call)
				}
			}
		}
	}
	return vars
}

func BuildMethodCallExprStmt(it *ast.ExprStmt, codeFunc *CodeFunction, fields []CodeField, imports []CodeImport, currentPackage string, localVars []CodeProperty) {
	switch expr := it.X.(type) {
	case *ast.CallExpr:
		call := BuildCallFromExpr(expr, codeFunc, fields, imports, currentPackage, localVars)
		codeFunc.FunctionCalls = append(codeFunc.FunctionCalls, call)
	default:
		fmt.Fprintf(output, "BuildMethodCallExprStmt: %s\n", reflect.TypeOf(expr))
	}
}

func BuildCallFromExpr(expr *ast.CallExpr, codeFunc *CodeFunction, fields []CodeField, imports []CodeImport, currentPackage string, localVars []CodeProperty) CodeCall {
	_, selector, selName := BuildExpr(expr.Fun.(ast.Expr))
	target := ParseTarget(selector, fields, localVars, codeFunc)
	packageName := getPackageName(target, selector, imports)
	if packageName == "" {
		packageName = currentPackage
	}

	call := CodeCall{
		Package:      packageName,
		Type:         target,
		NodeName:     selector,
		FunctionName: selName,
	}

	for _, arg := range expr.Args {
		if reflect.TypeOf(arg.(ast.Expr)).String() == "*ast.FuncLit" {
			funcLit := arg.(ast.Expr).(*ast.FuncLit)
			for _, item := range funcLit.Body.List {
				_, methodCall := BuildMethodCall(codeFunc, item, fields, localVars, imports, packageName)

				if methodCall.NodeName != "" {
					codeFunc.FunctionCalls = append(codeFunc.FunctionCalls, methodCall)
				}
			}
		} 

		_, value, kind := BuildExpr(arg.(ast.Expr))
		property := &CodeProperty{
			TypeValue: value,
			TypeType:  kind,
		}

		call.Parameters = append(call.Parameters, *property)
	}
	return call
}

func getPackageName(target string, selector string, imports []CodeImport) string {
	packageName := ""
	if strings.Contains(target, ".") {
		split := strings.Split(target, ".")
		for _, imp := range imports {
			if strings.HasSuffix(imp.Source, split[0]) {
				return split[0]
			}
		}
	}

	for _, imp := range imports {
		if imp.Source == target {
			return target
		}
	}

	for _, member := range identCodeMembers {
		if member.DataStructID == target {
			return member.AliasPackage
		}
	}

	packageName = currentPackage.Name
	return packageName
}

func ParseTarget(selector string, fields []CodeField, localVars []CodeProperty, codeFunc *CodeFunction) string {
	if codeFunc != nil {
		for _, param := range codeFunc.Parameters {
			if param.ParamName == selector {
				return param.TypeValue
			}
		}
	}
	for _, localVar := range localVars {
		if selector == localVar.TypeValue {
			return localVar.TypeType
		}
	}

	for _, field := range fields {
		if field.TypeValue == selector {
			return field.TypeType
		}
	}

	return selector
}
