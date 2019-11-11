package api

import (
	"github.com/phodal/coca/adapter/models"
	. "github.com/phodal/coca/language/java"
	"reflect"
	"strings"
)

var clz []models.JClassNode

type RestApi struct {
	Uri            string
	HttpMethod     string
	MethodName     string
	ResponseStatus string
	Body           string
	MethodParams   map[string]string
}

var hasEnterClass = false
var isSpringRestController = false
var hasEnterRestController = false
var baseApiUrlName = ""
var localVars = make(map[string]string)

var currentRestApi RestApi
var RestApis []RestApi

func NewJavaApiListener() *JavaApiListener {
	isSpringRestController = false
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

		buildRestApi(ctx)
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
			}
		}

		paramType := paramContext.TypeType().GetText()
		paramValue := paramContext.VariableDeclaratorId().(*VariableDeclaratorIdContext).IDENTIFIER().GetText()

		if hasRequestBody {
			requestBodyClass = paramType
		}

		localVars[paramValue] = paramType
	}
	currentRestApi.Body = requestBodyClass
	//currentRestApi.Body
	hasEnterRestController = false
	requestBodyClass = ""
	RestApis = append(RestApis, currentRestApi)
}

func (s *JavaApiListener) appendClasses(classes []models.JClassNode) {
	clz = classes
}

func (s *JavaApiListener) getApis() []RestApi {
	return RestApis
}
