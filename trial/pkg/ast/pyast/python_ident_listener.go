package pyast

import (
	"bytes"
	"fmt"
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
var hasEnterMember = false

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

func (s *PythonIdentListener) EnterImport_stmt(ctx *parser.Import_stmtContext) {
	var imports []trial.CodeImport
	for _, asName := range ctx.Dotted_as_names().(*parser.Dotted_as_namesContext).AllDotted_as_name() {
		nameContext := asName.(*parser.Dotted_as_nameContext)
		codeImport := BuildCodeImport(nameContext)

		imports = append(imports, *codeImport)
	}

	//currentCodeFile.Imports = append(currentCodeFile.Imports, imports)
}

func BuildCodeImport(nameContext *parser.Dotted_as_nameContext) *trial.CodeImport {
	asNameText := nameContext.Dotted_name().GetText()
	name := ""
	if nameContext.Name() != nil {
		name = nameContext.Name().GetText()
	}

	codeImport := &trial.CodeImport{
		Source: asNameText,
		AsName: name,
	}
	return codeImport
}

func (s *PythonIdentListener) EnterFrom_stmt(ctx *parser.From_stmtContext) {
	if ctx.Dotted_name() != nil {
		asNameText := ctx.Dotted_name().GetText()
		codeImport := &trial.CodeImport{
			Source: asNameText,
			AsName: "",
		}

		fmt.Println(codeImport)
	}
}

func (s *PythonIdentListener) EnterClassdef(ctx *parser.ClassdefContext) {
	hasEnterMember = true
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

func (s *PythonIdentListener) ExitClassdef(ctx *parser.ClassdefContext) {
	hasEnterMember = false
}

func (s *PythonIdentListener) EnterFuncdef(ctx *parser.FuncdefContext) {
	hasEnterMember = true
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

func (s *PythonIdentListener) ExitFuncdef(ctx *parser.FuncdefContext) {
	hasEnterMember = false
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
