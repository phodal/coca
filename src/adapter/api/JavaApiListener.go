package api

import (
	models2 "coca/src/adapter/models"
	"coca/src/language/java"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"strings"
)

var jClassNodes []models2.JClassNode

type RestApi struct {
	Uri              string
	HttpMethod       string
	MethodName       string
	ResponseStatus   string
	RequestBodyClass string
	MethodParams     map[string]string
	PackageName      string
	ClassName        string
}

var hasEnterClass = false
var isSpringRestController = false
var hasEnterRestController = false
var baseApiUrlName = ""
var localVars = make(map[string]string)

var currentRestApi RestApi
var RestApis []RestApi
var currentClz string
var currentPkg string

func NewJavaApiListener() *JavaApiListener {
	isSpringRestController = false
	currentClz = ""
	currentPkg = ""

	params := make(map[string]string)
	currentRestApi = *&RestApi{"", "", "", "", "", params, "", ""}
	return &JavaApiListener{}
}

type JavaApiListener struct {
	parser.BaseJavaParserListener
}

func (s *JavaApiListener) EnterPackageDeclaration(ctx *parser.PackageDeclarationContext) {
	currentPkg = ctx.QualifiedName().GetText()
}

func (s *JavaApiListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	hasEnterClass = true
	if ctx.IDENTIFIER() != nil {
		currentClz = ctx.IDENTIFIER().GetText()
	}
}

func (s *JavaApiListener) EnterAnnotation(ctx *parser.AnnotationContext) {
	annotationName := ctx.QualifiedName().GetText()
	if annotationName == "RestController" || annotationName == "Controller" {
		isSpringRestController = true
	}

	if !isSpringRestController {
		return
	}

	if !hasEnterClass {
		if annotationName == "RequestMapping" {
			if ctx.ElementValuePairs() != nil {
				firstPair := ctx.ElementValuePairs().GetChild(0).(*parser.ElementValuePairContext)
				if firstPair.IDENTIFIER().GetText() == "value" {
					baseApiUrlName = firstPair.ElementValue().GetText()
				}
			} else {
				baseApiUrlName = "/"
			}
		}
	}

	if !(annotationName == "RequestMapping" || annotationName == "GetMapping" || annotationName == "PutMapping" || annotationName == "PostMapping" || annotationName == "DeleteMapping") {
		return
	}

	hasEnterRestController = true
	uri := ""
	if ctx.ElementValue() != nil {
		uri = baseApiUrlName + ctx.ElementValue().GetText()
	} else {
		uri = baseApiUrlName
	}

	uriRemoveQuote := strings.ReplaceAll(uri, "\"", "")

	currentRestApi = RestApi{uriRemoveQuote, "", "", "", "", nil, "", ""}
	if annotationName != "RequestMapping" {
		if hasEnterClass {
			addApiMethod(annotationName)
		}
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
				currentRestApi.Uri = text[1 : len(text)-1]
			}
		}
	}
}

func addApiMethod(annotationName string) {
	switch annotationName {
	case
		"GetMapping",
		"RequestMethod.GET",
		"GET":
		currentRestApi.HttpMethod = "GET"

	case
		"PutMapping",
		"RequestMethod.PUT",
		"PUT":
		currentRestApi.HttpMethod = "PUT"

	case
		"PostMapping",
		"RequestMethod.POST",
		"POST":
		currentRestApi.HttpMethod = "POST"

	case
		"DeleteMapping",
		"RequestMethod.DELETE",
		"DELETE":

		currentRestApi.HttpMethod = "DELETE"
	}
}

var requestBodyClass string

func (s *JavaApiListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	if hasEnterRestController && ctx.FormalParameters() != nil {
		if ctx.FormalParameters().GetChild(0) == nil || ctx.FormalParameters().GetChild(1) == nil {
			return
		}

		currentRestApi.PackageName = currentPkg
		currentRestApi.ClassName = currentClz
		currentRestApi.MethodName = ctx.IDENTIFIER().GetText()
		if ctx.FormalParameters().GetText() == "()" {
			currentRestApi.RequestBodyClass = requestBodyClass
			hasEnterRestController = false
			requestBodyClass = ""
			RestApis = append(RestApis, currentRestApi)
		} else {
			buildRestApi(ctx)
		}
	}

	methodBody := ctx.MethodBody()
	blockContext := methodBody.GetChild(0)
	if reflect.TypeOf(blockContext).String() == "*parser.BlockContext" {
		filterMethodCall(blockContext)
	}
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

func buildRestApi(ctx *parser.MethodDeclarationContext) {
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

				if qualifiedName == "PathVariable" {
					//fmt.Println()
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
	currentRestApi.RequestBodyClass = requestBodyClass

	buildMethodParameters(requestBodyClass)

	hasEnterRestController = false
	requestBodyClass = ""
	RestApis = append(RestApis, currentRestApi)
}

func buildMethodParameters(requestBodyClass string) {
	params := make(map[string]string)
	for _, clz := range jClassNodes {
		if clz.Class == requestBodyClass {
			for _, field := range clz.Fields {
				params[field.Value] = field.Type
			}
		}
	}

	currentRestApi.MethodParams = params
}

func (s *JavaApiListener) appendClasses(classes []models2.JClassNode) {
	jClassNodes = classes
}

func (s *JavaApiListener) getClassApis() []RestApi {
	return RestApis
}

func (s *JavaApiListener) getCurrentApi() RestApi {
	return currentRestApi
}
