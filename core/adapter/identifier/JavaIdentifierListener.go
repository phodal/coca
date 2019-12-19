package identifier

import (
	"coca/core/languages/java"
	"coca/core/models"
	"fmt"
	"reflect"
)

var node *models.JIdentifier

var currentMethod models.JMethod
var hasEnterClass = false
var hasEnterMethod = false

type JavaIdentifierListener struct {
	parser.BaseJavaParserListener
}

func (s *JavaIdentifierListener) EnterPackageDeclaration(ctx *parser.PackageDeclarationContext) {
	node.Package = ctx.QualifiedName().GetText()
}

func (s *JavaIdentifierListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	hasEnterClass = true

	node.Type = "Class"
	if ctx.IDENTIFIER() != nil {
		node.ClassName = ctx.IDENTIFIER().GetText()
	}

	if ctx.EXTENDS() != nil {
		node.ExtendsName = ctx.TypeType().GetText()
	}
}

func (s *JavaIdentifierListener) ExitClassDeclaration(ctx *parser.ClassDeclarationContext) {
	hasEnterClass = false
}

func (s *JavaIdentifierListener) EnterInterfaceBodyDeclaration(ctx *parser.InterfaceBodyDeclarationContext) {
	fmt.Println(ctx.GetText())

	for _, modifier := range ctx.AllModifier() {
		modifier := modifier.(*parser.ModifierContext).GetChild(0)
		if reflect.TypeOf(modifier.GetChild(0)).String() == "*parser.AnnotationContext" {
			annotationContext := modifier.GetChild(0).(*parser.AnnotationContext)
			currentMethod.Annotations = append(currentMethod.Annotations, annotationContext.QualifiedName().GetText())
		}
	}
}

func (s *JavaIdentifierListener) EnterInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) {
	hasEnterMethod = true

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
	hasEnterMethod = false

	node.AddMethod(currentMethod)
	currentMethod = models.NewJMethod()
}

var isOverrideMethod = false

func (s *JavaIdentifierListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	hasEnterMethod = true

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

	isOverrideMethod = false
}

func (s *JavaIdentifierListener) ExitMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	hasEnterMethod = false

	node.AddMethod(currentMethod)
	currentMethod = models.NewJMethod()
}

func (s *JavaIdentifierListener) EnterAnnotation(ctx *parser.AnnotationContext) {
	// Todo: support override method
	annotationName := ctx.QualifiedName().GetText()
	if annotationName == "Override" {
		isOverrideMethod = true
	}

	if hasEnterClass {
		currentMethod.Annotations = append(currentMethod.Annotations, annotationName)
	}
}

func (s *JavaIdentifierListener) EnterInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	node.Type = "Interface"
	node.ClassName = ctx.IDENTIFIER().GetText()
}

func (s *JavaIdentifierListener) InitNode(identifier *models.JIdentifier) {
	node = identifier
}
