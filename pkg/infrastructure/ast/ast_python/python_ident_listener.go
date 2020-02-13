package ast_python

import (
	"bytes"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/python"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/astutil"
	"io"
	"os"
	"strings"
)

type PythonIdentListener struct {
	parser.BasePythonParserListener
}

var currentCodeFile *core_domain.CodeContainer
var currentDataStruct *core_domain.CodeDataStruct
var debug = false
var output io.Writer
var hasEnterMember = false

func NewPythonIdentListener(fileName string) *PythonIdentListener {
	currentCodeFile = &core_domain.CodeContainer{}
	currentCodeFile.FullName = fileName
	output = os.Stdout

	return &PythonIdentListener{}
}

func (s *PythonIdentListener) SetDebugOutput(isDebug bool) io.Writer {
	output = new(bytes.Buffer)
	debug = isDebug

	return output
}

func (s *PythonIdentListener) EnterImport_stmt(ctx *parser.Import_stmtContext) {
	dotNames := ctx.Dotted_as_names().(*parser.Dotted_as_namesContext).AllDotted_as_name()

	codeImport := &core_domain.CodeImport{}
	context := dotNames[0].(*parser.Dotted_as_nameContext)
	codeImport.Source = context.Dotted_name().GetText()
	if context.Name() != nil {
		codeImport.UsageName = append(codeImport.UsageName, context.Name().GetText())
	}

	for _, usageName := range dotNames[1:] {
		nameContext := usageName.(*parser.Dotted_as_nameContext)
		codeImport.UsageName = append(codeImport.UsageName, nameContext.GetText())
	}

	currentCodeFile.Imports = append(currentCodeFile.Imports, *codeImport)
}

func (s *PythonIdentListener) EnterFrom_stmt(ctx *parser.From_stmtContext) {
	codeImport := &core_domain.CodeImport{}
	codeImport.Source = ctx.From_stmt_source().GetText()
	usageName := ctx.From_stmt_as_names().GetText()

	asNameCtx := ctx.From_stmt_as_names().(*parser.From_stmt_as_namesContext)
	if asNameCtx.OPEN_PAREN() != nil {
		usageName = asNameCtx.Import_as_names().GetText()
	}
	if strings.Contains(usageName, ",") {
		usageNames := strings.Split(usageName, ",")
		codeImport.UsageName = append(codeImport.UsageName, usageNames...)
	} else {
		codeImport.UsageName = append(codeImport.UsageName, usageName)
	}

	currentCodeFile.Imports = append(currentCodeFile.Imports, *codeImport)
}

func (s *PythonIdentListener) EnterClassdef(ctx *parser.ClassdefContext) {
	hasEnterMember = true
	dataStruct := &core_domain.CodeDataStruct{
		NodeName:        ctx.Name().GetText(),
	}

	ctxIndex := astutil.GetNodeIndex(ctx)
	if ctxIndex > 0 {
		decorators := BuildDecoratorsByIndex(ctx, ctxIndex)
		dataStruct.Annotations = decorators
	}

	currentDataStruct = dataStruct
}

func (s *PythonIdentListener) ExitClassdef(ctx *parser.ClassdefContext) {
	hasEnterMember = false
	currentCodeFile.DataStructures = append(currentCodeFile.DataStructures, *currentDataStruct)
	currentDataStruct = nil
}

func (s *PythonIdentListener) EnterFuncdef(ctx *parser.FuncdefContext) {
	hasEnterMember = true
	function := core_domain.CodeFunction{
		Name: ctx.Name().GetText(),
	}

	ctxIndex := astutil.GetNodeIndex(ctx)
	if ctxIndex > 0 {
		decorators := BuildDecoratorsByIndex(ctx, ctxIndex)
		function.Annotations = decorators
	}

	member := core_domain.NewCodeMember()
	member.Name = ctx.Name().GetText()

	if currentDataStruct != nil {
		currentDataStruct.Functions = append(currentDataStruct.Functions, function)
	} else {
		member.FunctionNodes = append(member.FunctionNodes, function)
		currentCodeFile.Members = append(currentCodeFile.Members, *member)
	}
}

func (s *PythonIdentListener) ExitFuncdef(ctx *parser.FuncdefContext) {
	hasEnterMember = false
}

func BuildDecoratorsByIndex(node antlr.ParseTree, index int) []core_domain.CodeAnnotation {
	var nodes []parser.DecoratorContext
	for i := 0; i < index; i++ {
		context := node.GetParent().GetChild(i).(*parser.DecoratorContext)
		nodes = append(nodes, *context)
	}

	var annotations []core_domain.CodeAnnotation
	for _, node := range nodes {
		decorator := BuildDecorator(&node)
		annotations = append(annotations, *decorator)
	}

	return annotations
}

func BuildDecorator(x *parser.DecoratorContext) *core_domain.CodeAnnotation {
	text := x.Dotted_name().GetText()

	annotation := &core_domain.CodeAnnotation{
		Name: text,
	}

	if x.Arglist() != nil {
		annotation.KeyValues = BuildArgList(x.Arglist().(*parser.ArglistContext))
	}

	return annotation
}

func BuildArgList(context *parser.ArglistContext) []core_domain.AnnotationKeyValue {
	var arguments []core_domain.AnnotationKeyValue
	for _, arg := range context.AllArgument() {
		argContext := arg.(*parser.ArgumentContext)
		argument := &core_domain.AnnotationKeyValue{
			Key:   "",
			Value: argContext.GetText(),
		}
		arguments = append(arguments, *argument)
	}

	return arguments
}

func (s *PythonIdentListener) GetCodeFileInfo() core_domain.CodeContainer {
	return *currentCodeFile
}
