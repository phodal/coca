package ast_api_java

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/languages/java"
	api_domain2 "github.com/phodal/coca/pkg/domain/api_domain"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"reflect"
	"strings"
)

var jClassNodes []core_domain.CodeDataStruct

var hasEnterClass = false
var isSpringRestController = false
var hasEnterRestController = false
var baseApiUrl string
var localVars = make(map[string]string)

var currentRestAPI api_domain2.RestAPI
var restAPIs []api_domain2.RestAPI
var currentClz string
var currentPkg string

var identMap map[string]core_domain.CodeDataStruct
var imports []string
var currentImplements = ""

func NewJavaAPIListener(jIdentMap map[string]core_domain.CodeDataStruct, diMap map[string]string) *JavaAPIListener {
	isSpringRestController = false
	currentClz = ""
	currentPkg = ""
	currentImplements = ""

	imports = nil
	restAPIs = nil

	identMap = jIdentMap

	params := make(map[string]string)
	currentRestAPI = api_domain2.RestAPI{MethodParams: params}
	return &JavaAPIListener{}
}

type JavaAPIListener struct {
	parser.BaseJavaParserListener
}

func (s *JavaAPIListener) EnterImportDeclaration(ctx *parser.ImportDeclarationContext) {
	importText := ctx.QualifiedName().GetText()
	imports = append(imports, importText)
}

func (s *JavaAPIListener) EnterPackageDeclaration(ctx *parser.PackageDeclarationContext) {
	currentPkg = ctx.QualifiedName().GetText()
}

func (s *JavaAPIListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	hasEnterClass = true
	if ctx.IDENTIFIER() != nil {
		currentClz = ctx.IDENTIFIER().GetText()
	}

	if ctx.IMPLEMENTS() != nil {
		currentImplements = ctx.TypeList().GetText()
	}
}

func (s *JavaAPIListener) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {
	hasEnterClass = false
}

func (s *JavaAPIListener) EnterAnnotation(ctx *parser.AnnotationContext) {
	annotationName := ctx.QualifiedName().GetText()
	if annotationName == "RestController" || annotationName == "Controller" {
		isSpringRestController = true
	}

	if !isSpringRestController {
		return
	}

	if !hasEnterClass {
		buildBaseApiUrlString(annotationName, ctx)
	}

	notAPI := annotationName == "RequestMapping" || annotationName == "GetMapping" || annotationName == "PutMapping" || annotationName == "PostMapping" || annotationName == "DeleteMapping"
	if !notAPI {
		return
	}

	hasEnterRestController = true
	uri := ""
	if ctx.ElementValue() != nil {
		uri = baseApiUrl + ctx.ElementValue().GetText()
	} else {
		uri = baseApiUrl
	}

	uriRemoveQuote := strings.ReplaceAll(uri, "\"", "")

	currentRestAPI = api_domain2.RestAPI{Uri: uriRemoveQuote}
	if annotationName != "RequestMapping" {
		if hasEnterClass {
			addApiMethod(annotationName)
		}

		return
	}

	if ctx.ElementValuePairs() != nil {
		allValuePair := ctx.ElementValuePairs().(*parser.ElementValuePairsContext).AllElementValuePair()
		for _, valuePair := range allValuePair {
			pair := valuePair.(*parser.ElementValuePairContext)
			if pair.IDENTIFIER().GetText() == "method" {
				addApiMethod(pair.ElementValue().GetText())
			}
			if pair.IDENTIFIER().GetText() == "value" {
				text := pair.ElementValue().GetText()
				currentRestAPI.Uri = baseApiUrl + text[1:len(text)-1]
			}
		}
	}
}

func buildBaseApiUrlString(annotationName string, ctx *parser.AnnotationContext) {
	// 类声明处的注解
	if annotationName == "RequestMapping" {
		if ctx.ElementValuePairs() != nil {
			allValuePair := ctx.ElementValuePairs().(*parser.ElementValuePairsContext).AllElementValuePair()
			for _, valuePair := range allValuePair {
				pair := valuePair.(*parser.ElementValuePairContext)
				if pair.IDENTIFIER().GetText() == "value" {
					text := pair.ElementValue().GetText()
					baseApiUrl = text[1 : len(text)-1]
				}
			}
		} else if ctx.ElementValue() != nil {
			text := ctx.ElementValue().GetText()
			baseApiUrl = text[1 : len(text)-1]
		} else {
			baseApiUrl = "/"
		}
	}
}

func addApiMethod(annotationName string) {
	switch annotationName {
	case
		"GetMapping",
		"RequestMethod.GET",
		"GET":
		currentRestAPI.HttpMethod = "GET"

	case
		"PutMapping",
		"RequestMethod.PUT",
		"PUT":
		currentRestAPI.HttpMethod = "PUT"

	case
		"PostMapping",
		"RequestMethod.POST",
		"POST":
		currentRestAPI.HttpMethod = "POST"

	case
		"DeleteMapping",
		"RequestMethod.DELETE",
		"DELETE":

		currentRestAPI.HttpMethod = "DELETE"
	}
}

var requestBodyClass string

func (s *JavaAPIListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	methodName := ctx.IDENTIFIER().GetText()

	if currentImplements != "" {
		if buildApiForInterfaceAnnotation(methodName) {
			return
		}
	}

	if hasEnterRestController && ctx.FormalParameters() != nil {
		if ctx.FormalParameters().GetChild(0) == nil || ctx.FormalParameters().GetChild(1) == nil {
			return
		}

		currentRestAPI.PackageName = currentPkg
		currentRestAPI.ClassName = currentClz
		currentRestAPI.MethodName = methodName
		if ctx.FormalParameters().GetText() == "()" {
			currentRestAPI.RequestBodyClass = requestBodyClass
			hasEnterRestController = false
			requestBodyClass = ""
			restAPIs = append(restAPIs, currentRestAPI)
		} else {
			buildRestApiWithParameters(ctx)
		}
	}

	methodBody := ctx.MethodBody()
	blockContext := methodBody.GetChild(0)
	if reflect.TypeOf(blockContext).String() == "*parser.BlockContext" {
		filterMethodCall(blockContext)
	}
}

func buildApiForInterfaceAnnotation(methodName string) bool {
	var superClz = ""
	for index := range imports {
		imp := imports[index]
		if strings.HasSuffix(imp, "."+currentImplements) {
			superClz = imp
		}
		// TODO: 支持 interface 在同一个包内
	}

	if _, ok := identMap[superClz]; ok {
		for _, method := range identMap[superClz].Functions {
			if method.Name == methodName {
				for _, annotation := range method.Annotations {
					if annotation.Name == "ServiceMethod" {
						currentRestAPI.PackageName = currentPkg
						currentRestAPI.ClassName = currentClz
						currentRestAPI.MethodName = methodName

						restAPIs = append(restAPIs, currentRestAPI)
						return true
					}
				}
			}
		}
	}
	return false
}

func filterMethodCall(blockContext antlr.Tree) {
	blcStatement := blockContext.(*parser.BlockContext).AllBlockStatement()
	for _, rangeStatement := range blcStatement {
		if reflect.TypeOf(rangeStatement.GetChild(0)).String() == "*parser.StatementContext" {
			statement := rangeStatement.GetChild(0).(*parser.StatementContext)
			if reflect.TypeOf(statement.GetChild(0)).String() == "*parser.ExpressionContext" {
				express := statement.GetChild(0).(*parser.ExpressionContext)
				reflect.TypeOf(express.GetChild(0))
			}
		}
	}
}

func buildRestApiWithParameters(ctx *parser.MethodDeclarationContext) {
	parameterList := ctx.FormalParameters().GetChild(1).(*parser.FormalParameterListContext)
	formalParameter := parameterList.AllFormalParameter()
	for _, param := range formalParameter {
		paramContext := param.(*parser.FormalParameterContext)

		modifiers := paramContext.AllVariableModifier()
		hasRequestBody := false
		for _, modifier := range modifiers {
			childType := reflect.TypeOf(modifier.GetChild(0))
			if childType.String() == "*parser.AnnotationContext" {
				qualifiedName := modifier.GetChild(0).(*parser.AnnotationContext).QualifiedName().GetText()
				if qualifiedName == "RequestBody" {
					hasRequestBody = true
				}
			}
		}

		paramType := paramContext.TypeType().GetText()
		paramValue := paramContext.VariableDeclaratorId().(*parser.VariableDeclaratorIdContext).IDENTIFIER().GetText()

		if hasRequestBody {
			requestBodyClass = paramType
		}

		localVars[paramValue] = paramType
	}
	currentRestAPI.RequestBodyClass = requestBodyClass

	buildMethodParameters(requestBodyClass)

	hasEnterRestController = false
	requestBodyClass = ""
	restAPIs = append(restAPIs, currentRestAPI)
}

func buildMethodParameters(requestBodyClass string) {
	params := make(map[string]string)
	for _, clz := range jClassNodes {
		if clz.NodeName == requestBodyClass {
			for _, field := range clz.Fields {
				params[field.TypeValue] = field.TypeType
			}
		}
	}

	currentRestAPI.MethodParams = params
}

func (s *JavaAPIListener) AppendClasses(classes []core_domain.CodeDataStruct) {
	jClassNodes = classes
}

func (s *JavaAPIListener) GetClassApis() []api_domain2.RestAPI {
	return restAPIs
}
