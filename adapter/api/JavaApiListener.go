package api

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/adapter/models"
	. "github.com/phodal/coca/language/java"
	"reflect"
	"strings"
)

var jClassNodes []models.JClassNode

type RestApi struct {
	Uri              string
	HttpMethod       string
	MethodName       string
	ResponseStatus   string
	RequestBodyClass string
	MethodParams     map[string]string
}

var hasEnterClass = false
var isSpringRestController = false
var hasEnterRestController = false
var baseApiUrlName = ""
var localVars = make(map[string]string)

var currentRestApi RestApi
var RestApis []RestApi
var pathVars = make(map[string]string)

func NewJavaApiListener() *JavaApiListener {
	isSpringRestController = false
	params := make(map[string]string)
	currentRestApi = *&RestApi{"", "", "", "", "", params}
	return &JavaApiListener{}
}

type JavaApiListener struct {
	BaseJavaParserListener
}

func (s *JavaApiListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {
	hasEnterClass = true
}

func (s *JavaApiListener) EnterAnnotation(ctx *AnnotationContext) {
	annotationName := ctx.QualifiedName().GetText()
	if annotationName == "RestController" {
		isSpringRestController = true
	}

	if !isSpringRestController {
		return
	}

	if !hasEnterClass {
		if annotationName == "RequestMapping" {
			if ctx.ElementValuePairs() != nil {
				firstPair := ctx.ElementValuePairs().GetChild(0).(*ElementValuePairContext)
				if firstPair.IDENTIFIER().GetText() == "value" {
					baseApiUrlName = firstPair.ElementValue().GetText()
				}
			} else {
				baseApiUrlName = "/"
			}
		}
	}

	if !(annotationName == "GetMapping" || annotationName == "PutMapping" || annotationName == "PostMapping" || annotationName == "DeleteMapping") {
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

	currentRestApi = RestApi{uriRemoveQuote, "", "", "", "", nil}
	if hasEnterClass {
		switch annotationName {
		case "GetMapping":
			currentRestApi.HttpMethod = "GET"
		case "PutMapping":
			currentRestApi.HttpMethod = "PUT"
		case "PostMapping":
			currentRestApi.HttpMethod = "POST"
		case "DeleteMapping":
			currentRestApi.HttpMethod = "DELETE"
		}
	}
}

var requestBodyClass string

func (s *JavaApiListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {
	if hasEnterRestController && ctx.FormalParameters() != nil {
		if ctx.FormalParameters().GetChild(0) == nil || ctx.FormalParameters().GetText() == "()" || ctx.FormalParameters().GetChild(1) == nil {
			return
		}

		currentRestApi.MethodName = ctx.IDENTIFIER().GetText()
		buildRestApi(ctx)
	}

	methodBody := ctx.MethodBody()
	blockContext := methodBody.GetChild(0)
	if reflect.TypeOf(blockContext).String() == "*parser.BlockContext" {
		filterMethodCall(blockContext)
	}
}

func filterMethodCall(blockContext antlr.Tree) {
	blcStatement := blockContext.(*BlockContext).AllBlockStatement()
	for _, rangeStatement := range blcStatement {
		if reflect.TypeOf(rangeStatement.GetChild(0)).String() == "*parser.StatementContext" {
			statement := rangeStatement.GetChild(0).(*StatementContext)
			if reflect.TypeOf(statement.GetChild(0)).String() == "*parser.ExpressionContext" {
				express := statement.GetChild(0).(*ExpressionContext)
				reflect.TypeOf(express.GetChild(0))
			}
		}
	}
}

func buildRestApi(ctx *MethodDeclarationContext) {
	parameterList := ctx.FormalParameters().GetChild(1).(*FormalParameterListContext)
	formalParameter := parameterList.AllFormalParameter()
	for _, param := range formalParameter {
		paramContext := param.(*FormalParameterContext)

		modifiers := paramContext.AllVariableModifier()
		hasRequestBody := false
		for _, modifier := range modifiers {
			childType := reflect.TypeOf(modifier.GetChild(0))
			if childType.String() == "*parser.AnnotationContext" {
				qualifiedName := modifier.GetChild(0).(*AnnotationContext).QualifiedName().GetText()
				if qualifiedName == "RequestBody" {
					hasRequestBody = true
				}

				if qualifiedName == "PathVariable" {
					fmt.Println()
				}
			}
		}

		paramType := paramContext.TypeType().GetText()
		paramValue := paramContext.VariableDeclaratorId().(*VariableDeclaratorIdContext).IDENTIFIER().GetText()

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

func (s *JavaApiListener) appendClasses(classes []models.JClassNode) {
	jClassNodes = classes
}

func (s *JavaApiListener) getClassApis() []RestApi {
	return RestApis
}
