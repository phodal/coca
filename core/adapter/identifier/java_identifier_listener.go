package identifier

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/adapter/common_listener"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/languages/java"
	"reflect"
	"strings"
)

var currentNode *models.JIdentifier
var nodes []models.JIdentifier

var currentMethod models.JMethod
var hasEnterClass = false
var imports []string

func NewJavaIdentifierListener() *JavaIdentifierListener {
	nodes = nil
	currentNode = models.NewJIdentifier()
	currentMethod = models.NewJMethod()
	return &JavaIdentifierListener{}
}

type JavaIdentifierListener struct {
	parser.BaseJavaParserListener
}

func (s *JavaIdentifierListener) EnterImportDeclaration(ctx *parser.ImportDeclarationContext) {
	importText := ctx.QualifiedName().GetText()
	imports = append(imports, importText)
}

func (s *JavaIdentifierListener) EnterPackageDeclaration(ctx *parser.PackageDeclarationContext) {
	currentNode.Package = ctx.QualifiedName().GetText()
}

func (s *JavaIdentifierListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	hasEnterClass = true

	currentNode.ClassType = "Class"
	if ctx.IDENTIFIER() != nil {
		currentNode.ClassName = ctx.IDENTIFIER().GetText()
	}

	if ctx.EXTENDS() != nil {
		currentNode.ExtendsName = ctx.TypeType().GetText()
	}

	if ctx.IMPLEMENTS() != nil {
		types := ctx.TypeList().(*parser.TypeListContext).AllTypeType()
		for _, typ := range types {
			typeText := typ.GetText()
			for _, imp := range imports {
				if strings.HasSuffix(imp, "."+typeText) {
					currentNode.Implements = append(currentNode.Implements, imp)
				}
			}
		}
	}

	currentMethod = models.NewJMethod()
}

func (s *JavaIdentifierListener) ExitClassBody(ctx *parser.ClassBodyContext) {
	hasEnterClass = false
	if currentNode.ClassName != "" {
		currentNode.Methods = currentNode.GetMethods()
		nodes = append(nodes, *currentNode)
	}
	currentNode = models.NewJIdentifier()
}

func (s *JavaIdentifierListener) ExitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	hasEnterClass = false
	if currentNode.ClassName != "" {
		currentNode.Methods = currentNode.GetMethods()
		nodes = append(nodes, *currentNode)
	}
	currentNode = models.NewJIdentifier()
}

func (s *JavaIdentifierListener) EnterConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) {

	currentMethod = *&models.JMethod{
		Name:              ctx.IDENTIFIER().GetText(),
		Type:              "",
		StartLine:         ctx.GetStart().GetLine(),
		StartLinePosition: ctx.GetStart().GetColumn(),
		StopLine:          ctx.GetStop().GetLine(),
		StopLinePosition:  ctx.GetStop().GetColumn(),
		Override:          isOverrideMethod,
		Annotations:       currentMethod.Annotations,
		IsConstructor:     true,
	}
}

func (s *JavaIdentifierListener) ExitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) {

	currentNode.AddMethod(currentMethod)
	_ = models.NewJMethod()
}

func (s *JavaIdentifierListener) EnterInterfaceBodyDeclaration(ctx *parser.InterfaceBodyDeclarationContext) {
	hasEnterClass = true
	for _, modifier := range ctx.AllModifier() {
		modifier := modifier.(*parser.ModifierContext).GetChild(0)
		if reflect.TypeOf(modifier.GetChild(0)).String() == "*parser.AnnotationContext" {
			annotationContext := modifier.GetChild(0).(*parser.AnnotationContext)
			common_listener.BuildAnnotation(annotationContext)
		}
	}
}

func (s *JavaIdentifierListener) EnterInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().GetColumn()
	name := ctx.IDENTIFIER().GetText()
	//XXX: find the start position of {, not public
	typeType := ctx.TypeTypeOrVoid().GetText()

	annotations := currentMethod.Annotations
	currentMethod = *&models.JMethod{
		Name:              name,
		Type:              typeType,
		StartLine:         startLine,
		StartLinePosition: startLinePosition,
		StopLine:          stopLine,
		StopLinePosition:  stopLinePosition,
		Override:          isOverrideMethod,
		Annotations:       annotations,
	}
}

func (s *JavaIdentifierListener) ExitInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) {

	currentNode.AddMethod(currentMethod)
	_ = models.NewJMethod()
}

var isOverrideMethod = false

func (s *JavaIdentifierListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	hasEnterClass = true

	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().GetColumn()
	name := ctx.IDENTIFIER().GetText()

	typeType := ctx.TypeTypeOrVoid().GetText()

	annotations := currentMethod.Annotations
	currentMethod = *&models.JMethod{
		Name:              name,
		Type:              typeType,
		StartLine:         startLine,
		StartLinePosition: startLinePosition,
		StopLine:          stopLine,
		StopLinePosition:  stopLinePosition,
		Override:          isOverrideMethod,
		Annotations:       annotations,
	}

	if reflect.TypeOf(ctx.GetParent().GetParent()).String() == "*parser.ClassBodyDeclarationContext" {
		bodyCtx := ctx.GetParent().GetParent().(*parser.ClassBodyDeclarationContext)
		for _, modifier := range bodyCtx.AllModifier() {
			if !strings.Contains(modifier.GetText(), "@") {
				currentMethod.Modifiers = append(currentMethod.Modifiers, modifier.GetText())
			}
		}
	}

	isOverrideMethod = false
}

func (s *JavaIdentifierListener) ExitMethodDeclaration(ctx *parser.MethodDeclarationContext) {

	currentNode.AddMethod(currentMethod)
	_ = models.NewJMethod()
}

func (s *JavaIdentifierListener) EnterAnnotation(ctx *parser.AnnotationContext) {
	// Todo: support override method
	annotationName := ctx.QualifiedName().GetText()
	if annotationName == "Override" {
		isOverrideMethod = true
	}

	if hasEnterClass {
		annotation := common_listener.BuildAnnotation(ctx)
		currentMethod.Annotations = append(currentMethod.Annotations, annotation)
	} else {
		annotation := common_listener.BuildAnnotation(ctx)
		currentNode.Annotations = append(currentNode.Annotations, annotation)
	}
}

func (s *JavaIdentifierListener) EnterInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	hasEnterClass = true
	currentNode.ClassType = "Interface"
	currentNode.ClassName = ctx.IDENTIFIER().GetText()
}

func (s *JavaIdentifierListener) EnterExpression(ctx *parser.ExpressionContext) {
	if reflect.TypeOf(ctx.GetParent()).String() == "*parser.StatementContext" {
		statementCtx := ctx.GetParent().(*parser.StatementContext)
		firstChild := statementCtx.GetChild(0).(antlr.ParseTree).GetText()
		if strings.ToLower(firstChild) == "return" {
			currentMethod.IsReturnNull = strings.Contains(ctx.GetText(), "null")
		}
	}
}

func (s *JavaIdentifierListener) getNodes() []models.JIdentifier {
	return nodes
}
