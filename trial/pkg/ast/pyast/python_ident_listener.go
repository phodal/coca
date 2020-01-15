package pyast

import (
	"bytes"
	"fmt"
	parser "github.com/phodal/coca/languages/python"
	"github.com/phodal/coca/pkg/domain/trial"
	"io"
	"os"
	"reflect"
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

	switch x := ctx.GetParent().GetChild(0).(type) {
	case *parser.DecoratorContext:
		decorator := BuildDecorator(x)
		dataStruct.Annotations = append(dataStruct.Annotations, decorator)
	default:
		fmt.Fprintf(output, "EnterClassdef: %s\n", reflect.TypeOf(x))
	}

	currentCodeFile.DataStructures = append(currentCodeFile.DataStructures, dataStruct)
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
			Name: "",
			TypeName: argContext.GetText(),
		}
		arguments = append(arguments, *argument)
	}

	return arguments
}

func (s *PythonIdentListener) GetCodeFileInfo() *trial.CodeFile {
	return currentCodeFile
}
