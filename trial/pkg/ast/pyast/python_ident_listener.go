package pyast

import (
	"bytes"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/python"
	"github.com/phodal/coca/pkg/domain/trial"
	"io"
	"os"
)

type PythonIdentListener struct {
	parser.BasePythonParserListener
}

var currentCodeFile *trial.CodeFile
var debug = false
var output io.Writer

func NewPythonIdentListener(fileName string) *PythonIdentListener {
	currentCodeFile = &trial.CodeFile{}
	currentCodeFile.FullName = fileName
	output = os.Stdout

	return &PythonIdentListener{}
}

func (s *PythonIdentListener) SetDebugOutput(isDebug bool) io.Writer {
	output = new(bytes.Buffer)
	debug = isDebug

	return output
}

func (s *PythonIdentListener) EnterClassdef(ctx *parser.ClassdefContext) {
	dataStruct := trial.CodeDataStruct{
		Name:       ctx.Name().GetText(),
		ID:         "",
		MemberIds:  nil,
		Properties: nil,
	}

	ctxIndex := GetNodeIndex(ctx)
	if ctxIndex > 0 {
		decorators := BuildDecoratorsByIndex(ctx, ctxIndex)
		dataStruct.Annotations = decorators
	}

	currentCodeFile.DataStructures = append(currentCodeFile.DataStructures, dataStruct)
}

func (s *PythonIdentListener) EnterFuncdef(ctx *parser.FuncdefContext) {
	function := trial.CodeFunction{
		Name: ctx.Name().GetText(),
	}

	ctxIndex := GetNodeIndex(ctx)
	if ctxIndex > 0 {
		decorators := BuildDecoratorsByIndex(ctx, ctxIndex)
		function.Annotations = decorators
	}

	member := &trial.CodeMember{
		Name: ctx.Name().GetText(),
	}

	member.MethodNodes = append(member.MethodNodes, function)
	currentCodeFile.Members = append(currentCodeFile.Members, member)
}

func BuildDecoratorsByIndex(node antlr.ParseTree, index int) []trial.PythonAnnotation {
	var nodes []parser.DecoratorContext
	for i := 0; i < index; i++ {
		context := node.GetParent().GetChild(i).(*parser.DecoratorContext)
		nodes = append(nodes, *context)
	}

	var annotations []trial.PythonAnnotation
	for _, node := range nodes {
		decorator := BuildDecorator(&node)
		annotations = append(annotations, *decorator)
	}

	return annotations
}

func BuildDecorator(x *parser.DecoratorContext) *trial.PythonAnnotation {
	text := x.Dotted_name().GetText()

	annotation := &trial.PythonAnnotation{
		Name: text,
	}

	if x.Arglist() != nil {
		annotation.Properties = BuildArgList(x.Arglist().(*parser.ArglistContext))
	}

	return annotation
}

func BuildArgList(context *parser.ArglistContext) []trial.CodeProperty {
	var arguments []trial.CodeProperty
	for _, arg := range context.AllArgument() {
		argContext := arg.(*parser.ArgumentContext)
		argument := &trial.CodeProperty{
			Name:     "",
			TypeName: argContext.GetText(),
		}
		arguments = append(arguments, *argument)
	}

	return arguments
}

func (s *PythonIdentListener) GetCodeFileInfo() *trial.CodeFile {
	return currentCodeFile
}
