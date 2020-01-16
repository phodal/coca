package identifier

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/languages/java"
	"github.com/phodal/coca/pkg/domain/jdomain"
	common_listener2 "github.com/phodal/coca/pkg/infrastructure/ast/common_listener"
	"reflect"
	"strings"
)

var currentNode *jdomain.JIdentifier
var nodes []jdomain.JIdentifier

var currentMethod jdomain.JMethod
var hasEnterClass = false
var imports []string

func NewJavaIdentifierListener() *JavaIdentifierListener {
	nodes = nil
	currentNode = jdomain.NewJIdentifier()
	currentMethod = jdomain.NewJMethod()
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

	currentMethod = jdomain.NewJMethod()
}

func (s *JavaIdentifierListener) ExitClassBody(ctx *parser.ClassBodyContext) {
	hasEnterClass = false
	if currentNode.ClassName != "" {
		currentNode.Methods = currentNode.GetMethods()
		nodes = append(nodes, *currentNode)
	}
	currentNode = jdomain.NewJIdentifier()
}

func (s *JavaIdentifierListener) ExitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	hasEnterClass = false
	if currentNode.ClassName != "" {
		currentNode.Methods = currentNode.GetMethods()
		nodes = append(nodes, *currentNode)
	}
	currentNode = jdomain.NewJIdentifier()
}

func (s *JavaIdentifierListener) EnterConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) {
	currentMethod = jdomain.JMethod{
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
}

func (s *JavaIdentifierListener) EnterInterfaceBodyDeclaration(ctx *parser.InterfaceBodyDeclarationContext) {
	hasEnterClass = true
}

func (s *JavaIdentifierListener) EnterInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().GetColumn()
	name := ctx.IDENTIFIER().GetText()
	//XXX: find the start position of {, not public
	typeType := ctx.TypeTypeOrVoid().GetText()

	if reflect.TypeOf(ctx.GetParent().GetParent().GetChild(0)).String() == "*parser.ModifierContext" {
		common_listener2.BuildAnnotationForMethod(ctx.GetParent().GetParent().GetChild(0).(*parser.ModifierContext), &currentMethod)
	}

	currentMethod = jdomain.JMethod{
		Name:              name,
		Type:              typeType,
		StartLine:         startLine,
		StartLinePosition: startLinePosition,
		StopLine:          stopLine,
		StopLinePosition:  stopLinePosition,
		Override:          isOverrideMethod,
		Annotations:       currentMethod.Annotations,
	}
}

func (s *JavaIdentifierListener) ExitInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) {
	currentNode.AddMethod(currentMethod)
	currentMethod = jdomain.NewJMethod()
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

	if reflect.TypeOf(ctx.GetParent().GetParent().GetChild(0)).String() == "*parser.ModifierContext" {
		common_listener2.BuildAnnotationForMethod(ctx.GetParent().GetParent().GetChild(0).(*parser.ModifierContext), &currentMethod)
	}

	currentMethod = jdomain.JMethod{
		Name:              name,
		Type:              typeType,
		StartLine:         startLine,
		StartLinePosition: startLinePosition,
		StopLine:          stopLine,
		StopLinePosition:  stopLinePosition,
		Override:          isOverrideMethod,
		Annotations:       currentMethod.Annotations,
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
	currentMethod = jdomain.NewJMethod()
}

func (s *JavaIdentifierListener) EnterAnnotation(ctx *parser.AnnotationContext) {
	annotationName := ctx.QualifiedName().GetText()
	if annotationName == "Override" {
		isOverrideMethod = true
	}

	if !hasEnterClass {
		annotation := common_listener2.BuildAnnotation(ctx)
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

func (s *JavaIdentifierListener) GetNodes() []jdomain.JIdentifier {
	return nodes
}
